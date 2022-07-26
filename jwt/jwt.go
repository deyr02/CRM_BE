package jwt

import (
	"log"

	"github.com/deyr02/bnzlcrm/graph/model"
	"github.com/golang-jwt/jwt"
)

// secret key being used to sign tokens
var (
	SecretKey = []byte("bnzl_secret_123456")
)

// GenerateToken generates a jwt token and assign a username to it's claims and return it
func GenerateToken(tokenServiceDto *model.TokenServiceDto) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)
	/* Set token claims */
	claims["username"] = tokenServiceDto.UserName
	claims["roleid"] = tokenServiceDto.RoleID
	claims["expirydate"] = tokenServiceDto.ExpiryDate //time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in Generating key")
		return "", err
	}
	return tokenString, nil
}

// ParseToken parses a jwt token and returns the username in it's claims
func ParseToken(tokenStr string) (*model.TokenServiceDto, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		tokenServiceDto := &model.TokenServiceDto{
			UserName:   claims["username"].(string),
			RoleID:     claims["roleid"].(string),
			ExpiryDate: claims["expirydate"].(string),
		}

		return tokenServiceDto, nil
	} else {
		return nil, err
	}
}
