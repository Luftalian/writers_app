package domain

type Tag struct {
	TagID   UUID   `json:"tag_id" db:"tag_id"`
	TagName string `json:"tag_name" db:"tag_name"`
}

type Tags []Tag

type TagList struct {
	TagID   UUID   `json:"tag_id" db:"tag_id"`
	TagName string `json:"tag_name" db:"name"`
	TextID  UUID   `json:"text_id" db:"text_id"`
}

type TagLists []TagList
