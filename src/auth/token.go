package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(id uint) (string, error) {
	permissions := jwt.MapClaims{}

	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["id"] = id

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(config.SecretKey))
}

func ValidateToken(r *http.Request) error {
	token := extractToken(r)

	parsedToken, err := jwt.Parse(token, getVerificationKey)

	if err != nil {
		return err
	}

	if _, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return nil
	}

	return errors.New("token inválido")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	tokenSplit := strings.Split(token, " ")

	if len(tokenSplit) != 2 {
		return ""
	}

	return tokenSplit[1]
}

func getVerificationKey(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}

func ExtractUserIdToken(r *http.Request) (uint64, error) {
	token := extractToken(r)

	parsedToken, err := jwt.Parse(token, getVerificationKey)

	if err != nil {
		return 0, err
	}

	if permissions, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		id, err := strconv.ParseUint(
			fmt.Sprintf("%.0f", permissions["id"]),
			10,
			64,
		)

		if err != nil {
			return 0, err
		}

		return id, nil
	}

	return 0, errors.New("token inválido")
}
