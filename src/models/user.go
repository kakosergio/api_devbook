package models

import "time"

// User representa um usuário e suas informações
type User struct {
	Id   uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Nick string `json:"nick,omitempty"`
	Email string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	CreatedOn time.Time `json:"created_on,omitempty"`
}
