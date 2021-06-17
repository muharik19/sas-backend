package typedefs

import (
	"github.com/graphql-go/graphql"
)

var enumOperatorType = graphql.NewEnum(
	graphql.EnumConfig{
		Name: "enumFilterOperatorType",
		Values: graphql.EnumValueConfigMap{
			"eq": &graphql.EnumValueConfig{
				Value: "eq",
			},
			"not_eq": &graphql.EnumValueConfig{
				Value: "not_eq",
			},
			"like": &graphql.EnumValueConfig{
				Value: "like",
			},
			"not_like": &graphql.EnumValueConfig{
				Value: "not_like",
			},
			"greater_than": &graphql.EnumValueConfig{
				Value: "greater_than",
			},
			"less_than": &graphql.EnumValueConfig{
				Value: "less_than",
			},
			"greater_than_equal": &graphql.EnumValueConfig{
				Value: "greater_than_equal",
			},
			"less_than_equal": &graphql.EnumValueConfig{
				Value: "less_than_equal",
			},
			"between": &graphql.EnumValueConfig{
				Value: "between",
			},
			"in": &graphql.EnumValueConfig{
				Value: "in",
			},
		},
	},
)

var enumLogicType = graphql.NewEnum(
	graphql.EnumConfig{
		Name: "enumLogicType",
		Values: graphql.EnumValueConfigMap{
			"and": &graphql.EnumValueConfig{
				Value: "and",
			},
			"or": &graphql.EnumValueConfig{
				Value: "or",
			},
		},
	},
)

var enumDataType = graphql.NewEnum(
	graphql.EnumConfig{
		Name: "enumDataType",
		Values: graphql.EnumValueConfigMap{
			"string": &graphql.EnumValueConfig{
				Value: "string",
			},
			"float": &graphql.EnumValueConfig{
				Value: "float",
			},
			"int": &graphql.EnumValueConfig{
				Value: "int",
			},
		},
	},
)

// FilterArgs ...
func FilterArgs(typeName string) graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"orderBy": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"skip": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"limit": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"filters": &graphql.ArgumentConfig{
			Type: graphql.NewList(
				graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name: typeName,
						Fields: graphql.InputObjectConfigFieldMap{
							"field": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"value": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"operator": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(enumOperatorType),
							},
							"logic": &graphql.InputObjectFieldConfig{
								Type: enumLogicType,
							},
						},
					},
				),
			),
		},
	}
}
