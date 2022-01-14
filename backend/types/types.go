package types

import (
	"github.com/graphql-go/graphql"
	"reflect"
)

type Product struct {
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
		Name: "Product",
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
