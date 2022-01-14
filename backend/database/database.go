package database

import (
	"backend/types"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
	SetEnv("user", "backend")
	SetEnv("db", "blades")
	SetEnv("pass", "password")
	SetEnv("root_pass", "example")
	SetEnv("port", "3306")
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

// ExtractProductsFromRows
// Take the result of a query and return an array of products
// /**
func ExtractProductsFromRows(rows *sql.Rows) ([]types.Product, error) {
	var tempProducts []types.Product
	var colNames []string
	var err error
	colNames, err = rows.Columns()
	if err != nil {
		panic(err)
	}
	//fmt.Println(colNames)
	for rows.Next() {
		var tempProduct types.Product

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
	//log.Println(tempProducts)
	return tempProducts, nil
}

func ProductCol(ColName string, product *types.Product) interface{} {
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
