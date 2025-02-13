package model

type Goal struct {
	ID    int
	Goal  string
	UserID int
	Chapters []string
}