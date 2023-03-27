package usecase

import (
	"reflect"
	"testing"

	"github.com/Luftalian/writers_app/domain"
)

func TestUserLikeInteractor_Add(t *testing.T) {
	type fields struct {
		UserLikeRepository UserLikeRepository
	}
	type args struct {
		u domain.UserLike
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.UserLike
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interactor := &UserLikeInteractor{
				UserLikeRepository: tt.fields.UserLikeRepository,
			}
			got, err := interactor.Add(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserLikeInteractor.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserLikeInteractor.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
