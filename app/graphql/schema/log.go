package schema

import (
	"sas-backend/app/graphql/resolver"
	typedefs "sas-backend/app/graphql/type-def"

	"github.com/graphql-go/graphql"
)

// ListLogSchema func
func ListLogSchema() *graphql.Field {
	return &graphql.Field{
		Type:        typedefs.ListLogType(),
		Args:        typedefs.FilterArgs("listLogFilterArgsType"),
		Description: "Get List Log",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return resolver.GetListLog(params)
		},
	}
}

// CreateLogSchema func
func CreateLogSchema() *graphql.Field {
	return &graphql.Field{
		Args:        typedefs.CreateLogType(),
		Type:        typedefs.LogType("returnLogType"),
		Description: "Create Log",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return resolver.InsertLog(params)
		},
	}
}
