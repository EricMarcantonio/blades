package database

import (
	"backend/types"
	"database/sql"
	"fmt"
	"reflect"
	"testing"
)

func connectToDatabase() {
	var err error
	err, DB = ConnectToDatabase()
	if err != nil {
		return
	}
}

func TestBuildParamUpdateColumn(t *testing.T) {
	type args struct {
		changingColumns []string
	}
	var tests = []struct {
		name string
		args args
		want string
	}{
		{
			name: "Real columns",
			args: args{changingColumns: []string{"id", "name", "price", "modified_date", "added_date", "is_active", "units"}},
			want: "name = ?,price = ?,is_active = ?,units = ?",
		},
		{
			name: "No correct columns",
			args: args{changingColumns: []string{"test"}},
			want: "",
		},
		{
			name: "Empty string",
			args: args{changingColumns: []string{""}},
			want: "",
		},
		{
			name: "One column, shouldn't be returned",
			args: args{changingColumns: []string{"modified_date"}},
			want: "",
		},
		{
			name: "One column, should be returned",
			args: args{changingColumns: []string{"name"}},
			want: "name = ?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildParamUpdateColumn(tt.args.changingColumns); got != tt.want {
				t.Errorf("BuildParamUpdateColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateProduct(t *testing.T) {
	connectToDatabase()
	type args struct {
		product          types.Product
		requestedColumns []string
	}
	tests := []struct {
		name    string
		args    args
		want    types.Product
		wantErr bool
	}{
		{
			name: "Create a product",
			args: args{
				product: types.Product{
					Name:  "Skates",
					Price: 10.01,
					Units: 1,
				},
				requestedColumns: []string{"id", "name"},
			},
			want: types.Product{
				ID:   5,
				Name: "Skates",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateProduct(tt.args.product, tt.args.requestedColumns)
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

func TestExtractProductsFromRows(t *testing.T) {
	type args struct {
		rows *sql.Rows
	}
	connectToDatabase()
	rows, err := DB.Query("SELECT * from skates order by id asc limit 1")
	col, _ := rows.Columns()
	fmt.Println(col)
	if err != nil {
		return
	}
	tests := []struct {
		name    string
		args    args
		want    []types.Product
		wantErr bool
	}{
		{
			name: "Get a product struct from a row in the db",
			args: args{rows: rows},
			want: []types.Product{
				{
					ID:       1,
					Name:     "Bauer Supreme Ultrasonic Skates",
					Price:    599.99,
					IsActive: "yes",
					Units:    12,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractProductsFromRows(tt.args.rows)
			got[0].ModifiedDate = ""
			got[0].AddedDate = ""
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractProductsFromRows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractProductsFromRows() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAProductById(t *testing.T) {
	type args struct {
		id               int
		requestedColumns []string
	}
	connectToDatabase()
	tests := []struct {
		name    string
		args    args
		want    types.Product
		wantErr bool
	}{
		{
			name: "ID passed",
			args: args{
				id:               1,
				requestedColumns: []string{"id", "name", "price", "is_active", "units"},
			},

			want: types.Product{
				ID:       1,
				Name:     "Bauer Supreme Ultrasonic Skates",
				Price:    599.99,
				IsActive: "yes",
				Units:    12,
			},
		},
		{
			name: "Wrong ID passed",
			args: args{
				id:               0,
				requestedColumns: []string{"id", "name", "price", "is_active", "units"},
			},

			want: types.Product{
				ID:           0,
				Name:         "",
				Price:        0,
				ModifiedDate: "",
				AddedDate:    "",
				IsActive:     "",
				Units:        0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAProductById(tt.args.id, tt.args.requestedColumns)
			got.ModifiedDate = ""
			got.AddedDate = ""
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAProductById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAProductById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllProducts(t *testing.T) {
	connectToDatabase()
	type args struct {
		requestedFields []string
	}
	tests := []struct {
		name    string
		args    args
		want    []types.Product
		wantErr bool
	}{
		{
			name: "Whole Database",
			args: args{requestedFields: []string{"id"}},
			want: []types.Product{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
				{
					ID: 3,
				},
				{
					ID: 4,
				},
				{
					ID: 5,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllProducts(tt.args.requestedFields)
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

func TestUpdateAProductById(t *testing.T) {
	connectToDatabase()
	type args struct {
		product         types.Product
		requestedFields []string
	}
	tests := []struct {
		name    string
		args    args
		want    types.Product
		wantErr bool
	}{
		{
			name: "Change the number of units for an ID",
			args: args{
				product:         types.Product{ID: 1, Units: 1200},
				requestedFields: []string{"id", "units"},
			},
			want: types.Product{
				ID:    1,
				Units: 1200,
			},
		},
		{
			name: "Change the number of units for an ID",
			args: args{
				product:         types.Product{ID: 1, Units: 12},
				requestedFields: []string{"id", "units"},
			},
			want: types.Product{
				ID:    1,
				Units: 12,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UpdateAProductById(tt.args.product, tt.args.requestedFields)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateAProductById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateAProductById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeactivateProductById(t *testing.T) {
	connectToDatabase()
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    types.Product
		wantErr bool
	}{
		{
			name:    "Deactivate number 5",
			args:    args{id: 5},
			want:    types.Product{ID: 5, IsActive: "no"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeactivateProductById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeactivateProductById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeactivateProductById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
