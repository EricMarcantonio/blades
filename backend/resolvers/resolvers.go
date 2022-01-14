package resolvers

import (
	"backend/database"
	"backend/types"
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

func GetProductById(p graphql.ResolveParams) (interface{}, error) {
	i := p.Args["id"].(int)
	if i == 0 {
		return nil, errors.New("no id passed, or wrong type")
	}
	product, err := database.GetAProductById(i, GetSelectedFields([]string{"product"}, p))
	if err != nil {
		return nil, err
	}
	return product, nil
}

func GetAllProducts(p graphql.ResolveParams) (interface{}, error) {
	products, err := database.GetAllProducts(GetSelectedFields([]string{"products"}, p))
	if err != nil {
		return nil, err
	}
	return products, nil
}

func CreateProduct(p graphql.ResolveParams) (interface{}, error) {
	if p.Args["name"].(string) == "" {
		return types.Product{}, errors.New("no name passed")
	}
	if p.Args["price"].(float64) == 0.0 {
		return types.Product{}, errors.New("no price passed")
	}
	product := types.Product{
		Name:  p.Args["name"].(string),
		Price: p.Args["price"].(float64),
	}
	product, err := database.CreateProduct(product, GetSelectedFields([]string{"createProduct"}, p))
	if err != nil {
		return types.Product{}, err
	}
	return product, nil
}

func UpdateAProduct(p graphql.ResolveParams) (interface{}, error) {
	i := p.Args["id"].(int)
	if i == 0 {
		return types.Product{}, errors.New("ID not passed to UpdateAProduct")
	}
	product := types.Product{
		ID:       p.Args["id"].(int),
		Name:     p.Args["name"].(string),
		Price:    p.Args["price"].(float64),
		IsActive: p.Args["is_active"].(string),
		Units:    p.Args["units"].(int),
	}

	aProductById, err := database.UpdateAProductById(product, GetSelectedFields([]string{"updateProduct"}, p))
	if err != nil {
		return nil, err
	}
	return aProductById, nil
}

func GetSelectedFields(selectionPath []string, resolveParams graphql.ResolveParams) []string {
	fields := resolveParams.Info.FieldASTs
	for _, propName := range selectionPath {
		found := false
		for _, field := range fields {
			if field.Name.Value == propName {
				selections := field.SelectionSet.Selections
				fields = make([]*ast.Field, 0)
				for _, selection := range selections {
					fields = append(fields, selection.(*ast.Field))
				}
				found = true
				break
			}
		}
		if !found {
			return []string{}
		}
	}
	var collect []string
	for _, field := range fields {
		collect = append(collect, field.Name.Value)
	}
	return collect
}

func DeactivateProduct(p graphql.ResolveParams) (interface{}, error) {
	i := p.Args["id"].(int)
	if i == 0 {
		return types.Product{}, errors.New("ID not passed to Deactivate")
	}

	aProductById, err := database.DeactivateProductById(i)
	if err != nil {
		return nil, err
	}
	return aProductById, nil
}
