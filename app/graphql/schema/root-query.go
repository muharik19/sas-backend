package schema

import (
	"github.com/graphql-go/graphql"
)

func initRootQuery() graphql.ObjectConfig {
	queryFields := graphql.Fields{
		"login":   LoginSchema(),
		"user":    UserSchema(),
		"users":   UsersSchema(),
		"modules": ModuleSchema(),
		"role":    RoleSchema(),
		"roles":   RolesSchema(),
		"item":    ItemSchema(),
		"items":   ItemsSchema(),
		"listLog": ListLogSchema(),
	}
	return graphql.ObjectConfig{
		Name:   "Query",
		Fields: queryFields,
	}
}
