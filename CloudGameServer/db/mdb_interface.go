package db

type Use interface {
	Save() error
	Delete() error
	Update() error
}
