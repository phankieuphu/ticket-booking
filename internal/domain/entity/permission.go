package entity

type Permission struct {
	ID          int
	Name        string
	Description string
	Resource    string
	Action      string // enum
}
