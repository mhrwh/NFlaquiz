package models

type QuizWithOption struct {
	ID          int
	CountryName string
	CountryID   string
	Hint1       string
	Hint2       string
	Hint3       string
	Options     []string
}