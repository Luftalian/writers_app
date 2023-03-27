package usecase

import (
	"reflect"
	"testing"

	"github.com/Luftalian/writers_app/domain"
)

func TestUserCreateInteractor_Add(t *testing.T) {
	type fields struct {
		UserCreateRepository UserCreateRepository
	}
	type args struct {
		u domain.UserCreate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.UserCreate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &UserCreateInteractor{
				UserCreateRepository: tt.fields.UserCreateRepository,
			}
			got, err := interactor.Add(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserCreateInteractor.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserCreateInteractor.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
