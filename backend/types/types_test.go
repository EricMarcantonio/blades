package types

import (
	"reflect"
	"testing"
)

func TestStatFields(t *testing.T) {
	type args struct {
		object interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := StatFields(tt.args.object)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StatFields() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("StatFields() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
