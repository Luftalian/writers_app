package usecase

import (
	"reflect"
	"testing"

	"github.com/Luftalian/writers_app/domain"
)

func TestTagInteractor_Tags(t *testing.T) {
	type fields struct {
		TagRepository TagRepository
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
			interactor := &TagInteractor{
				TagRepository: tt.fields.TagRepository,
			}
			got, err := interactor.Tags()
			if (err != nil) != tt.wantErr {
				t.Errorf("TagInteractor.Tags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagInteractor.Tags() = %v, want %v", got, tt.want)
			}
		})
	}
}
