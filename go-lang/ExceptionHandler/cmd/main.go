package main

import "fmt"

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
		err = Exception{Code: 400, Message: "Campo Nome não pode estar vazio"}
		return
	}

	return
}

type User struct {
	Nome  string
	Email string
}

type Exception struct {
	Code    int
	Message string
}

func (e Exception) Error() string {
	return fmt.Sprintf("Erro ao processar dados, code=%d, message=%s",
		e.Code,
		e.Message,
	)
}
