package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	ErrorCreatingToken   = errors.New("error creating token")
	ErrorExtractingToken = errors.New("error extracting token")
	ErrorValidatingToken = errors.New("error validating token")
)

const (
	ExpiryTimeInMinutes = 10
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Authentication testing", r.Header.Get("x-auth"))

		next.ServeHTTP(rw, r)
	})
}

func CreateAuthToken(auth *Auth) (string, error) {
	if auth.UserId == "" {
		return "", ErrorCreatingToken
	}

	token := fmt.Sprintf("%s|%d", auth.UserId, auth.CreatedAt)

	return token, nil
}

func ExtractAuthToken(token string) (*Auth, error) {
	auth := &Auth{}

	components := strings.Split(token, "|")

	if len(components) != 2 {
		return nil, ErrorExtractingToken
	}

	auth.UserId = components[0]

	createdTimeInt, err := strconv.ParseInt(components[1], 10, 64)
	if err != nil {
		return nil, ErrorExtractingToken
	}

	auth.CreatedAt = time.Unix(createdTimeInt, 0).Unix()

	return auth, nil
}

func CheckTokenValidity(auth *Auth) bool {
	createAt := time.Unix(auth.CreatedAt, 0)

	if time.Now().Sub(createAt).Minutes() > ExpiryTimeInMinutes {
		log.Printf("Token expired for %s\n", auth.UserId)
		return false
	}

	return true
}
