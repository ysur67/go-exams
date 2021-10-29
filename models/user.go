package models

type User struct {
	Id       string
	Username string
	Password string
}

type LoginParam struct {
	Username string
	Password string
}
