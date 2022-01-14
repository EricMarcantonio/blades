package resolvers

import (
	"github.com/graphql-go/graphql"
	"reflect"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	type args struct {
		p graphql.ResolveParams
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateProduct(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateProduct() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllProducts(t *testing.T) {
	type args struct {
		p graphql.ResolveParams
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllProducts(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllProducts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetProductById(t *testing.T) {
	type args struct {
		p graphql.ResolveParams
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetProductById(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProductById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProductById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSelectedFields(t *testing.T) {
	type args struct {
		selectionPath []string
		resolveParams graphql.ResolveParams
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSelectedFields(tt.args.selectionPath, tt.args.resolveParams); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSelectedFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateAProduct(t *testing.T) {
	type args struct {
		p graphql.ResolveParams
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UpdateAProduct(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateAProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateAProduct() got = %v, want %v", got, tt.want)
			}
		})
	}
}
