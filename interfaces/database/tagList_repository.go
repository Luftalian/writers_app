package database

import (
	"github.com/google/uuid"

	"github.com/Luftalian/writers_app/domain"
)

type TagListRepository struct {
	DbHandler
}

func (repo *TagListRepository) FindAll() (domain.TagLists, error) {
	tagLists := domain.TagLists{}
	err := repo.Select(&tagLists, "SELECT * FROM tagLists")
	if err != nil {
		return nil, err
	}
	return tagLists, nil
}

func (repo *TagListRepository) Store(tagList domain.TagList) (domain.TagList, error) {
	tagList.TagID = uuid.New()
	_, err := repo.Exec("INSERT INTO tagLists (tag_id, name, text_id) VALUES (?, ?, ?)", tagList.TagID, tagList.TagName, tagList.TextID)
	if err != nil {
		return domain.TagList{}, err
	}
	return tagList, nil
}

func (repo *TagListRepository) FindByTextID(textID uuid.UUID) (domain.TagList, error) {
	tagList := domain.TagList{}
	err := repo.Select(&tagList, "SELECT * FROM tagLists WHERE text_id = ?", textID)
	if err != nil {
		return domain.TagList{}, err
	}
	return tagList, nil
}

func (repo *TagListRepository) FindByTagID(tagID uuid.UUID) (domain.TagLists, error) {
	tagLists := domain.TagLists{}
	err := repo.Select(&tagLists, "SELECT * FROM tagLists WHERE tag_id = ?", tagID)
	if err != nil {
		return nil, err
	}
	return tagLists, nil
}

func (repo *TagListRepository) FindByName(name string) (domain.Texts, error) {
	texts := domain.Texts{}
	err := repo.Select(&texts, "SELECT texts.text_id, texts.title, texts.content, texts.user_name, texts.user_id, texts.created_at, texts.changed_at, texts.good_count, texts.bad_count FROM texts INNER JOIN tagLists ON texts.text_id = tagLists.text_id WHERE tagLists.name = ?", name)
	if err != nil {
		return nil, err
	}
	return texts, nil
}

func (repo *TagListRepository) Update(tagList domain.TagList) (domain.TagList, error) {
	_, err := repo.Exec("UPDATE tagLists SET name = ? WHERE tag_id = ?", tagList.TagName, tagList.TagID)
	if err != nil {
		return domain.TagList{}, err
	}
	return tagList, nil
}

func (repo *TagListRepository) Delete(tagID uuid.UUID) error {
	_, err := repo.Exec("DELETE FROM tagLists WHERE tag_id = ?", tagID)
	if err != nil {
		return err
	}
	return nil
}
