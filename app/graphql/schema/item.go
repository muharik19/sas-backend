package schema

import (
	"sas-backend/app/graphql/resolver"
	typedefs "sas-backend/app/graphql/type-def"

	"github.com/graphql-go/graphql"
)

// ItemsSchema func
func ItemsSchema() *graphql.Field {
	return &graphql.Field{
		Type:        typedefs.ItemsType(),
		Args:        typedefs.FilterArgs("itemsFilterArgsType"),
		Description: "Get All Item",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return resolver.GetItems(params)
		},
	}
}

// ItemSchema func
func ItemSchema() *graphql.Field {
	return &graphql.Field{
		Type: typedefs.ItemType("itemFieldType"),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Description: "Get Item By ID",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			ID := params.Args["id"].(int)
			return resolver.GetItem(ID, params)
		},
	}
}

// CreateItemSchema func
func CreateItemSchema() *graphql.Field {
	return &graphql.Field{
		Args:        typedefs.ItemCreateFieldType(),
		Type:        typedefs.ItemType("itemCreateType"),
		Description: "Create New Item",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return resolver.InsertItem(params)
		},
	}
}

// UpdateItemSchema func
func UpdateItemSchema() *graphql.Field {
	return &graphql.Field{
		Args:        typedefs.ItemUpdateFieldType(),
		Type:        typedefs.ItemType("itemUpdateType"),
		Description: "Update Item",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return resolver.UpdateItem(params)
		},
	}
}

// DeleteItemSchema func
func DeleteItemSchema() *graphql.Field {
	return &graphql.Field{
		Type: typedefs.ItemType("itemDeleteFieldType"),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Description: "Delete Item",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id := params.Args["id"].(int)
			return resolver.DeleteItem(id, params)
		},
	}
}
