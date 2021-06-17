package models

import "time"

//LoginModel struct
type LoginModel struct {
	Token string `json:"token"`
}

//UserModel struct
type UserModel struct {
	Email     string      `json:"email"`
	RoleID    interface{} `json:"role_id"`
	RoleName  string      `json:"role_name"`
	Fullname  string      `json:"fullname"`
	Username  string      `json:"username"`
	Photo     interface{} `json:"photo"`
	Grade     interface{} `json:"grade"`
	Positions interface{} `json:"positions"`
	Permision interface{} `json:"permision"`
}

//UsersModel struct
type UsersModel struct {
	Total uint16      `json:"total"`
	Data  []UserModel `json:"data"`
}

//UserTypeModel struct
type UserTypeModel struct {
	ID         int         `json:"id"`
	Email      string      `json:"email"`
	Username   string      `json:"username"`
	Password   string      `json:"password"`
	EmpNo      string      `json:"emp_no"`
	Fullname   string      `json:"fullname"`
	Grade      interface{} `json:"grade"`
	Positions  interface{} `json:"positions"`
	Photo      interface{} `json:"photo"`
	RoleID     int         `json:"role_id"`
	RoleName   string      `json:"role_name"`
	CreatedAt  time.Time   `json:"created_at"`
	CreatedBy  string      `json:"created_by"`
	ModifiedAt interface{} `json:"modified_at"`
	ModifiedBy interface{} `json:"modified_by"`
	IsDeleted  interface{} `json:"is_deleted"`
}

//ListUserModel struct
type ListUserModel struct {
	Total int             `json:"total"`
	Data  []UserTypeModel `json:"data"`
}
