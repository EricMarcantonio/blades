package database

import (
	"backend/types"
	"fmt"
	"strings"
)

func GetAProductById(id int, requestedColumns []string) (types.Product, error) {
	rows, err := DB.Query(fmt.Sprintf("SELECT %s FROM skates WHERE id=?", strings.Join(requestedColumns, ",")), id)
	if err != nil {
		return types.Product{}, err
	}
	productsFromRows, err := ExtractProductsFromRows(rows)
	if err != nil {
		return types.Product{}, err
	}
	if len(productsFromRows) == 0 {
		return types.Product{}, nil
	}
	return productsFromRows[0], nil
}

func GetAllProducts(requestedFields []string) ([]types.Product, error) {
	rows, err := DB.Query(fmt.Sprintf("SELECT %s FROM skates WHERE is_active = 'yes' LIMIT 100", strings.Join(requestedFields, ",")))
	if err != nil {
		return nil, err
	}
	productsFromRows, err := ExtractProductsFromRows(rows)
	if err != nil {
		return nil, err
	}
	return productsFromRows, nil
}
