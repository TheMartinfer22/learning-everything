package main

import (
	"ExceptionHandlerProject/exception"
	"fmt"
)

func main() {
	user, err := GetUser()
	if err != nil {
		fmt.Printf("%#v", err)
	}
	println(user.Nome)
}

func GetUser() (user User, err error) {
	user.Nome = ""
	user.Email = "fulano@example.com"
	if user.Nome == "" {
		err = exception.Exception{Code: 400, Message: "Campo Nome não pode estar vazio"}
		return
	}
	return
}

type User struct {
	Nome  string
	Email string
}
