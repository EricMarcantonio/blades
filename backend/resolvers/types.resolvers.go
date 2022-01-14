package resolvers

import (
	"backend/types"
	"github.com/graphql-go/graphql"
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"product": &graphql.Field{
				Type:        types.GProduct,
				Description: "Get product by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: GetProductById,
			},
			"products": &graphql.Field{
				Type:        graphql.NewList(types.GProduct),
				Description: "Get product list",
				Resolve:     GetAllProducts,
			},
		},
	})

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createProduct": &graphql.Field{
			Type:        types.GProduct,
			Description: "Create new product",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"price": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Float),
				},
				"units": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 0,
				},
			},
			Resolve: CreateProduct,
		},
		"updateProduct": &graphql.Field{
			Type:        types.GProduct,
			Description: "Update a Product",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"name": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},
				"price": &graphql.ArgumentConfig{
					Type:         graphql.Float,
					DefaultValue: 0.0,
				},
				"is_active": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},
				"units": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: -1, //no value was passed if -1, so we can update with a 0
				},
			},
			Resolve: UpdateAProduct,
		},
		"deactivateProduct": &graphql.Field{
			Type:        types.GProduct,
			Description: "Update a Product",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: DeactivateProduct,
		},
	},
})
