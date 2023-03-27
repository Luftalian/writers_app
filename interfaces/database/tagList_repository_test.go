package database

import (
	"reflect"
	"testing"

	"github.com/Luftalian/writers_app/domain"
	"github.com/google/uuid"
)

func TestTagListRepository_FindAll(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	tests := []struct {
		name    string
		fields  fields
		want    domain.TagLists
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &TagListRepository{
				DbHandler: tt.fields.DbHandler,
			}
			got, err := repo.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("TagListRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagListRepository.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTagListRepository_Store(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	type args struct {
		tagList domain.TagList
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.TagList
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &TagListRepository{
				DbHandler: tt.fields.DbHandler,
			}
			got, err := repo.Store(tt.args.tagList)
			if (err != nil) != tt.wantErr {
				t.Errorf("TagListRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagListRepository.Store() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTagListRepository_FindByTextID(t *testing.T) {
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
		want    domain.TagLists
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &TagListRepository{
				DbHandler: tt.fields.DbHandler,
			}
			got, err := repo.FindByTextID(tt.args.textID)
			if (err != nil) != tt.wantErr {
				t.Errorf("TagListRepository.FindByTextID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagListRepository.FindByTextID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTagListRepository_FindByTagID(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	type args struct {
		tagID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.TagLists
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &TagListRepository{
				DbHandler: tt.fields.DbHandler,
			}
			got, err := repo.FindByTagID(tt.args.tagID)
			if (err != nil) != tt.wantErr {
				t.Errorf("TagListRepository.FindByTagID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagListRepository.FindByTagID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTagListRepository_FindByName(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	type args struct {
		name string
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
			repo := &TagListRepository{
				DbHandler: tt.fields.DbHandler,
			}
			got, err := repo.FindByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("TagListRepository.FindByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagListRepository.FindByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTagListRepository_Update(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	type args struct {
		tagList domain.TagList
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.TagList
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &TagListRepository{
				DbHandler: tt.fields.DbHandler,
			}
			got, err := repo.Update(tt.args.tagList)
			if (err != nil) != tt.wantErr {
				t.Errorf("TagListRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagListRepository.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTagListRepository_Delete(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	type args struct {
		tagID uuid.UUID
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
			repo := &TagListRepository{
				DbHandler: tt.fields.DbHandler,
			}
			if err := repo.Delete(tt.args.tagID); (err != nil) != tt.wantErr {
				t.Errorf("TagListRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
