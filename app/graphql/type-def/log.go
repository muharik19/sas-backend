package typedefs

import (
	"github.com/graphql-go/graphql"
)

var actionType = graphql.NewEnum(
	graphql.EnumConfig{
		Name: "actionType",
		Values: graphql.EnumValueConfigMap{
			"C": &graphql.EnumValueConfig{
				Value: "C",
			},
			"R": &graphql.EnumValueConfig{
				Value: "R",
			},
			"U": &graphql.EnumValueConfig{
				Value: "U",
			},
			"D": &graphql.EnumValueConfig{
				Value: "D",
			},
		},
	},
)

// LogType func
func LogType(typeName string) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: typeName,
		Fields: graphql.Fields{
			"id":          &graphql.Field{Type: graphql.Int},
			"modules":     &graphql.Field{Type: graphql.String},
			"refId":       &graphql.Field{Type: graphql.Int},
			"actions":     &graphql.Field{Type: graphql.String},
			"logActivity": &graphql.Field{Type: graphql.String},
			"createdAt":   &graphql.Field{Type: graphql.String},
			"createdBy":   &graphql.Field{Type: graphql.String},
		},
	})
}

//ListLogType function
func ListLogType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "listLogType",
		Fields: graphql.Fields{
			"data":  &graphql.Field{Type: graphql.NewList(LogType("logTypeForList"))},
			"total": &graphql.Field{Type: graphql.Int},
		},
	})
}

// CreateLogType func
func CreateLogType() graphql.FieldConfigArgument {
	data := graphql.InputObjectConfigFieldMap{
		"module":      &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"refId":       &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"actions":     &graphql.InputObjectFieldConfig{Type: actionType},
		"logActivity": &graphql.InputObjectFieldConfig{Type: graphql.NewList(graphql.String)},
	}

	return graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewInputObject(graphql.InputObjectConfig{
				Name:   "logInputType",
				Fields: data,
			}),
		},
	}
}
