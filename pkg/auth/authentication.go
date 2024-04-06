package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(username string) (string, error) {
	switch CryptographyAlgorithm {
	case HMAC:
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
		return token.SignedString([]byte(encryptionKey))

	case ECDSA:
		token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
		tokenString, err := token.SignedString(privateKeyECDSA)
		if err != nil {
			return "", err
		}
		return tokenString, nil
	}
	return "", errors.New("invalid algorithm")
}

func ValidateToken(tokenString string) (string, error) {
	token, err := parseByCryptographyAlgorithm(tokenString)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}
	username, ok := claims["username"].(string)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}
	return username, nil

}

func parseByCryptographyAlgorithm(tokenString string) (*jwt.Token, error) {
	var (
		token *jwt.Token
		err   error
	)
	switch CryptographyAlgorithm {
	case HMAC:
		token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(encryptionKey), nil
		})
		if err != nil {
			return nil, err
		}

	case ECDSA:
		token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return &privateKeyECDSA.PublicKey, nil
		})
		if err != nil {
			return nil, err
		}
		if !token.Valid {
			return nil, fmt.Errorf("token is not valid")
		}
	default:
		return nil, errors.New("invalid algorithm")

	}
	return token, nil
}
