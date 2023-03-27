package database

import (
	"reflect"
	"testing"

	"github.com/Luftalian/writers_app/domain"
	"github.com/google/uuid"
)

func TestTagRepository_FindAll(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	tests := []struct {
		name    string
		fields  fields
		want    domain.Tags
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &TagRepository{
				DbHandler: tt.fields.DbHandler,
			}
			got, err := repo.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("TagRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagRepository.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTagRepository_FindByID(t *testing.T) {
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
		want    domain.Tag
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &TagRepository{
				DbHandler: tt.fields.DbHandler,
			}
			got, err := repo.FindByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("TagRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagRepository.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTagRepository_Store(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	type args struct {
		tag domain.Tag
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Tag
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &TagRepository{
				DbHandler: tt.fields.DbHandler,
			}
			got, err := repo.Store(tt.args.tag)
			if (err != nil) != tt.wantErr {
				t.Errorf("TagRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagRepository.Store() = %v, want %v", got, tt.want)
			}
		})
	}
}
