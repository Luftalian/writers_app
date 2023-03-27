package database

import (
	"reflect"
	"testing"

	"github.com/Luftalian/writers_app/domain"
	"github.com/google/uuid"
)

func TestUserCreateRepository_FindAll(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	tests := []struct {
		name    string
		fields  fields
		want    domain.UserCreates
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &UserCreateRepository{
				DbHandler: tt.fields.DbHandler,
			}
			got, err := repo.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("UserCreateRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserCreateRepository.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserCreateRepository_FindByUserID(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	type args struct {
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Texts
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &UserCreateRepository{
				DbHandler: tt.fields.DbHandler,
			}
			got, err := repo.FindByUserID(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserCreateRepository.FindByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserCreateRepository.FindByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserCreateRepository_FindByTextID(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	type args struct {
		textID uuid.UUID
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
			repo := &UserCreateRepository{
				DbHandler: tt.fields.DbHandler,
			}
			got, err := repo.FindByTextID(tt.args.textID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserCreateRepository.FindByTextID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserCreateRepository.FindByTextID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserCreateRepository_Store(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	type args struct {
		userCreate domain.UserCreate
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
			repo := &UserCreateRepository{
				DbHandler: tt.fields.DbHandler,
			}
			got, err := repo.Store(tt.args.userCreate)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserCreateRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserCreateRepository.Store() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserCreateRepository_Update(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	type args struct {
		userCreate domain.UserCreate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &UserCreateRepository{
				DbHandler: tt.fields.DbHandler,
			}
			if err := repo.Update(tt.args.userCreate); (err != nil) != tt.wantErr {
				t.Errorf("UserCreateRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserCreateRepository_Delete(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	type args struct {
		u domain.UserCreate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &UserCreateRepository{
				DbHandler: tt.fields.DbHandler,
			}
			if err := repo.Delete(tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("UserCreateRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserCreateRepository_DeleteByUserID(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	type args struct {
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &UserCreateRepository{
				DbHandler: tt.fields.DbHandler,
			}
			if err := repo.DeleteByUserID(tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("UserCreateRepository.DeleteByUserID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserCreateRepository_DeleteByTextID(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	type args struct {
		textID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &UserCreateRepository{
				DbHandler: tt.fields.DbHandler,
			}
			if err := repo.DeleteByTextID(tt.args.textID); (err != nil) != tt.wantErr {
				t.Errorf("UserCreateRepository.DeleteByTextID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
