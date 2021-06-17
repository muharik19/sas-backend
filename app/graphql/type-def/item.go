package typedefs

import (
	"github.com/graphql-go/graphql"
)

// ItemType func
func ItemType(typeName string) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: typeName,
		Fields: graphql.Fields{
			"id":         &graphql.Field{Type: graphql.Int},
			"itemName":   &graphql.Field{Type: graphql.String},
			"price":      &graphql.Field{Type: graphql.Float},
			"weightML":   &graphql.Field{Type: graphql.Float},
			"weightMG":   &graphql.Field{Type: graphql.Float},
			"createdBy":  &graphql.Field{Type: graphql.String},
			"createdAt":  &graphql.Field{Type: graphql.DateTime},
			"modifiedBy": &graphql.Field{Type: graphql.String},
			"modifiedAt": &graphql.Field{Type: graphql.DateTime},
			"isDeleted":  &graphql.Field{Type: graphql.Int},
		},
	})
}

// ItemsType func
func ItemsType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "ItemsFieldType",
		Fields: graphql.Fields{
			"data":  &graphql.Field{Type: graphql.NewList(ItemType("itemsFieldType"))},
			"total": &graphql.Field{Type: graphql.Int},
		},
	})
}

// ItemCreateFieldType func
func ItemCreateFieldType() graphql.FieldConfigArgument {
	data := graphql.InputObjectConfigFieldMap{
		"itemName": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"price":    &graphql.InputObjectFieldConfig{Type: graphql.Float},
		"weightML": &graphql.InputObjectFieldConfig{Type: graphql.Float},
		"weightMG": &graphql.InputObjectFieldConfig{Type: graphql.Float},
	}

	return graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewInputObject(graphql.InputObjectConfig{
				Name:   "itemCreateFieldType",
				Fields: data,
			}),
		},
	}
}

// ItemUpdateFieldType func
func ItemUpdateFieldType() graphql.FieldConfigArgument {
	data := graphql.InputObjectConfigFieldMap{
		"id":       &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"itemName": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"price":    &graphql.InputObjectFieldConfig{Type: graphql.Float},
		"weightML": &graphql.InputObjectFieldConfig{Type: graphql.Float},
		"weightMG": &graphql.InputObjectFieldConfig{Type: graphql.Float},
	}

	return graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewInputObject(graphql.InputObjectConfig{
				Name:   "itemUpdateFieldType",
				Fields: data,
			}),
		},
	}
}
