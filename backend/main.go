package main

import (
	"backend/database"
	"backend/resolvers"
	"database/sql"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/joho/sqltocsv"
	"log"
	"net/http"
	"os"
)

func main() {
	var err error
	err, database.DB = database.ConnectToDatabase()

	if err != nil {
		os.Exit(1)
	}

	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    resolvers.QueryType,
			Mutation: resolvers.MutationType,
		},
	)

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     false,
		GraphiQL:   true,
		Playground: false,
	})

	http.HandleFunc("/gql", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		h.ServeHTTP(w, r)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		keys, ok := r.URL.Query()["seed"]
		if !ok {
			_, err := fmt.Fprintf(w, "Missing params!")
			if err != nil {
				return
			}
		} else {
			if keys[0] == "yes" {
				database.SeedFromFile()
			}
		}
	})
	http.HandleFunc("/csv", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		rows, err := database.DB.Query("SELECT * FROM skates LIMIT 10000")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		w.Header().Set("Content-type", "text/csv")
		w.Header().Set("Content-Disposition", "attachment; filename=\"report.csv\"")

		err = sqltocsv.Write(w, rows)
		if err != nil {
			return
		}
	})

	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalln(err)
	}

	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatalln("couldn't close connection the DB")
		}
	}(database.DB)
}
