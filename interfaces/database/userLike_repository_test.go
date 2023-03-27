package database

import (
	"reflect"
	"testing"

	"github.com/Luftalian/writers_app/domain"
	"github.com/google/uuid"
)

func TestUserLikeRepository_FindAll(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	tests := []struct {
		name    string
		fields  fields
		want    domain.UserLikes
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &UserLikeRepository{
				DbHandler: tt.fields.DbHandler,
			}
			got, err := repo.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("UserLikeRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserLikeRepository.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserLikeRepository_FindByUserID(t *testing.T) {
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
			repo := &UserLikeRepository{
				DbHandler: tt.fields.DbHandler,
			}
			got, err := repo.FindByUserID(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserLikeRepository.FindByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserLikeRepository.FindByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserLikeRepository_FindByTextID(t *testing.T) {
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
		want    domain.UserLike
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &UserLikeRepository{
				DbHandler: tt.fields.DbHandler,
			}
			got, err := repo.FindByTextID(tt.args.textID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserLikeRepository.FindByTextID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserLikeRepository.FindByTextID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserLikeRepository_Store(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	type args struct {
		userLike domain.UserLike
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
			repo := &UserLikeRepository{
				DbHandler: tt.fields.DbHandler,
			}
			got, err := repo.Store(tt.args.userLike)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserLikeRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserLikeRepository.Store() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserLikeRepository_CheckLikeUser(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	type args struct {
		userLike domain.UserLike
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &UserLikeRepository{
				DbHandler: tt.fields.DbHandler,
			}
			got, err := repo.CheckLikeUser(tt.args.userLike)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserLikeRepository.CheckLikeUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UserLikeRepository.CheckLikeUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserLikeRepository_Update(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	type args struct {
		userLike domain.UserLike
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
			repo := &UserLikeRepository{
				DbHandler: tt.fields.DbHandler,
			}
			if err := repo.Update(tt.args.userLike); (err != nil) != tt.wantErr {
				t.Errorf("UserLikeRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserLikeRepository_Delete(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	type args struct {
		u domain.UserLike
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
			repo := &UserLikeRepository{
				DbHandler: tt.fields.DbHandler,
			}
			if err := repo.Delete(tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("UserLikeRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserLikeRepository_DeleteByUserID(t *testing.T) {
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
			repo := &UserLikeRepository{
				DbHandler: tt.fields.DbHandler,
			}
			if err := repo.DeleteByUserID(tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("UserLikeRepository.DeleteByUserID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserLikeRepository_DeleteByTextID(t *testing.T) {
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
			repo := &UserLikeRepository{
				DbHandler: tt.fields.DbHandler,
			}
			if err := repo.DeleteByTextID(tt.args.textID); (err != nil) != tt.wantErr {
				t.Errorf("UserLikeRepository.DeleteByTextID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
