package resolver

import (
	"errors"
	"sas-backend/app/database/models"
	"sas-backend/app/utils"

	"fmt"
	"log"
	"os"
	dbconnect "sas-backend/app/database/connectors/postgres"
	"strconv"
	"strings"

	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
)

// Login func
func Login(params graphql.ResolveParams) (models.LoginModel, error) {
	stringToken := models.LoginModel{}
	username := params.Args["username"].(string)
	password := params.Args["password"].(string)
	var ID int

	attrUser := getUserByUsername(username)

	// deskripsi dan compare password
	loginPassword := bcrypt.CompareHashAndPassword([]byte(attrUser.Password), []byte(password))

	if loginPassword != nil {
		return stringToken, errors.New("Invalid User or Password")
	}

	query := fmt.Sprintf(`
		update sas_user set
			last_login=now()
		where username='%s'
		returning id;
	`,
		username,
	)

	rowsQ, err := dbconnect.Query(query)
	if err != nil {
		log.Printf(err.Error())
		return stringToken, err
	}

	for rowsQ.Next() {
		err = rowsQ.Scan(
			&ID,
		)

		if err != nil {
			log.Printf(err.Error())
			return stringToken, err
		}
	}

	_, errLog := utils.InsertLog("login", "R", "Login "+username, username, ID)

	if errLog != nil {
		return stringToken, nil
	}

	stringToken.Token = utils.GenerateToken(attrUser, getPermission(attrUser.RoleID))

	return stringToken, nil
}

// getUserByUsername func
func getUserByUsername(username string) models.UserTypeModel {
	rowData := models.UserTypeModel{}

	query := fmt.Sprintf(`
		select 
			su.id,su.username,su.password,su.email,su.emp_no,su.fullname,su.grade,su.positions,su.photo,su.role_id,sr.role_name,su.created_at,su.created_by,su.modified_at,su.modified_by
		from sas_user su 
		join sas_role sr on sr.id = su.role_id
		where su.is_deleted = 0 and su.username = '%s';
	`, username)

	rowsQ, err := dbconnect.Query(query)

	if err != nil {
		log.Printf(err.Error())
		return rowData
	}

	for rowsQ.Next() {
		err = rowsQ.Scan(
			&rowData.ID,
			&rowData.Username,
			&rowData.Password,
			&rowData.Email,
			&rowData.EmpNo,
			&rowData.Fullname,
			&rowData.Grade,
			&rowData.Positions,
			&rowData.Photo,
			&rowData.RoleID,
			&rowData.RoleName,
			&rowData.CreatedAt,
			&rowData.CreatedBy,
			&rowData.ModifiedAt,
			&rowData.ModifiedBy,
		)

		if err != nil {
			log.Printf(err.Error())
			return rowData
		}

	}

	return rowData
}

// getPermission func
func getPermission(userRoleID int) map[string]map[string]bool {
	permission := make(map[string]map[string]bool)

	query := fmt.Sprintf(`
		select 
			sra.role_id,sra.module_id,replace(lower(sm.name), ' ', '_') module_code,sm.name,sra.c,sra.r,sra.u,sra.d
		from sas_role_access sra
		join sas_module sm on sm.id = sra.module_id
		where sra.role_id = %d and (sra.c = 1 or sra.r = 1 or sra.u = 1  or sra.d = 1);
	`, userRoleID)

	rowsQ, err := dbconnect.Query(query)

	if err != nil {
		log.Printf(err.Error())
		return permission
	}

	var roleID, moduleID int
	var create, read, update, delete bool
	var moduleCode, moduleName string

	for rowsQ.Next() {
		err = rowsQ.Scan(
			&roleID,
			&moduleID,
			&moduleCode,
			&moduleName,
			&create,
			&read,
			&update,
			&delete,
		)

		aMap := map[string]bool{
			"create": create,
			"read":   read,
			"update": update,
			"delete": delete,
		}
		permission[moduleCode] = aMap

		if err != nil {
			log.Printf(err.Error())
			return permission
		}
	}

	return permission
}

// getAccessByUserID func
func getUserByID(UserID int, isDeleted int) models.UserTypeModel {
	rowData := models.UserTypeModel{}

	query := fmt.Sprintf(`
		select 
			su.id,su.username,su.email,su.emp_no,su.fullname,su.grade,su.positions,su.photo,su.role_id,sr.role_name,su.created_at,su.created_by,su.modified_at,su.modified_by
		from sas_user su 
		join sas_role sr on sr.id = su.role_id
		where su.is_deleted = %d and su.id = %d;
	`, isDeleted, UserID)

	rowsQ, err := dbconnect.Query(query)

	if err != nil {
		log.Printf(err.Error())
		return rowData
	}

	for rowsQ.Next() {
		err = rowsQ.Scan(
			&rowData.ID,
			&rowData.Username,
			&rowData.Email,
			&rowData.EmpNo,
			&rowData.Fullname,
			&rowData.Grade,
			&rowData.Positions,
			&rowData.Photo,
			&rowData.RoleID,
			&rowData.RoleName,
			&rowData.CreatedAt,
			&rowData.CreatedBy,
			&rowData.ModifiedAt,
			&rowData.ModifiedBy,
		)

		if err != nil {
			log.Printf(err.Error())
			return rowData
		}

	}

	return rowData
}

// GetUser func
func GetUser(id int, params graphql.ResolveParams) (models.UserTypeModel, error) {
	rowData := models.UserTypeModel{}

	isAuthorized := utils.IsAuthorized(params.Context.Value("token").(string))
	if isAuthorized == false {
		return rowData, errors.New("Anda tidak memiliki otorisasi")
	}

	decodes := utils.JwtClaim(params.Context.Value("token").(string))

	rowData = getUserByID(id, 0)

	_, errLog := utils.InsertLog("master_data_user", "R", "Read by id master data pengguana "+rowData.Username, decodes.Username, id)

	if errLog != nil {
		return rowData, nil
	}

	return rowData, nil
}

// getListUserCount func
func getListUserCount(params graphql.ResolveParams) int {
	searchs := ""
	filters := params.Args["filters"]
	if filters != nil {
		for _, filterItem := range filters.([]interface{}) {
			tempFilter := filterItem.(map[string]interface{})
			field := utils.CamelToSnakeCase(tempFilter["field"].(string))
			value := tempFilter["value"].(string)
			operator := tempFilter["operator"].(string)

			switch strings.ToLower(field) {
			case "role_name":
				searchs += utils.GenerateFilter("sr."+field, operator, value, "")
			default:
				searchs += utils.GenerateFilter("su."+field, operator, value, "")
			}
		}
	}

	queryCount := fmt.Sprintf(`
		select 
			coalesce(count(1), 0)
		from sas_user su
		join sas_role sr on sr.id = su.role_id
		where su.is_deleted = 0 %s;
	`, searchs)

	return dbconnect.QueryCount(queryCount)
}

// GetListUser func
func GetListUser(params graphql.ResolveParams) (models.ListUserModel, error) {
	rowData := models.ListUserModel{}

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

	orderBy := fmt.Sprintf("order by su.id desc")
	if params.Args["orderBy"] != nil {
		if params.Args["orderBy"].(string) != "" {
			splitted := strings.Split(params.Args["orderBy"].(string), "_")
			col := "id"

			switch strings.ToLower(splitted[0]) {
			case "rolename":
				col = "sr.role_name"
			default:
				col = fmt.Sprintf("su.%s", utils.CamelToSnakeCase(splitted[0]))
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
			case "role_name":
				searchs += utils.GenerateFilter("sr."+field, operator, value, "")
			default:
				searchs += utils.GenerateFilter("su."+field, operator, value, "")
			}
		}
	}

	arrData := []models.UserTypeModel{}
	query := fmt.Sprintf(`
		select 
			su.id,su.username,su.email,su.emp_no,su.fullname,su.grade,su.positions,su.photo,su.role_id,sr.role_name,su.created_at,su.created_by,su.modified_at,su.modified_by
		from sas_user su
		join sas_role sr on sr.id = su.role_id
		where su.is_deleted = 0 %s
		%s
		offset %d rows fetch next %d rows only;
	`, searchs, orderBy, skip, limit)

	rowsQ, err := dbconnect.Query(query)

	if err != nil {
		log.Printf(err.Error())
		return rowData, err
	}

	for rowsQ.Next() {
		item := models.UserTypeModel{}
		err = rowsQ.Scan(
			&item.ID,
			&item.Username,
			&item.Email,
			&item.EmpNo,
			&item.Fullname,
			&item.Grade,
			&item.Positions,
			&item.Photo,
			&item.RoleID,
			&item.RoleName,
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
	rowData.Total = getListUserCount(params)

	_, errLog := utils.InsertLog("master_data_user", "R", "Read list master data user", decodes.Username, 0)

	if errLog != nil {
		return rowData, nil
	}

	return rowData, nil
}

// InsertUser func
func InsertUser(params graphql.ResolveParams) (models.UserTypeModel, error) {
	rowData := models.UserTypeModel{}

	isAuthorized := utils.IsAuthorized(params.Context.Value("token").(string))
	if isAuthorized == false {
		return rowData, errors.New("Anda tidak memiliki otorisasi")
	}

	decodes := utils.JwtClaim(params.Context.Value("token").(string))

	data := params.Args["data"].(map[string]interface{})

	var username string
	usernameTrim := strings.TrimSpace(data["username"].(string))
	checkUsername := fmt.Sprintf(`select su.username from sas_user su where su.username = '%s' and su.is_deleted = 0;`, usernameTrim)
	rowUsername := dbconnect.QueryRow(checkUsername).Scan(&username)

	if rowUsername == nil {
		errMsg := fmt.Sprintf("Username sudah ada '%s'", username)
		return rowData, errors.New(errMsg)
	}

	password := os.Getenv("DEFAULT_PASSWORD_USER")
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	query := fmt.Sprintf(`
		insert into sas_user (username,password,email,emp_no,fullname,grade,positions,photo,role_id,created_at,created_by)
		values ('%s','%s','%s','%s','%s','%s','%s','%s',%d,now(),'%s') returning id,email,fullname,username;
	`,
		usernameTrim,
		hashedPassword,
		data["email"].(string),
		data["empNo"].(string),
		data["fullname"].(string),
		utils.StringNullable(data["grade"]),
		utils.StringNullable(data["positions"]),
		utils.StringNullable(data["photo"]),
		data["roleId"].(int),
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
			&rowData.Email,
			&rowData.Fullname,
			&rowData.Username,
		)

		if err != nil {
			log.Printf(err.Error())
			return rowData, err
		}

		rowData = getUserByID(rowData.ID, 0)
	}

	utils.RegisterNotification(rowData.Email, rowData.Fullname, rowData.Username)

	_, errLog := utils.InsertLog("master_data_user", "C", "Create master data user "+rowData.Username, decodes.Username, rowData.ID)

	if errLog != nil {
		return rowData, nil
	}

	return rowData, nil
}

// UpdateUser func
func UpdateUser(params graphql.ResolveParams) (models.UserTypeModel, error) {
	rowData := models.UserTypeModel{}

	isAuthorized := utils.IsAuthorized(params.Context.Value("token").(string))
	if isAuthorized == false {
		return rowData, errors.New("Anda tidak memiliki otorisasi")
	}

	decodes := utils.JwtClaim(params.Context.Value("token").(string))

	data := params.Args["data"].(map[string]interface{})

	var usernameID int
	usernameTrim := strings.TrimSpace(data["username"].(string))
	checkUsernameID := fmt.Sprintf(`select su.id from sas_user su where su.username = '%s' and su.is_deleted = 0;`, usernameTrim)
	rowUsernameID := dbconnect.QueryRow(checkUsernameID).Scan(&usernameID)

	usernameUpdate := ""
	if rowUsernameID != nil {
		usernameUpdate = fmt.Sprintf(`username='%s',`, usernameTrim)
	} else if usernameID != data["id"] {
		errMsg := fmt.Sprintf("Username sudah ada '%s'", usernameTrim)
		return rowData, errors.New(errMsg)
	}

	query := fmt.Sprintf(`
		update sas_user set
			%s
			email='%s', 
			emp_no='%s', 
			fullname='%s', 
			grade='%s', 
			positions='%s', 
			photo='%s', 
			role_id=%d,
			modified_at=now(), 
			modified_by='%s'
		where id=%d
		returning id;
	`,
		usernameUpdate,
		data["email"],
		data["empNo"],
		data["fullname"],
		utils.StringNullable(data["grade"]),
		utils.StringNullable(data["positions"]),
		utils.StringNullable(data["photo"]),
		data["roleId"],
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

		rowData = getUserByID(rowData.ID, 0)
	}

	return rowData, nil
}

// DeleteUser func
func DeleteUser(id int, params graphql.ResolveParams) (models.UserTypeModel, error) {
	rowData := models.UserTypeModel{}

	isAuthorized := utils.IsAuthorized(params.Context.Value("token").(string))
	if isAuthorized == false {
		return rowData, errors.New("Anda tidak memiliki otorisasi")
	}

	decodes := utils.JwtClaim(params.Context.Value("token").(string))

	query := fmt.Sprintf(`
		update sas_user set
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
		)

		if err != nil {
			log.Printf(err.Error())
			return rowData, err
		}

		rowData = getUserByID(rowData.ID, 1)
	}

	_, errLog := utils.InsertLog("master_data_user", "D", "Delete master data user "+rowData.Username, decodes.Username, id)

	if errLog != nil {
		return rowData, nil
	}

	return rowData, nil
}
