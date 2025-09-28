package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(pwd string) (string, error) {
	hpwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

	return string(hpwd), err
}

func Compare(hpwd, pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hpwd), []byte(pwd))
}

func Sign(secretID string, sercretKey string, iss string, aud string) string {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Minute).Unix(),
		"nbf": time.Now().Unix(),
		"iat": time.Now().Unix(),
		"iss": iss,
		"aud": aud,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token.Header["kid"] = secretID

	signedToken, _ := token.SignedString([]byte(sercretKey))

	return signedToken
}
