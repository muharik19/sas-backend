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

// getAccessByRoleID func
func getAccessByRoleID(roleID int) []models.RoleAccess {
	rowData := []models.RoleAccess{}

	query := fmt.Sprintf(`
		select 
			sm.id module_id,sm.name module_name,coalesce(sra.c, 0) c,coalesce(sra.r, 0) r,coalesce(sra.u, 0) u,coalesce(sra.d, 0) d
		from sas_module sm
		left join sas_role_access sra on sra.module_id = sm.id and sra.role_id = %d;
	`, roleID)
	rowsQ, err := dbconnect.Query(query)

	if err != nil {
		log.Printf(err.Error())
		return rowData
	}

	for rowsQ.Next() {
		item := models.RoleAccess{}
		err = rowsQ.Scan(
			&item.ModuleID,
			&item.ModuleName,
			&item.Create,
			&item.Read,
			&item.Update,
			&item.Delete,
		)

		if err != nil {
			log.Printf(err.Error())
			return rowData
		}

		rowData = append(rowData, item)
	}

	return rowData
}

// GetRole func
func GetRole(id int, params graphql.ResolveParams) (models.RoleModel, error) {
	rowData := models.RoleModel{}

	isAuthorized := utils.IsAuthorized(params.Context.Value("token").(string))
	if isAuthorized == false {
		return rowData, errors.New("Anda tidak memiliki otorisasi")
	}

	decodes := utils.JwtClaim(params.Context.Value("token").(string))

	query := fmt.Sprintf(`
		select 
			sr.id,sr.role_name,sr.is_active,sr.created_at,sr.created_by,sr.modified_at,sr.modified_by
		from sas_role sr where sr.id = %d and sr.is_deleted = 0;
	`, id)

	rowsQ, err := dbconnect.Query(query)

	if err != nil {
		log.Printf(err.Error())
		return rowData, err
	}

	for rowsQ.Next() {
		err = rowsQ.Scan(
			&rowData.ID,
			&rowData.RoleName,
			&rowData.IsActive,
			&rowData.CreatedAt,
			&rowData.CreatedBy,
			&rowData.ModifiedAt,
			&rowData.ModifiedBy,
		)

		if err != nil {
			log.Printf(err.Error())
			return rowData, err
		}

		rowData.RoleAccess = getAccessByRoleID(rowData.ID)
	}

	_, errLog := utils.InsertLog("master_data_group_access", "R", "Read by id master data grup akses "+rowData.RoleName, decodes.Username, id)

	if errLog != nil {
		return rowData, nil
	}

	return rowData, nil
}

// getListRoleCount func
func getListRoleCount(params graphql.ResolveParams) int {
	searchs := ""
	filters := params.Args["filters"]
	if filters != nil {
		for _, filterItem := range filters.([]interface{}) {
			tempFilter := filterItem.(map[string]interface{})
			field := utils.CamelToSnakeCase(tempFilter["field"].(string))
			value := tempFilter["value"].(string)
			operator := tempFilter["operator"].(string)

			switch strings.ToLower(field) {
			case "is_active":
				val := "1"
				if value == "false" {
					val = "0"
				}
				searchs += utils.GenerateFilter("sr."+field, operator, val, "int")
			default:
				searchs += utils.GenerateFilter("sr."+field, operator, value, "")
			}
		}
	}

	queryCount := fmt.Sprintf(`
		select 
			coalesce(count(1), 0)
		from sas_role sr where sr.is_deleted = 0 %s;
	`, searchs)

	return dbconnect.QueryCount(queryCount)
}

// GetListRole func
func GetListRole(params graphql.ResolveParams) (models.ListRoleModel, error) {
	rowData := models.ListRoleModel{}

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

	orderBy := fmt.Sprintf("order by sr.id desc")
	if params.Args["orderBy"] != nil {
		if params.Args["orderBy"].(string) != "" {
			splitted := strings.Split(params.Args["orderBy"].(string), "_")
			col := "id"

			switch strings.ToLower(splitted[0]) {
			default:
				col = fmt.Sprintf("sr.%s", utils.CamelToSnakeCase(splitted[0]))
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
			case "is_active":
				val := "1"
				if value == "false" {
					val = "0"
				}
				searchs += utils.GenerateFilter("sr."+field, operator, val, "int")
			default:
				searchs += utils.GenerateFilter("sr."+field, operator, value, "")
			}
		}
	}

	arrData := []models.RoleModel{}
	query := fmt.Sprintf(`
		select 
			sr.id,sr.role_name,sr.is_active,sr.created_at,sr.created_by,sr.modified_at,sr.modified_by
		from sas_role sr
		where sr.is_deleted = 0 %s
		%s
		offset %d rows fetch next %d rows only;
	`, searchs, orderBy, skip, limit)

	rowsQ, err := dbconnect.Query(query)

	if err != nil {
		log.Printf(err.Error())
		return rowData, err
	}

	for rowsQ.Next() {
		item := models.RoleModel{}
		err = rowsQ.Scan(
			&item.ID,
			&item.RoleName,
			&item.IsActive,
			&item.CreatedAt,
			&item.CreatedBy,
			&item.ModifiedAt,
			&item.ModifiedBy,
		)
		if err != nil {
			log.Printf(err.Error())
			return rowData, err
		}

		arrData = append(arrData, item)
	}

	rowData.Data = arrData
	rowData.Total = getListRoleCount(params)

	_, err = utils.InsertLog("master_data_group_access", "R", "Read list master data grup akses", decodes.Username, 0)

	if err != nil {
		return rowData, nil
	}

	return rowData, nil
}

// GetModule func
func GetModule(params graphql.ResolveParams) ([]models.Modules, error) {
	rowData := []models.Modules{}

	isAuthorized := utils.IsAuthorized(params.Context.Value("token").(string))
	if isAuthorized == false {
		return rowData, errors.New("Anda tidak memiliki otorisasi")
	}

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

	orderBy := fmt.Sprintf("order by sm.id desc")
	if params.Args["orderBy"] != nil {
		if params.Args["orderBy"].(string) != "" {
			splitted := strings.Split(params.Args["orderBy"].(string), "_")
			col := "id"

			switch strings.ToLower(splitted[0]) {
			default:
				col = fmt.Sprintf("sm.%s", utils.CamelToSnakeCase(splitted[0]))
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
			default:
				searchs += utils.GenerateFilter(field, operator, value, "")
			}
		}
	}

	query := fmt.Sprintf(`
		select
			sm.id, sm.name from sas_module sm
		where 1=1 %s
		%s
		offset %d rows fetch next %d rows only;
	`, searchs, orderBy, skip, limit)

	rowsQ, err := dbconnect.Query(query)

	if err != nil {
		log.Printf(err.Error())
		return rowData, err
	}

	for rowsQ.Next() {
		item := models.Modules{}
		err = rowsQ.Scan(
			&item.ID,
			&item.Name,
		)
		if err != nil {
			log.Printf(err.Error())
			return rowData, err
		}

		rowData = append(rowData, item)
	}

	return rowData, nil
}

// InsertRole func
func InsertRole(params graphql.ResolveParams) (models.RoleModel, error) {
	rowData := models.RoleModel{}

	isAuthorized := utils.IsAuthorized(params.Context.Value("token").(string))
	if isAuthorized == false {
		return rowData, errors.New("Anda tidak memiliki otorisasi")
	}

	decodes := utils.JwtClaim(params.Context.Value("token").(string))

	data := params.Args["data"].(map[string]interface{})

	query := fmt.Sprintf(`
		insert into sas_role (role_name,created_at,created_by)
		values ('%s',now(),'%s') returning id;
	`,
		data["roleName"],
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

		roleAccess := data["roleAccess"].([]interface{})

		insertRoleAccess := "insert into sas_role_access (role_id,module_id,c,r,u,d) values "

		loop := 1
		for _, item := range roleAccess {
			field := item.(map[string]interface{})
			insertRoleAccess += fmt.Sprintf(`(%d,%d,%d,%d,%d,%d)`, rowData.ID, field["moduleId"], utils.BoolToInteger(field["create"].(bool)), utils.BoolToInteger(field["read"].(bool)), utils.BoolToInteger(field["update"].(bool)), utils.BoolToInteger(field["delete"].(bool)))
			if loop == len(roleAccess) {
				insertRoleAccess += ";"
			} else {
				insertRoleAccess += ","
			}

			loop++
		}

		_, errInsert := dbconnect.Query(insertRoleAccess)

		if errInsert == nil {
			rowData.RoleAccess = getAccessByRoleID(rowData.ID)
		}
	}

	_, errLog := utils.InsertLog("master_data_group_access", "C", "Create master data grup akses "+rowData.RoleName, decodes.Username, rowData.ID)

	if errLog != nil {
		return rowData, nil
	}

	return rowData, nil
}

// UpdateRole func
func UpdateRole(params graphql.ResolveParams) (models.RoleModel, error) {
	rowData := models.RoleModel{}

	isAuthorized := utils.IsAuthorized(params.Context.Value("token").(string))
	if isAuthorized == false {
		return rowData, errors.New("Anda tidak memiliki otorisasi")
	}

	decodes := utils.JwtClaim(params.Context.Value("token").(string))

	data := params.Args["data"].(map[string]interface{})

	query := fmt.Sprintf(`
		update sas_role set
			role_name='%s',
			is_active=%d,
			modified_at=now(),
			modified_by='%s'
		where id=%d
		returning id;
	`,
		data["roleName"],
		utils.IntNullableForIsActive(data["isActive"]),
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

		roleAccess := data["roleAccess"].([]interface{})

		deleteRoleAccessByID := fmt.Sprintf(`delete from sas_role_access where role_id=%d;`, rowData.ID)
		insertRoleAccess := "insert into sas_role_access (role_id,module_id,c,r,u,d) values "

		loop := 1
		for _, item := range roleAccess {
			field := item.(map[string]interface{})
			insertRoleAccess += fmt.Sprintf(`(%d,%d,%d,%d,%d,%d)`, rowData.ID, field["moduleId"], utils.BoolToInteger(field["create"].(bool)), utils.BoolToInteger(field["read"].(bool)), utils.BoolToInteger(field["update"].(bool)), utils.BoolToInteger(field["delete"].(bool)))
			if loop == len(roleAccess) {
				insertRoleAccess += ";"
			} else {
				insertRoleAccess += ","
			}

			loop++
		}

		execQuery := fmt.Sprintf(`%s%s`, deleteRoleAccessByID, insertRoleAccess)

		_, errInsert := dbconnect.Query(execQuery)

		if errInsert == nil {
			rowData.RoleAccess = getAccessByRoleID(rowData.ID)
		}
	}

	return rowData, nil
}

// DeleteRole func
func DeleteRole(id int, params graphql.ResolveParams) (models.RoleModel, error) {
	rowData := models.RoleModel{}

	isAuthorized := utils.IsAuthorized(params.Context.Value("token").(string))
	if isAuthorized == false {
		return rowData, errors.New("Anda tidak memiliki otorisasi")
	}

	decodes := utils.JwtClaim(params.Context.Value("token").(string))

	query := fmt.Sprintf(`
		update sas_role set
			is_deleted=1,
			modified_at=now(), 
			modified_by='%s'
		where id=%d
		returning id,is_deleted;
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

		rowData.RoleAccess = getAccessByRoleID(rowData.ID)
	}

	_, err = utils.InsertLog("master_data_group_access", "D", "Delete master data grup akses "+rowData.RoleName, decodes.Username, id)

	if err != nil {
		return rowData, nil
	}

	return rowData, nil
}
