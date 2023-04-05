package auth

import (
	"api/src/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CreateToken cria e retorna um token assinado com as permissões do usuário
func CreateToken (userID uint64) (string, error){
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))
}