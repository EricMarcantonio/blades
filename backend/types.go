package main

import (
	"github.com/graphql-go/graphql"
	"reflect"
)

/**
A Struct
*/
type Skate struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	ModifiedDate string  `json:"modified_date"`
	AddedDate    string  `json:"added_date"`
	IsActive     string  `json:"is_active"`
	Units        int     `json:"units"`
}

var GProduct = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Skate",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
			"modified_date": &graphql.Field{
				Type: graphql.String,
			},
			"added_date": &graphql.Field{
				Type: graphql.String,
			},
			"is_active": &graphql.Field{
				Type: graphql.String,
			},
			"units": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"product": &graphql.Field{
				Type:        GProduct,
				Description: "Get product by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: GetProductById,
			},
			"products": &graphql.Field{
				Type:        graphql.NewList(GProduct),
				Description: "Get product list",
				Resolve:     GetAllProducts,
			},
		},
	})

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createProduct": &graphql.Field{
			Type:        GProduct,
			Description: "Create new product",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"price": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Float),
				},
				"units": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 0,
				},
			},
			Resolve: CreateProduct,
		},
		"updateProduct": &graphql.Field{
			Type:        GProduct,
			Description: "Update a Skate",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"name": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},
				"price": &graphql.ArgumentConfig{
					Type:         graphql.Float,
					DefaultValue: 0.0,
				},
				"is_active": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},
				"units": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: -1, //no value was passed if -1, so we can update with a 0
				},
			},
			Resolve: UpdateAProduct,
		},
		"deactivateProduct": &graphql.Field{
			Type:        GProduct,
			Description: "Update a Skate",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: DeactivateProduct,
		},
	},
})

// StatFields
//
//	Returns the NonNilFields, nilFields
///**
func StatFields(object interface{}) ([]string, []string) {
	val := reflect.ValueOf(object)
	var nilFields []string
	var nonNilFields []string
	for i := 0; i < val.NumField(); i++ {
		if val.Type().Field(i).Name == "Units" && val.Field(i).Interface().(int) < 0 {
			nilFields = append(nilFields, val.Type().Field(i).Name)
			continue
		} else if val.Type().Field(i).Name == "Units" {
			nonNilFields = append(nonNilFields, val.Type().Field(i).Name)
			continue
		}
		if val.Field(i).Interface() == reflect.Zero(val.Type().Field(i).Type).Interface() {
			nilFields = append(nilFields, val.Type().Field(i).Name)
		} else {
			nonNilFields = append(nonNilFields, val.Type().Field(i).Name)
		}
	}
	return nonNilFields, nilFields
}
