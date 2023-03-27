package controller

import (
	"reflect"
	"testing"

	"github.com/Luftalian/writers_app/interfaces/database"
)

func TestNewTextController(t *testing.T) {
	type args struct {
		handler database.DbHandler
	}
	tests := []struct {
		name string
		args args
		want *TextController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTextController(tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTextController() = %v, want %v", got, tt.want)
			}
		})
	}
}
