package database

import (
	"reflect"
	"testing"

	"github.com/Luftalian/writers_app/domain"
	"github.com/google/uuid"
)

func TestTextRepository_FindAll(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	tests := []struct {
		name    string
		fields  fields
		want    domain.Texts
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &TextRepository{
				DbHandler: tt.fields.DbHandler,
			}
			got, err := repo.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("TextRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TextRepository.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTextRepository_FindByID(t *testing.T) {
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
		want    domain.Text
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &TextRepository{
				DbHandler: tt.fields.DbHandler,
			}
			got, err := repo.FindByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("TextRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TextRepository.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTextRepository_Store(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	type args struct {
		text domain.Text
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
			repo := &TextRepository{
				DbHandler: tt.fields.DbHandler,
			}
			got, err := repo.Store(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("TextRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TextRepository.Store() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTextRepository_Update(t *testing.T) {
	type fields struct {
		DbHandler DbHandler
	}
	type args struct {
		text domain.Text
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
			repo := &TextRepository{
				DbHandler: tt.fields.DbHandler,
			}
			if err := repo.Update(tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("TextRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTextRepository_DeleteByID(t *testing.T) {
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
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &TextRepository{
				DbHandler: tt.fields.DbHandler,
			}
			if err := repo.DeleteByID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("TextRepository.DeleteByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
