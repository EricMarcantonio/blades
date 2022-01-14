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

func ConnectToDatabase() (error, *sql.DB) {
	_, ok := os.LookupEnv("user")
	if !ok {
		SetEnv("user", "backend")
	}
	_, ok = os.LookupEnv("db")
	if !ok {
		SetEnv("db", "blades")
	}
	_, ok = os.LookupEnv("pass")
	if !ok {
		SetEnv("pass", "password")
	}
	_, ok = os.LookupEnv("root_pass")
	if !ok {
		SetEnv("root_pass", "example")
	}
	_, ok = os.LookupEnv("port")
	if !ok {
		SetEnv("port", "3306")
	}
	_, ok = os.LookupEnv("domain")
	if !ok {
		SetEnv("domain", "127.0.0.1")
	}

	SetEnv("domain", "127.0.0.1")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", os.Getenv("user"), os.Getenv("pass"), os.Getenv("domain"), os.Getenv("port"), os.Getenv("db")))
	if err != nil {
		return err, nil
	}
	return nil, db
}

func SetEnv(key string, value string) {
	_, exists := os.LookupEnv(key)
	if !exists {
		err := os.Setenv(key, value)
		if err != nil {
			os.Exit(1)
		}
	}
}

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
		panic("Not impletmented")
	}
}

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

func CreateProductInDB(product Skate, requestedColumns []string) (Skate, error) {
	_, err := DB.Exec("INSERT INTO skates (name, price, modified_date, added_date, is_active, units) VALUES (?, ?, NOW(), NOW(),'yes', ?)", product.Name, product.Price, product.Units)
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
	return productsFromRows[0], nil
}
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

func UpdateAProductById(product Skate, requestedFields []string) (Skate, error) {
	// Check if an ID was given
	if product.ID == 0 {
		return Skate{}, errors.New("no ID passed to UpdateAProductById")
	}

	//Build the query string with this
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
	queryString.WriteString(", modified_date = NOW()")
	queryString.WriteString(" WHERE id = ? ")
	var err error
	//var tempProduct types.Skate
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
	return productsFromRows[0], nil
}

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
