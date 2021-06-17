package schema

import (
	"sas-backend/app/graphql/resolver"
	typedefs "sas-backend/app/graphql/type-def"

	"github.com/graphql-go/graphql"
)

// ModuleSchema func
func ModuleSchema() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(typedefs.ModuleType()),
		Args:        typedefs.FilterArgs("ModuleMasterFilterArgsType"),
		Description: "Get All Module",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return resolver.GetModule(params)
		},
	}
}

// RoleSchema func
func RoleSchema() *graphql.Field {
	return &graphql.Field{
		Type: typedefs.RoleType("RoleTypeById"),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Description: "Get Role By Id",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id := params.Args["id"].(int)
			return resolver.GetRole(id, params)
		},
	}
}

// RolesSchema func
func RolesSchema() *graphql.Field {
	return &graphql.Field{
		Type:        typedefs.ListRoleType(),
		Args:        typedefs.FilterArgs("RoleMasterFilterArgsType"),
		Description: "Get Role",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return resolver.GetListRole(params)
		},
	}
}

// CreateRoleSchema func
func CreateRoleSchema() *graphql.Field {
	return &graphql.Field{
		Args:        typedefs.CreateRoleFieldType(),
		Type:        typedefs.RoleType("returnCreateRoleType"),
		Description: "Create Role",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return resolver.InsertRole(params)
		},
	}
}

// UpdateRoleSchema func
func UpdateRoleSchema() *graphql.Field {
	return &graphql.Field{
		Args:        typedefs.UpdateRoleFieldType(),
		Type:        typedefs.RoleType("returnUpdateRoleType"),
		Description: "Update Role",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return resolver.UpdateRole(params)
		},
	}
}

// DeleteRoleSchema func
func DeleteRoleSchema() *graphql.Field {
	return &graphql.Field{
		Type: typedefs.RoleType("returnDeleteRoleType"),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Description: "Delete Role",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id := params.Args["id"].(int)
			return resolver.DeleteRole(id, params)
		},
	}
}
