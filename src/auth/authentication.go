package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

func ValidateToken (r *http.Request) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, getVerificationKey)

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
		return nil
	}
	return errors.New("invalid token")
}

func extractToken (r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

// ExtractIdFromToken extrai o userId do token para ser utilizado em algumas validações de permissões.
func ExtractIdFromToken(r *http.Request) (uint64, error){
	// Pega a tokenString e extrai da request
	tokenString := extractToken(r)
	// Faz a verificação do token pra saber se ele é válido
	token, err := jwt.Parse(tokenString, getVerificationKey)
	if err != nil {
		return 0, err
	}
	// Se for válido, verifica se o token tem os Claims, salva o Claims na variável permissions
	// e se tem o campo Valid (verificação de validade do próprio pacote jwt)
	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Faz um parse para recuperar do token o campo userId
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
		if err != nil {
			return 0, err
		}
		// Se tudo der certo, retorna o ID
		return userID, nil
	}
	// Senão, retorna erro
	return 0, errors.New("invalid token")
}

func getVerificationKey (token *jwt.Token) (interface{}, error){
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("signing method unexpected! %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}