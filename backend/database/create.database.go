package database

import (
	"backend/types"
	"fmt"
	"strings"
)

// CreateProduct
//- creates a product in the skates table
///**
func CreateProduct(product types.Product, requestedColumns []string) (types.Product, error) {
	_, err := DB.Exec("INSERT INTO skates (name, price, modified_date, added_date, is_active, units) VALUES (?, ?, NOW(), NOW(),'yes', ?)", product.Name, product.Price, product.Units)
	if err != nil {
		return types.Product{}, err
	}
	rows, err := DB.Query(fmt.Sprintf("SELECT %s from skates WHERE id = LAST_INSERT_ID()", strings.Join(requestedColumns, ",")))
	if err != nil {
		return types.Product{}, err
	}
	productsFromRows, err := ExtractProductsFromRows(rows)
	if err != nil {
		return types.Product{}, err
	}
	return productsFromRows[0], nil
}
