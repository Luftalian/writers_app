package usecase

import (
	"reflect"
	"testing"

	"github.com/Luftalian/writers_app/domain"
)

func TestTagListInteractor_TagLists(t *testing.T) {
	type fields struct {
		TagListRepository TagListRepository
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
			interactor := &TagListInteractor{
				TagListRepository: tt.fields.TagListRepository,
			}
			got, err := interactor.TagLists()
			if (err != nil) != tt.wantErr {
				t.Errorf("TagListInteractor.TagLists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagListInteractor.TagLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
