package controller

import (
	"reflect"
	"testing"

	"github.com/Luftalian/writers_app/interfaces/database"
)

func TestNewUserLikeController(t *testing.T) {
	type args struct {
		handler database.DbHandler
	}
	tests := []struct {
		name string
		args args
		want *UserLikeController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserLikeController(tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserLikeController() = %v, want %v", got, tt.want)
			}
		})
	}
}
