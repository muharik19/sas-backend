package typedefs

import (
	"github.com/graphql-go/graphql"
)

// LoginTypedefs func
func LoginTypedefs() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "loginType",
		Fields: graphql.Fields{
			"token": &graphql.Field{Type: graphql.String},
		},
	})
}

// UserType func
func UserType(typeName string) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: typeName,
		Fields: graphql.Fields{
			"id":         &graphql.Field{Type: graphql.Int},
			"email":      &graphql.Field{Type: graphql.String},
			"username":   &graphql.Field{Type: graphql.String},
			"empNo":      &graphql.Field{Type: graphql.String},
			"fullname":   &graphql.Field{Type: graphql.String},
			"grade":      &graphql.Field{Type: graphql.String},
			"positions":  &graphql.Field{Type: graphql.String},
			"photo":      &graphql.Field{Type: graphql.String},
			"roleId":     &graphql.Field{Type: graphql.Int},
			"roleName":   &graphql.Field{Type: graphql.String},
			"createdAt":  &graphql.Field{Type: graphql.DateTime},
			"createdBy":  &graphql.Field{Type: graphql.String},
			"modifiedAt": &graphql.Field{Type: graphql.DateTime},
			"modifiedBy": &graphql.Field{Type: graphql.String},
			"isDeleted":  &graphql.Field{Type: graphql.Int},
		},
	})
}

//ListUserType function
func ListUserType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "listUserType",
		Fields: graphql.Fields{
			"data":  &graphql.Field{Type: graphql.NewList(UserType("UserListType"))},
			"total": &graphql.Field{Type: graphql.Int},
		},
	})
}

// CreateUserFieldType func
func CreateUserFieldType() graphql.FieldConfigArgument {
	data := graphql.InputObjectConfigFieldMap{
		"email":     &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"username":  &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"empNo":     &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"fullname":  &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"grade":     &graphql.InputObjectFieldConfig{Type: graphql.String},
		"positions": &graphql.InputObjectFieldConfig{Type: graphql.String},
		"photo":     &graphql.InputObjectFieldConfig{Type: graphql.String},
		"roleId":    &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
	}

	return graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewInputObject(graphql.InputObjectConfig{
				Name:   "UserInputFieldType",
				Fields: data,
			}),
		},
	}
}

// UpdateUserFieldType func
func UpdateUserFieldType() graphql.FieldConfigArgument {
	data := graphql.InputObjectConfigFieldMap{
		"id":        &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"email":     &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"username":  &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"empNo":     &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"fullname":  &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"grade":     &graphql.InputObjectFieldConfig{Type: graphql.String},
		"positions": &graphql.InputObjectFieldConfig{Type: graphql.String},
		"photo":     &graphql.InputObjectFieldConfig{Type: graphql.String},
		"roleId":    &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
	}

	return graphql.FieldConfigArgument{
		"data": &graphql.ArgumentConfig{
			Type: graphql.NewInputObject(graphql.InputObjectConfig{
				Name:   "UserUpdateFieldType",
				Fields: data,
			}),
		},
	}
}
