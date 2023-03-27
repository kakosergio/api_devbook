package models

import (
	"errors"
	"strings"
	"time"
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
func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}
	user.format()
	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}
	if user.Nick == "" {
		return errors.New("o nick é obrigatório e não pode estar em branco")
	}
	if user.Email == "" {
		return errors.New("o e-mail é obrigatório e não pode estar em branco")
	}
	if user.Password == "" {
		return errors.New("a senha é obrigatória e não pode estar em branco")
	}
	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
