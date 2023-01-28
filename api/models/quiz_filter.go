package models

type QuizFilter struct {
	Colors   []string `form:"colors[]"`
	Areas    []string `form:"areas[]"`
	Bookmark string   `form:"bookmark"`
}
