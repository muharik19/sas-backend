package models

import (
	"time"
)

// Modules ...
type Modules struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// RoleAccess ...
type RoleAccess struct {
	ID         int    `json:"id"`
	ModuleID   int    `json:"module_id"`
	ModuleName string `json:"module_name"`
	Create     bool   `json:"c"`
	Read       bool   `json:"r"`
	Update     bool   `json:"u"`
	Delete     bool   `json:"d"`
}

// ListRoleModel struct
type ListRoleModel struct {
	Data  []RoleModel `json:"data"`
	Total int         `json:"total"`
}

// RoleModel struct
type RoleModel struct {
	ID         int          `json:"id"`
	RoleName   string       `json:"role_name"`
	IsActive   bool         `json:"is_active"`
	CreatedAt  time.Time    `json:"created_at"`
	CreatedBy  string       `json:"created_by"`
	ModifiedAt interface{}  `json:"modified_at"`
	ModifiedBy interface{}  `json:"modified_by"`
	IsDeleted  interface{}  `json:"is_deleted"`
	RoleAccess []RoleAccess `json:"access"`
}
