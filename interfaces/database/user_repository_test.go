package database

import (
	"reflect"
	"testing"

	"github.com/Luftalian/writers_app/domain"
	"github.com/google/uuid"
)

func TestUserRepository_FindAll(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	tests := []struct {
		name    string
		fields  fields
		want    domain.Users
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &UserRepository{
				DbHandler: tt.fields.DbHandler,
			}
			got, err := repo.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_FindByID(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	type args struct {
		id uuid.UUID
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
			repo := &UserRepository{
				DbHandler: tt.fields.DbHandler,
			}
			got, err := repo.FindByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_Store(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	type args struct {
		user domain.User
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
			repo := &UserRepository{
				DbHandler: tt.fields.DbHandler,
			}
			got, err := repo.Store(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.Store() = %v, want %v", got, tt.want)
			}
		})
	}
}
