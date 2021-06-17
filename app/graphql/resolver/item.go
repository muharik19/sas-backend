package resolver

import (
	"errors"
	"fmt"
	"log"
	"os"
	dbconnect "sas-backend/app/database/connectors/postgres"
	"sas-backend/app/database/models"
	"sas-backend/app/utils"
	"strconv"
	"strings"

	"github.com/graphql-go/graphql"
)

// GetItem func
func GetItem(id int, params graphql.ResolveParams) (models.ItemModel, error) {
	rowData := models.ItemModel{}

	isAuthorized := utils.IsAuthorized(params.Context.Value("token").(string))
	if isAuthorized == false {
		return rowData, errors.New("Anda tidak memiliki otorisasi")
	}

	decodes := utils.JwtClaim(params.Context.Value("token").(string))

	query := fmt.Sprintf(`
		select
			id,item_name,price,weight_ml,weight_mg,is_deleted,created_by,created_at,modified_by,modified_at
		from sas_item
		where is_deleted = 0 and id = %d
	`, id)

	rowsQ, err := dbconnect.Query(query)

	if err != nil {
		log.Printf(err.Error())
		return rowData, err
	}

	for rowsQ.Next() {
		err = rowsQ.Scan(
			&rowData.ID,
			&rowData.ItemName,
			&rowData.Price,
			&rowData.WeightML,
			&rowData.WeightMG,
			&rowData.IsDeleted,
			&rowData.CreatedBy,
			&rowData.CreatedAt,
			&rowData.ModifiedBy,
			&rowData.ModifiedAt,
		)

		if err != nil {
			log.Printf(err.Error())
			return rowData, err
		}
	}

	_, errLog := utils.InsertLog("master_data_item", "R", "Read by id master data item "+rowData.ItemName, decodes.Username, id)

	if errLog != nil {
		return rowData, nil
	}

	return rowData, nil
}

// getListItemCount func
func getListItemCount(params graphql.ResolveParams) int {
	searchs := ""
	filters := params.Args["filters"]
	if filters != nil {
		for _, filterItem := range filters.([]interface{}) {
			tempFilter := filterItem.(map[string]interface{})
			field := utils.CamelToSnakeCase(tempFilter["field"].(string))
			value := tempFilter["value"].(string)
			operator := tempFilter["operator"].(string)

			switch strings.ToLower(field) {
			case "price", "weight_ml", "weight_mg":
				searchs += utils.GenerateFilter("CAST(si."+field+" AS TEXT)", operator, value, "")
			default:
				searchs += utils.GenerateFilter("si."+field, operator, value, "")
			}
		}
	}

	queryCount := fmt.Sprintf(`
		select coalesce(count(1), 0) total from sas_item si where si.is_deleted = 0 %s;
	`, searchs)

	return dbconnect.QueryCount(queryCount)
}

// GetItems func
func GetItems(params graphql.ResolveParams) (models.ItemsModel, error) {
	rowData := models.ItemsModel{}

	isAuthorized := utils.IsAuthorized(params.Context.Value("token").(string))
	if isAuthorized == false {
		return rowData, errors.New("Anda tidak memiliki otorisasi")
	}

	decodes := utils.JwtClaim(params.Context.Value("token").(string))

	skip, _ := strconv.Atoi(os.Getenv("PAGE_SKIP"))
	limit, _ := strconv.Atoi(os.Getenv("PAGE_LIMIT"))
	searchs := ""

	if params.Args["skip"] != nil {
		if params.Args["skip"] != "" {
			skip = params.Args["skip"].(int)
		}
	}

	if params.Args["limit"] != nil {
		if params.Args["limit"] != "" {
			limit = params.Args["limit"].(int)
		}
	}

	orderBy := fmt.Sprintf("order by si.id desc")
	if params.Args["orderBy"] != nil {
		if params.Args["orderBy"].(string) != "" {
			splitted := strings.Split(params.Args["orderBy"].(string), "_")
			col := "id"

			switch strings.ToLower(splitted[0]) {
			default:
				col = fmt.Sprintf("si.%s", utils.CamelToSnakeCase(splitted[0]))
			}

			orderBy = fmt.Sprintf("order by %s %s", col, splitted[1])
		}
	}

	filters := params.Args["filters"]
	if filters != nil {
		for _, filterItem := range filters.([]interface{}) {
			tempFilter := filterItem.(map[string]interface{})
			field := utils.CamelToSnakeCase(tempFilter["field"].(string))
			value := tempFilter["value"].(string)
			operator := tempFilter["operator"].(string)

			switch strings.ToLower(field) {
			case "price", "weight_ml", "weight_mg":
				searchs += utils.GenerateFilter("CAST(si."+field+" AS TEXT)", operator, value, "")
			default:
				searchs += utils.GenerateFilter("si."+field, operator, value, "")
			}
		}
	}

	arrData := []models.ItemModel{}
	query := fmt.Sprintf(`
		select
			si.id,si.item_name,si.price,si.weight_ml,si.weight_mg,si.is_deleted,si.created_by,si.created_at,si.modified_by,si.modified_at
		from sas_item si where si.is_deleted = 0 %s
		%s
			offset %d rows fetch next %d rows only;
	`, searchs, orderBy, skip, limit)

	rowsQ, err := dbconnect.Query(query)

	if err != nil {
		log.Printf(err.Error())
		return rowData, err
	}

	for rowsQ.Next() {
		item := models.ItemModel{}
		err = rowsQ.Scan(
			&item.ID,
			&item.ItemName,
			&item.Price,
			&item.WeightML,
			&item.WeightMG,
			&item.IsDeleted,
			&item.CreatedBy,
			&item.CreatedAt,
			&item.ModifiedBy,
			&item.ModifiedAt,
		)
		if err != nil {
			log.Printf(err.Error())
			return rowData, err
		}

		arrData = append(arrData, item)
	}

	rowData.Data = arrData
	rowData.Total = getListItemCount(params)

	_, errLog := utils.InsertLog("master_data_item", "R", "Read list master data item", decodes.Username, 0)

	if errLog != nil {
		return rowData, nil
	}

	return rowData, nil
}

// InsertItem func
func InsertItem(params graphql.ResolveParams) (models.ItemModel, error) {
	rowData := models.ItemModel{}

	isAuthorized := utils.IsAuthorized(params.Context.Value("token").(string))
	if isAuthorized == false {
		return rowData, errors.New("Anda tidak memiliki otorisasi")
	}

	decodes := utils.JwtClaim(params.Context.Value("token").(string))

	data := params.Args["data"].(map[string]interface{})

	itemNameTrim := strings.TrimSpace(data["itemName"].(string))

	checkItemName := fmt.Sprintf(`select count(1) cnt from sas_item si where si.item_name = '%s' and si.is_deleted = 0;`, itemNameTrim)
	rowItemName := dbconnect.QueryCount(checkItemName)

	if rowItemName > 0 {
		errMsg := fmt.Sprintf("Nama barang sudah ada '%s'", itemNameTrim)
		return rowData, errors.New(errMsg)
	}

	query := fmt.Sprintf(`
		insert into sas_item (item_name,price,weight_ml,weight_mg,created_by,created_at)
		values ('%s',%f,%f,%f,'%s',now()) returning id;
	`,
		itemNameTrim,
		utils.FloatNullable(data["price"]),
		utils.FloatNullable(data["weightML"]),
		utils.FloatNullable(data["weightMG"]),
		decodes.Username,
	)

	rowsQ, err := dbconnect.Query(query)

	if err != nil {
		log.Printf(err.Error())
		return rowData, err
	}

	for rowsQ.Next() {
		err = rowsQ.Scan(
			&rowData.ID,
		)

		if err != nil {
			log.Printf(err.Error())
			return rowData, err
		}
	}

	_, errLog := utils.InsertLog("master_data_item", "C", "Create master data item "+rowData.ItemName, decodes.Username, rowData.ID)

	if errLog != nil {
		return rowData, nil
	}

	return rowData, nil
}

// UpdateItem func
func UpdateItem(params graphql.ResolveParams) (models.ItemModel, error) {
	rowData := models.ItemModel{}

	isAuthorized := utils.IsAuthorized(params.Context.Value("token").(string))
	if isAuthorized == false {
		return rowData, errors.New("Anda tidak memiliki otorisasi")
	}

	decodes := utils.JwtClaim(params.Context.Value("token").(string))

	data := params.Args["data"].(map[string]interface{})

	var itemNameID int
	itemNameTrim := strings.TrimSpace(data["itemName"].(string))

	checkItemNameID := fmt.Sprintf(`select si.id from sas_item si where si.item_name = '%s' and si.is_deleted = 0;`, itemNameTrim)
	rowItemNameID := dbconnect.QueryRow(checkItemNameID).Scan(&itemNameID)

	ItemNameUpdate := ""
	if rowItemNameID != nil {
		ItemNameUpdate = fmt.Sprintf(`item_name='%s',`, itemNameTrim)
	} else if itemNameID != data["id"] {
		errMsg := fmt.Sprintf("Nama barang sudah ada '%s'", itemNameTrim)
		return rowData, errors.New(errMsg)
	}

	query := fmt.Sprintf(`
		update sas_item set
			%s
			price=%f,
			weight_ml=%f,
			weight_mg=%f,
			modified_by='%s',
			modified_at=now()
		WHERE id=%d
		returning id;
	`,
		ItemNameUpdate,
		utils.FloatNullable(data["price"]),
		utils.FloatNullable(data["weightML"]),
		utils.FloatNullable(data["weightMG"]),
		decodes.Username,
		data["id"],
	)

	rowsQ, err := dbconnect.Query(query)

	if err != nil {
		log.Printf(err.Error())
		return rowData, err
	}

	for rowsQ.Next() {
		err = rowsQ.Scan(
			&rowData.ID,
		)

		if err != nil {
			log.Printf(err.Error())
			return rowData, err
		}
	}

	return rowData, nil
}

// DeleteItem func
func DeleteItem(id int, params graphql.ResolveParams) (models.ItemModel, error) {
	rowData := models.ItemModel{}

	isAuthorized := utils.IsAuthorized(params.Context.Value("token").(string))
	if isAuthorized == false {
		return rowData, errors.New("Anda tidak memiliki otorisasi")
	}

	decodes := utils.JwtClaim(params.Context.Value("token").(string))

	query := fmt.Sprintf(`
		update sas_item set
			is_deleted=1, 
			modified_at=now(), 
			modified_by='%s'
		WHERE id=%d
		returning id, is_deleted;
	`,
		decodes.Username,
		id,
	)

	rowsQ, err := dbconnect.Query(query)

	if err != nil {
		log.Printf(err.Error())
		return rowData, err
	}

	for rowsQ.Next() {
		err = rowsQ.Scan(
			&rowData.ID,
			&rowData.IsDeleted,
		)

		if err != nil {
			log.Printf(err.Error())
			return rowData, err
		}
	}

	_, errLog := utils.InsertLog("master_data_item", "D", "Delete master data item "+rowData.ItemName, decodes.Username, rowData.ID)

	if errLog != nil {
		return rowData, nil
	}

	return rowData, nil
}
