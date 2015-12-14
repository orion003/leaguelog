package jwt

import (
	"errors"
	"fmt"

	jwt "leaguelog/Godeps/_workspace/src/github.com/dgrijalva/jwt-go"
)

type Jwt struct {
	key []byte
}

func InitializeJwt(k []byte) *Jwt {
	j := &Jwt{
		key: k,
	}

	return j
}

func (j *Jwt) Generate(claims map[string]interface{}) string {
	token := jwt.New(jwt.SigningMethodHS512)
	token.Claims = claims

	tokenString, err := token.SignedString(j.key)
	if err != nil {
		fmt.Printf("Unable to sign token: %v\n", err)
		return ""
	}

	return tokenString
}

func (j *Jwt) Validate(tString string) error {
	t, err := jwt.Parse(tString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return j.key, nil
	})

	var errorString string
	if t.Valid {
		fmt.Println("Valid token")
		return nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			errorString = "Invalid token found"
		} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
			errorString = "Token has expired"
		} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
			errorString = "Token is not yet valid"
		} else {
			errorString = fmt.Sprintf("Unable to handle token: %v", err)
		}
	} else {
		errorString = fmt.Sprintf("Unable to handle token: %v", err)
	}

	return errors.New(errorString)
}
