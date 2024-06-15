package models

type User struct {
	ID 	int16
	Email string
	PassHash []byte
}