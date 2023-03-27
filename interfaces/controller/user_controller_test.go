package controller

import (
	"reflect"
	"testing"

	"github.com/Luftalian/writers_app/interfaces/database"
)

func TestNewUserController(t *testing.T) {
	type args struct {
		handler database.DbHandler
	}
	tests := []struct {
		name string
		args args
		want *UserController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserController(tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserController() = %v, want %v", got, tt.want)
			}
		})
	}
}
