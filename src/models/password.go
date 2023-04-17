package models

// Senha representa o formato da requisição de atualização de senha
type Password struct {
	New string `json:"new"`
	Current string `json:"current"`
}