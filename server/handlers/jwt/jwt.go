package jwt

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const secretKey = "lahnasti"

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := isTokenValid(ctx)
		if err != nil {
			ctx.String(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func GenerateToken(uid string) (string, error) {
	token_life := 1           //время жизни токена 1 час
	claims := jwt.MapClaims{} // это набор данных (ключ-значение), которые будут храниться в токене. Эти данные могут быть проверены при декодировании токена.
	claims["authorized"] = true
	claims["uid"] = uid
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_life)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func ExtraToken(ctx *gin.Context) string {
	token := ctx.Query("token")
	if token != "" {
		return token
	}
	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func isTokenValid(ctx *gin.Context) error {
	tokenString := ExtraToken(ctx)
	// разбор токена из строки и проверка его подписи с функ обратного вызова
	_, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		//проверка метода подписания токена
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	return err
}
