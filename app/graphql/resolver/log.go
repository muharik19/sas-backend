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

// getListLogCount func
func getListLogCount(params graphql.ResolveParams) int {
	searchs := ""
	filters := params.Args["filters"]
	if filters != nil {
		for _, filterItem := range filters.([]interface{}) {
			tempFilter := filterItem.(map[string]interface{})
			field := utils.CamelToSnakeCase(tempFilter["field"].(string))
			value := tempFilter["value"].(string)
			operator := tempFilter["operator"].(string)

			switch field {
			default:
				searchs += utils.GenerateFilter("sat."+field, operator, value, "")
			}
		}
	}

	queryCount := fmt.Sprintf(`
		select
			coalesce(count(1), 0)
		from sas_audit_trail sat
		where sat.actions in ('C', 'U')  %s;
	`, searchs)

	return dbconnect.QueryCount(queryCount)
}

// GetListLog func
func GetListLog(params graphql.ResolveParams) (models.ListLogModel, error) {
	rowData := models.ListLogModel{}

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

	orderBy := fmt.Sprintf("order by sat.id desc")
	if params.Args["orderBy"] != nil {
		if params.Args["orderBy"].(string) != "" {
			splitted := strings.Split(params.Args["orderBy"].(string), "_")
			col := "id"

			switch strings.ToLower(splitted[0]) {
			default:
				col = fmt.Sprintf("sat.%s", utils.CamelToSnakeCase(splitted[0]))
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

			switch field {
			default:
				searchs += utils.GenerateFilter("sat."+field, operator, value, "")
			}
		}
	}

	arrData := []models.LogModel{}
	query := fmt.Sprintf(`
		select
			sat.id,sat.modules,sat.ref_id,sat.actions,sat.log_activity,sat.created_at,sat.created_by
		from sas_audit_trail sat
		where sat.actions in ('C', 'U')  %s
		%s
		offset %d rows fetch next %d rows only;
	`, searchs, orderBy, skip, limit)

	rowsQ, err := dbconnect.Query(query)

	if err != nil {
		log.Printf(err.Error())
		return rowData, err
	}

	for rowsQ.Next() {
		item := models.LogModel{}
		err = rowsQ.Scan(
			&item.ID,
			&item.Modules,
			&item.RefID,
			&item.Actions,
			&item.LogActivity,
			&item.CreatedAt,
			&item.CreatedBy,
		)
		if err != nil {
			log.Printf(err.Error())
			return rowData, err
		}

		arrData = append(arrData, item)
	}

	rowData.Data = arrData
	rowData.Total = getListLogCount(params)

	return rowData, nil
}

// InsertLog func
func InsertLog(params graphql.ResolveParams) (models.LogModel, error) {
	rowData := models.LogModel{}

	isAuthorized := utils.IsAuthorized(params.Context.Value("token").(string))
	if isAuthorized == false {
		return rowData, errors.New("Anda tidak memiliki otorisasi")
	}

	decodes := utils.JwtClaim(params.Context.Value("token").(string))

	data := params.Args["data"].(map[string]interface{})

	actions := "U"
	if data["action"] != nil {
		actions = data["action"].(string)
	}

	logAct := ""
	if data["logActivity"] != nil {
		arrData := data["logActivity"].([]interface{})
		for _, v := range arrData {
			logAct += v.(string) + "\n"
		}
	}

	r, err := utils.InsertLog(data["module"].(string), actions, logAct, decodes.Username, data["refId"].(int))

	if err != nil {
		return rowData, nil
	}

	return r, nil
}
