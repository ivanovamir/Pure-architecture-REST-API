package dto

type User struct {
	Id           int
	Name         string
	Surname      string
	Patronymic   string
	Email        string
	Phone        string
	Password     string
	PasswordHash []byte
}
