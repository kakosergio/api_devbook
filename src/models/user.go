package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User representa um usuário e suas informações
type User struct {
	Id        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedOn time.Time `json:"created_on,omitempty"`
}

// Prepare chamará os metodos para validar e formatar os dados do usuário recebido
func (user *User) Prepare(stage string) error {
	if err := user.validate(stage); err != nil {
		return err
	}
	user.format()
	return nil
}

func (user *User) validate(stage string) error {
	// inseriu o stage por causa do metodo de atualização do usuario, quando não for alterar a senha, só nome, nick e email
	if user.Name == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}
	if user.Nick == "" {
		return errors.New("o nick é obrigatório e não pode estar em branco")
	}
	if user.Email == ""{
		return errors.New("o e-mail é obrigatório e não pode estar em branco")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil{
		return errors.New("o email inserido é inválido")
	}

	if stage == "signup" && user.Password == "" {
		return errors.New("a senha é obrigatória e não pode estar em branco")
	}
	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
