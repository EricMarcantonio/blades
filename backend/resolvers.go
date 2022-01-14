package main

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

/*
	GetProductById
	Handles a request for a single product
	Only sends the requested fields to the db for projection, using GetSelectedFields
	Cleans and parses the request and sends it to a GetAProductByIdFromDB handler
	returns an error if no id is passed (it's enforced by gql but just in case)
	returns a single product, even if the id was not found
*/
func GetProductById(p graphql.ResolveParams) (interface{}, error) {
	i := p.Args["id"].(int)
	if i == 0 {
		return nil, errors.New("no id passed, or wrong type")
	}
	product, err := GetAProductByIdFromDB(i, GetSelectedFields([]string{"product"}, p))
	if err != nil {
		return nil, err
	}
	return product, nil
}

/*
	GetAllProducts
	Handles a request for all products
	Only sends the requested fields to the db for projection, using GetSelectedFields
	Cleans and parses the request and sends it to a GetAllProductsFromDB handler
	returns nil if nothing is returned by the DB
	returns an array of products that are converted to GProduct slice
*/
func GetAllProducts(p graphql.ResolveParams) (interface{}, error) {
	products, err := GetAllProductsFromDB(GetSelectedFields([]string{"products"}, p))
	if err != nil {
		return nil, err
	}
	return products, nil
}

/*
	CreateProduct
	Parses a request and sends it to CreateProductInDB
	Only sends the requested fields to the db for projection, using GetSelectedFields
	returns an error if CreateProductInDB returns an error
	returns a Skate, read and parsed from the database (confirmation of commit)
*/
func CreateProduct(p graphql.ResolveParams) (interface{}, error) {
	if p.Args["name"].(string) == "" {
		return Skate{}, errors.New("no name passed")
	}
	if p.Args["price"].(float64) == 0.0 {
		return Skate{}, errors.New("no price passed")
	}
	product := Skate{
		Name:  p.Args["name"].(string),
		Price: p.Args["price"].(float64),
	}
	product, err := CreateProductInDB(product, GetSelectedFields([]string{"createProduct"}, p))
	if err != nil {
		return Skate{}, err
	}
	return product, nil
}

/*
	UpdateAProduct
	Parses a request to update a Skate in the DB
	Only sends the requested fields to the db for projection, using GetSelectedFields
*/
func UpdateAProduct(p graphql.ResolveParams) (interface{}, error) {
	i := p.Args["id"].(int)
	if i == 0 {
		return Skate{}, errors.New("ID not passed to UpdateAProduct")
	}
	product := Skate{
		ID:    p.Args["id"].(int),
		Name:  p.Args["name"].(string),
		Price: p.Args["price"].(float64),
		Units: p.Args["units"].(int),
	}
	aProductById, err := UpdateAProductById(product, GetSelectedFields([]string{"updateProduct"}, p))
	if err != nil {
		return nil, err
	}
	return aProductById, nil
}

/*
	DeactivateProduct
	Deactivates a product, essentially deleting it from the user view
	returns an error if no ID is passed
*/
func DeactivateProduct(p graphql.ResolveParams) (interface{}, error) {
	i := p.Args["id"].(int)
	if i == 0 {
		return Skate{}, errors.New("ID not passed to Deactivate")
	}

	aProductById, err := DeactivateProductById(i)
	if err != nil {
		return nil, err
	}
	return aProductById, nil
}

/*
	GetSelectedFields
	Parses a graphql.ResolveParams to return the fields requested by the user
	Allows for lean queries; only need to project fields requested by the user
*/
func GetSelectedFields(selectionPath []string, resolveParams graphql.ResolveParams) []string {
	fields := resolveParams.Info.FieldASTs
	for _, propName := range selectionPath {
		found := false
		for _, field := range fields {
			if field.Name.Value == propName {
				selections := field.SelectionSet.Selections
				fields = make([]*ast.Field, 0)
				for _, selection := range selections {
					fields = append(fields, selection.(*ast.Field))
				}
				found = true
				break
			}
		}
		if !found {
			return []string{}
		}
	}
	var collect []string
	for _, field := range fields {
		collect = append(collect, field.Name.Value)
	}
	return collect
}
