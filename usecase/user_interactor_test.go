package usecase

import (
	"reflect"
	"testing"

	"github.com/Luftalian/writers_app/domain"
)

func TestUserInteractor_Add(t *testing.T) {
	type fields struct {
		UserRepository UserRepository
	}
	type args struct {
		u domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &UserInteractor{
				UserRepository: tt.fields.UserRepository,
			}
			got, err := interactor.Add(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserInteractor.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserInteractor.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
