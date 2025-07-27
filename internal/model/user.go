package model

type RegisterInput struct {
	Name         string
	Avatar       string
	Password     string
	UserSalt     string
	Sex          int
	Status       int
	Sign         string
	SecretAnswer string
}

type RegisterOutput struct {
	Id uint
}

type LoginInput struct {
	Name     string
	Password string
}

type UpdatePasswordInput struct {
	Password     string
	UserSalt     string
	SecretAnswer string
}

type UpdatePasswordOutput struct {
	Id uint
}
