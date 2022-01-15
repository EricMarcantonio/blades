package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var DB *sql.DB

var ProductToMap = map[string]string{
	"id":           "id",
	"Name":         "name",
	"Price":        "price",
	"ModifiedDate": "modified_date",
	"AddedDate":    "added_date",
	"IsActive":     "is_active",
	"Units":        "units",
}

/*
	Connect to a Maria/MysqlDB using EnvVars if provided
*/
func ConnectToDatabase() (error, *sql.DB) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", localDev("user"), localDev("pass"), localDev("domain"), localDev("port"), localDev("db")))
	if err != nil {
		return err, nil
	}
	return nil, db
}

func env(key string) string {
	v, exists := os.LookupEnv(key)
	if !exists {
		return ""
	} else {
		return v
	}
}

func localDev(param string) string {
	if env(param) == "" {
		switch param {
		case "user":
			return "backend"
		case "db":
			return "blades"
		case "pass":
			return "password"
		case "port":
			return "3306"
		case "domain":
			return "127.0.0.1"
		}
	}
	return env(param)
}

/*
	ExtractProductsFromRows
	Takes rows returned by a query and converts it to a Skate slice
	Dynamic setting of Skate fields, not all fields are guaranteed
*/
func ExtractProductsFromRows(rows *sql.Rows) ([]Skate, error) {
	var tempProducts []Skate
	var colNames []string
	var err error
	colNames, err = rows.Columns()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var tempProduct Skate

		cols := make([]interface{}, len(colNames))
		for i := 0; i < len(colNames); i++ {
			cols[i] = ProductCol(colNames[i], &tempProduct)
		}
		err := rows.Scan(cols...)

		if err != nil {
			return nil, err
		}
		tempProducts = append(tempProducts, tempProduct)
	}
	if rows.Err() != nil {
		panic(err)
	}
	return tempProducts, nil
}

/*
	ProductCol
	return the address of a field in a skate
*/
func ProductCol(ColName string, product *Skate) interface{} {
	switch ColName {
	case "id":
		return &product.ID
	case "name":
		return &product.Name
	case "price":
		return &product.Price
	case "modified_date":
		return &product.ModifiedDate
	case "added_date":
		return &product.AddedDate
	case "is_active":
		return &product.IsActive
	case "units":
		return &product.Units
	default:
		panic("Not implemented")
	}
}

/*
	BuildParamUpdateColumn
	Created parameterized queries to combat SQL injections
	Creates a string with parameters built in
*/
func BuildParamUpdateColumn(changingColumns []string) string {
	var tempString strings.Builder
	var columnName []string
	for _, eachString := range changingColumns {
		switch eachString {
		case "name":
			columnName = append(columnName, "name = ?")
		case "price":
			columnName = append(columnName, "price = ?")
		case "is_active":
			columnName = append(columnName, "is_active = ?")
		case "units":
			columnName = append(columnName, "units = ?")
		}
	}
	if len(columnName) == 0 {
		return ""
	}
	tempString.WriteString(strings.Join(columnName, ","))
	return tempString.String()
}

/*
	CreateProductInDB
	Creates a Skate in the db, returning the fields requested by the user
	Returns an error if any queries go wrong
*/
func CreateProductInDB(product Skate, requestedColumns []string) (Skate, error) {
	_, err := DB.Exec("INSERT INTO skates (name, price, modified_date, added_date, is_active, units) VALUES (?, ?, UTC_TIMESTAMP(), UTC_TIMESTAMP(),'yes', ?)", product.Name, product.Price, product.Units)
	if err != nil {
		return Skate{}, err
	}
	rows, err := DB.Query(fmt.Sprintf("SELECT %s from skates WHERE id = LAST_INSERT_ID()", strings.Join(requestedColumns, ",")))
	if err != nil {
		return Skate{}, err
	}
	productsFromRows, err := ExtractProductsFromRows(rows)
	if err != nil {
		return Skate{}, err
	}
	err = rows.Close()
	if err != nil {
		return Skate{}, err
	}
	if len(productsFromRows) == 0 {
		return Skate{}, nil
	}
	return productsFromRows[0], nil
}

/*
	GetAProductByIdFromDB
	Return a skate from the database, using the id, with only the fields requested
	returns an error if an error has been thrown
*/
func GetAProductByIdFromDB(id int, requestedColumns []string) (Skate, error) {
	rows, err := DB.Query(fmt.Sprintf("SELECT %s FROM skates WHERE id=?", strings.Join(requestedColumns, ",")), id)
	if err != nil {
		return Skate{}, err
	}
	productsFromRows, err := ExtractProductsFromRows(rows)
	if err != nil {
		return Skate{}, err
	}
	if len(productsFromRows) == 0 {
		return Skate{}, nil
	}
	return productsFromRows[0], nil
}

/*
	GetAllProductsFromDB
	return a slice of Skate that has the requested fields filled out
*/
func GetAllProductsFromDB(requestedFields []string) ([]Skate, error) {
	rows, err := DB.Query(fmt.Sprintf("SELECT %s FROM skates WHERE is_active = 'yes' LIMIT 100", strings.Join(requestedFields, ",")))
	if err != nil {
		return nil, err
	}
	productsFromRows, err := ExtractProductsFromRows(rows)
	if err != nil {
		return nil, err
	}
	err = rows.Close()
	if err != nil {
		return []Skate{}, err
	}
	return productsFromRows, nil
}

/*
	UpdateAProductById
	Updates a product with dynamic fields
	Checks `product` for ZeroValue fields, and creates a query that only updates those fields
	Returns the requestedFields of the updated product from the database (after commit)
*/
func UpdateAProductById(product Skate, requestedFields []string) (Skate, error) {

	if product.ID == 0 {
		return Skate{}, errors.New("no ID passed to UpdateAProductById")
	}

	var fieldsToUpdateWithId, _ = StatFields(product)
	fieldsToUpdateWithId = fieldsToUpdateWithId[1:]

	if len(fieldsToUpdateWithId) == 0 {
		rows, err := DB.Query(fmt.Sprintf("SELECT %s from skates WHERE id = ? ", strings.Join(requestedFields, ",")), product.ID)
		if err != nil {
			return Skate{}, err
		}
		productsFromRows, err := ExtractProductsFromRows(rows)
		if err != nil {
			return Skate{}, err
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
	queryString.WriteString(", modified_date = UTC_TIMESTAMP()")
	queryString.WriteString(" WHERE id = ? ")
	var err error
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

		return Skate{}, err
	}
	rows, err := DB.Query(fmt.Sprintf("SELECT %s FROM skates WHERE id = ?", strings.Join(requestedFields, ",")), product.ID)
	if err != nil {
		return Skate{}, err
	}
	productsFromRows, err := ExtractProductsFromRows(rows)
	if err != nil {
		return Skate{}, err
	}
	err = rows.Close()
	if err != nil {
		return Skate{}, err
	}
	if len(productsFromRows) == 0 {
		return Skate{}, nil
	}
	return productsFromRows[0], nil
}

/*
	DeactivateProductById
	Deactivates a skate in the db
*/
func DeactivateProductById(id int) (Skate, error) {
	_, err := DB.Exec("UPDATE skates set is_active = 'no' where id = ?", id)
	if err != nil {
		return Skate{}, err
	}
	rows, err := DB.Query("SELECT id, is_active from skates where id = ?", id)
	if err != nil {
		return Skate{}, err
	}
	productsFromRows, err := ExtractProductsFromRows(rows)
	if err != nil {
		return Skate{}, err
	}
	if len(productsFromRows) == 0 {
		return Skate{}, nil
	}
	err = rows.Close()
	if err != nil {
		return Skate{}, err
	}
	return productsFromRows[0], nil

}

func SeedFromFile() {
	c, ioErr := ioutil.ReadFile("./blades_skates.sql")
	if ioErr != nil {
		log.Println(ioErr)
		os.Exit(1)
	}
	_, err := DB.Query(string(c))
	if err != nil {
		log.Println(err)
	}
}
