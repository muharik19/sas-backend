package utils

import (
	"fmt"
	"log"
	dbconnect "sas-backend/app/database/connectors/postgres"
	"sas-backend/app/database/models"
)

// InsertLog func
func InsertLog(module, action, logActivity, createdBy string, refID int) (models.LogModel, error) {
	rowData := models.LogModel{}

	query := fmt.Sprintf(`
		insert into sas_audit_trail (modules,ref_id,actions,log_activity,created_at,created_by)
		values('%s',%d,'%s','%s',now(),'%s') returning id,modules,ref_id,actions,log_activity,created_at,created_by
	`, module, refID, action, logActivity, createdBy)

	rowsQ, err := dbconnect.Query(query)

	if err != nil {
		log.Printf(err.Error())
		return rowData, err
	}

	for rowsQ.Next() {
		err = rowsQ.Scan(
			&rowData.ID,
			&rowData.Modules,
			&rowData.RefID,
			&rowData.Actions,
			&rowData.LogActivity,
			&rowData.CreatedAt,
			&rowData.CreatedBy,
		)
		if err != nil {
			return rowData, err
		}
	}
	return rowData, nil
}
