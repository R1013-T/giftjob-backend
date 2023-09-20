package middleware

import (
	"giftjob-backend/utils"
	"github.com/go-jose/go-jose/v3"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
)

func DecryptAndVerifyJWEToken(token string, privateKey string) ([]byte, error) {
	object, err := jose.ParseEncrypted(token)
	if err != nil {
		log.Println("Error parsing JWE token:", err)
		return nil, err
	}

	decrypted, err := object.Decrypt(privateKey)
	if err != nil {
		log.Println("Error decrypting JWE token:", err)
		return nil, err
	}

	log.Printf("Decrypted token: %s", decrypted)

	return decrypted, nil
}

func JWEAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	log.Println("JWEAuthentication middleware started")

	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
		}

		// "Bearer"の後にトークンがあるはずなので、それを取得
		splitted := strings.Split(authHeader, " ")
		if len(splitted) != 2 {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization header format")
		}

		token := splitted[1]

		_, err := DecryptAndVerifyJWEToken(token, utils.Getenv("JWE_SECRET"))
		if err != nil {
			log.Println("Error decrypting and verifying JWE token:", err)
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		log.Printf("JWE token verified: %s", token)
		return next(c)
	}
}
