package main

import (
	"backend/database"
	"backend/resolvers"
	"database/sql"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
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

	http.HandleFunc("/gql", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "*")
		h.ServeHTTP(writer, request)
	})
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		keys, ok := request.URL.Query()["seed"]
		if !ok {
			_, err := fmt.Fprintf(writer, "Missing params!")
			if err != nil {
				return
			}
		} else {
			if keys[0] == "yes" {
				database.SeedFromFile()
			}
		}

	})
	//http.Handle("/graphql", h)
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
