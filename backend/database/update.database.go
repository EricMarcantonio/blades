package database

import (
	"backend/types"
	"errors"
	"fmt"
	"strings"
)

func UpdateAProductById(product types.Product, requestedFields []string) (types.Product, error) {
	// Check if an ID was given
	if product.ID == 0 {
		return types.Product{}, errors.New("no ID passed to UpdateAProductById")
	}

	//Build the query string with this
	var fieldsToUpdateWithId, _ = types.StatFields(product)
	fieldsToUpdateWithId = fieldsToUpdateWithId[1:]

	if len(fieldsToUpdateWithId) == 0 {
		rows, err := DB.Query(fmt.Sprintf("SELECT %s from skates WHERE id = ? ", strings.Join(requestedFields, ",")), product.ID)
		if err != nil {
			return types.Product{}, err
		}
		productsFromRows, err := ExtractProductsFromRows(rows)
		if err != nil {
			return types.Product{}, err
		}
		return productsFromRows[0], nil
	}
	var queryString strings.Builder
	var cleanedNames []string
	for _, e := range fieldsToUpdateWithId {
		cleanedNames = append(cleanedNames, ProductToMap[e])
	}
	queryString.WriteString(" UPDATE skates set ")
	queryString.WriteString(BuildParamUpdateColumn(cleanedNames))
	queryString.WriteString(", modified_date = NOW()")
	queryString.WriteString(" WHERE id = ? ")
	var err error
	//var tempProduct types.Product
	cols := make([]interface{}, len(cleanedNames)+1)
	for i := 0; i < len(cleanedNames); i++ {
		if cleanedNames[i] == "is_active" {
			if product.IsActive == "yes" {
				cols[i] = 1
			} else {
				cols[i] = 0
			}
			continue
		}
		cols[i] = ProductCol(cleanedNames[i], &product)
	}
	cols[len(cleanedNames)] = &product.ID
	_, err = DB.Exec(queryString.String(), cols...)
	if err != nil {

		return types.Product{}, err
	}
	rows, err := DB.Query(fmt.Sprintf("SELECT %s FROM skates WHERE id = ?", strings.Join(requestedFields, ",")), product.ID)
	if err != nil {
		return types.Product{}, err
	}
	productsFromRows, err := ExtractProductsFromRows(rows)
	if err != nil {
		return types.Product{}, err
	}
	err = rows.Close()
	if err != nil {
		return types.Product{}, err
	}
	return productsFromRows[0], nil
}
