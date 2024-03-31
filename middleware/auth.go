package middleware

import (
	"encoding/base64"
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
		log.Println("Path : ", r.URL.Path)

		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, x-auth")
		rw.Header().Set("Access-Control-Allow-Methods", "*")
		rw.Header().Set("Content-Type", "application/json")

		if r.URL.Path == "/user/authenticate" || r.Method == "OPTIONS" {
			next.ServeHTTP(rw, r)
			return
		}

		log.Println("Authentication testing", r.Header.Get("x-auth"))

		// auth, err := ExtractAuthToken(r.Header.Get("x-auth"))

		// if err != nil {
		// 	fmt.Println("dinga error", err)
		// 	b, _ := json.Marshal(model.Error{
		// 		Message: "Not Authenticated",
		// 		Status:  400,
		// 	})
		// 	rw.Header().Set("Access-Control-Allow-Origin", "*")
		// 	rw.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, x-auth")
		// 	rw.Header().Set("Access-Control-Allow-Methods", "*")
		// 	rw.Header().Set("Content-Type", "application/json")
		// 	rw.WriteHeader(http.StatusForbidden)
		// 	rw.Write([]byte(b))
		// 	return
		// }

		// isValid := CheckTokenValidity(auth)
		// if !isValid {
		// 	fmt.Println("Not Valid token", err)
		// 	b, _ := json.Marshal(model.Error{
		// 		Message: "Invalid Token",
		// 		Status:  400,
		// 	})
		// 	rw.Header().Set("Access-Control-Allow-Origin", "*")
		// 	rw.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, x-auth")
		// 	rw.Header().Set("Access-Control-Allow-Methods", "*")
		// 	rw.Header().Set("Content-Type", "application/json")
		// 	rw.WriteHeader(http.StatusForbidden)
		// 	rw.Write([]byte(b))
		// 	return
		// }

		next.ServeHTTP(rw, r)
	})
}

func CreateAuthToken(auth *Auth) (string, error) {
	if auth.UserId == "" {
		return "", ErrorCreatingToken
	}

	token := fmt.Sprintf("%s|%d", auth.UserId, auth.CreatedAt)

	token = base64.StdEncoding.EncodeToString([]byte(token))

	return token, nil
}

func ExtractAuthToken(token string) (*Auth, error) {
	auth := &Auth{}

	tokenDecoded, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, ErrorExtractingToken
	}

	components := strings.Split(string(tokenDecoded), "|")

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
