package typedefs

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

// ModuleType func
func ModuleType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "moduleType",
		Fields: graphql.Fields{
			"id":   &graphql.Field{Type: graphql.Int},
			"name": &graphql.Field{Type: graphql.String},
		},
	})
}

// RoleAccessType func
func RoleAccessType(typeName string) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: typeName,
		Fields: graphql.Fields{
			"moduleId":   &graphql.Field{Type: graphql.Int},
			"moduleName": &graphql.Field{Type: graphql.String},
			"create":     &graphql.Field{Type: graphql.Boolean},
			"read":       &graphql.Field{Type: graphql.Boolean},
			"update":     &graphql.Field{Type: graphql.Boolean},
			"delete":     &graphql.Field{Type: graphql.Boolean},
		},
	})
}

// RoleType func
func RoleType(typeName string) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: typeName,
		Fields: graphql.Fields{
			"id":         &graphql.Field{Type: graphql.Int},
			"roleName":   &graphql.Field{Type: graphql.String},
			"isActive":   &graphql.Field{Type: graphql.Boolean},
			"createdAt":  &graphql.Field{Type: graphql.DateTime},
			"createdBy":  &graphql.Field{Type: graphql.String},
			"modifiedAt": &graphql.Field{Type: graphql.DateTime},
			"modifiedBy": &graphql.Field{Type: graphql.String},
			"isDeleted":  &graphql.Field{Type: graphql.Int},
			"roleAccess": &graphql.Field{Type: graphql.NewList(RoleAccessType(fmt.Sprintf("RoleAccessTypeFor%s", typeName)))},
		},
	})
}

// RolesType func
func RolesType(typeName string) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: typeName,
		Fields: graphql.Fields{
			"id":         &graphql.Field{Type: graphql.Int},
			"roleName":   &graphql.Field{Type: graphql.String},
			"isActive":   &graphql.Field{Type: graphql.Boolean},
			"createdAt":  &graphql.Field{Type: graphql.DateTime},
			"createdBy":  &graphql.Field{Type: graphql.String},
			"modifiedAt": &graphql.Field{Type: graphql.DateTime},
			"modifiedBy": &graphql.Field{Type: graphql.String},
		},
	})
}

//ListRoleType function
func ListRoleType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "listRoleType",
		Fields: graphql.Fields{
			"data":  &graphql.Field{Type: graphql.NewList(RolesType("RolesType"))},
			"total": &graphql.Field{Type: graphql.Int},
		},
	})
}

// CreateRoleFieldType func
func CreateRoleFieldType() graphql.FieldConfigArgument {
	data := graphql.InputObjectConfigFieldMap{
		"roleName": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"roleAccess": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.NewList(
			graphql.NewInputObject(graphql.InputObjectConfig{
				Name: "RoleAccessInputFieldType",
				Fields: graphql.InputObjectConfigFieldMap{
					"moduleId": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
					"create":   &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Boolean)},
					"read":     &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Boolean)},
					"update":   &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Boolean)},
					"delete":   &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Boolean)},
				},
			}),
		))},
	}

	return graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewInputObject(graphql.InputObjectConfig{
				Name:   "RoleInputFieldType",
				Fields: data,
			}),
		},
	}
}

// UpdateRoleFieldType func
func UpdateRoleFieldType() graphql.FieldConfigArgument {
	data := graphql.InputObjectConfigFieldMap{
		"id":       &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"roleName": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"roleAccess": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.NewList(
			graphql.NewInputObject(graphql.InputObjectConfig{
				Name: "RoleAccessInputUpdateFieldType",
				Fields: graphql.InputObjectConfigFieldMap{
					"moduleId": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
					"create":   &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Boolean)},
					"read":     &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Boolean)},
					"update":   &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Boolean)},
					"delete":   &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Boolean)},
				},
			}),
		))},
	}

	return graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewInputObject(graphql.InputObjectConfig{
				Name:   "RoleInputUpdateFieldType",
				Fields: data,
			}),
		},
	}
}
