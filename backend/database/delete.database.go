package database

import "backend/types"

func DeactivateProductById(id int) (types.Product, error) {
	_, err := DB.Exec("UPDATE skates set is_active = 'no' where id = ?", id)
	if err != nil {
		return types.Product{}, err
	}
	rows, err := DB.Query("SELECT id, is_active from skates where id = ?", id)
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
	err = rows.Close()
	if err != nil {
		return types.Product{}, err
	}
	return productsFromRows[0], nil
}
