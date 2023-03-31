package security

import "golang.org/x/crypto/bcrypt"

// Hash cria um hash da senha fornecida
func Hash(pwd string)([]byte, error){
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
}

// VerifyPassword compara o hash com a senha
func VerifyPassword (pwdHashed, pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(pwdHashed), []byte(pwd))
}