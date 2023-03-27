package usecase

import (
	"reflect"
	"testing"

	"github.com/Luftalian/writers_app/domain"
)

func TestTextInteractor_Add(t *testing.T) {
	type fields struct {
		TextRepository TextRepository
	}
	type args struct {
		t domain.Text
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Text
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &TextInteractor{
				TextRepository: tt.fields.TextRepository,
			}
			got, err := interactor.Add(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("TextInteractor.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TextInteractor.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
