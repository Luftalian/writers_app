package controller

import (
	"reflect"
	"testing"

	"github.com/Luftalian/writers_app/interfaces/database"
)

func TestNewTagListController(t *testing.T) {
	type args struct {
		handler database.DbHandler
	}
	tests := []struct {
		name string
		args args
		want *TagListController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTagListController(tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTagListController() = %v, want %v", got, tt.want)
			}
		})
	}
}
