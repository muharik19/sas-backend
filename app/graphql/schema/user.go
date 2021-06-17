package schema

import (
	"sas-backend/app/graphql/resolver"
	typedefs "sas-backend/app/graphql/type-def"

	"github.com/graphql-go/graphql"
)

//LoginSchema func
func LoginSchema() *graphql.Field {
	return &graphql.Field{
		Type: typedefs.LoginTypedefs(),
		Args: graphql.FieldConfigArgument{
			"username": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Description: "User Login",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return resolver.Login(params)
		},
	}
}

// UserSchema func
func UserSchema() *graphql.Field {
	return &graphql.Field{
		Type: typedefs.UserType("UserTypeById"),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Description: "Get User By Id",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id := params.Args["id"].(int)
			return resolver.GetUser(id, params)
		},
	}
}

// ListUserSchema func
func UsersSchema() *graphql.Field {
	return &graphql.Field{
		Type:        typedefs.ListUserType(),
		Args:        typedefs.FilterArgs("UserMasterFilterArgsType"),
		Description: "Get User",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return resolver.GetListUser(params)
		},
	}
}

// CreateUserSchema func
func CreateUserSchema() *graphql.Field {
	return &graphql.Field{
		Args:        typedefs.CreateUserFieldType(),
		Type:        typedefs.UserType("returnCreateUserType"),
		Description: "Create User",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return resolver.InsertUser(params)
		},
	}
}

// UpdateUserSchema func
func UpdateUserSchema() *graphql.Field {
	return &graphql.Field{
		Args:        typedefs.UpdateUserFieldType(),
		Type:        typedefs.UserType("returnUpdateUserType"),
		Description: "Update User",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return resolver.UpdateUser(params)
		},
	}
}

// DeleteUserSchema func
func DeleteUserSchema() *graphql.Field {
	return &graphql.Field{
		Type: typedefs.UserType("returnDeleteUserType"),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Description: "Delete User By Id",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id := params.Args["id"].(int)
			return resolver.DeleteUser(id, params)
		},
	}
}
