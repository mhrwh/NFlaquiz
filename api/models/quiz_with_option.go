package models

type QuizWithOption struct {
	ID          int
	CountryName string
	CountryID   string
	Hints       []string
	Options     []string
}
