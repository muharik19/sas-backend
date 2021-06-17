package schema

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

//InitGraphQLSchema func
func InitGraphQLSchema() graphql.Schema {
	rootQuery := initRootQuery()
	rootMutation := initRootMutation()

	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: graphql.NewObject(rootMutation),
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		fmt.Printf("failed to create new schema, error: %v", err)
	}
	return schema
}
