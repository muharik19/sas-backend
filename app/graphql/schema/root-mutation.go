package schema

import (
	"github.com/graphql-go/graphql"
)

func initRootMutation() graphql.ObjectConfig {
	mutationFields := graphql.Fields{
		"userCreate": CreateUserSchema(),
		"userUpdate": UpdateUserSchema(),
		"userDelete": DeleteUserSchema(),
		"roleCreate": CreateRoleSchema(),
		"roleUpdate": UpdateRoleSchema(),
		"roleDelete": DeleteRoleSchema(),
		"itemCreate": CreateItemSchema(),
		"itemUpdate": UpdateItemSchema(),
		"itemDelete": DeleteItemSchema(),
		"logCreate":  CreateLogSchema(),
	}
	return graphql.ObjectConfig{
		Name:   "Mutation",
		Fields: mutationFields,
	}
}
