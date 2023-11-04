package middleware

import (
	"encoding/base64"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateAuthToken(t *testing.T) {
	auth := &Auth{
		"user_id",
		1697246834,
	}

	token, err := CreateAuthToken(auth)

	expected := fmt.Sprintf("%s|%d", auth.UserId, auth.CreatedAt)
	expected = base64.StdEncoding.EncodeToString([]byte(expected))

	assert.Nil(t, err)
	assert.Equal(t, expected, token)
}

func TestExtractAuthToken(t *testing.T) {
	userId := "san"
	createdAt := time.Now().Unix()
	token := fmt.Sprintf("%s|%d", userId, createdAt)
	token = base64.StdEncoding.EncodeToString([]byte(token))

	auth, err := ExtractAuthToken(token)

	assert.Nil(t, err)
	assert.Equal(t, userId, auth.UserId)
	assert.Equal(t, createdAt, auth.CreatedAt)
}

func TestValidateAuthToken(t *testing.T) {
	userId := "san"

	t.Run("valid token", func(t *testing.T) {
		createdAt := time.Now().Unix()
		token := fmt.Sprintf("%s|%d", userId, createdAt)
		token = base64.StdEncoding.EncodeToString([]byte(token))

		auth, err := ExtractAuthToken(token)

		assert.Nil(t, err)
		assert.Equal(t, userId, auth.UserId)
		assert.Equal(t, createdAt, auth.CreatedAt)

		valid := CheckTokenValidity(auth)

		assert.Equal(t, true, valid)
	})

	t.Run("invalid token", func(t *testing.T) {
		createdAt := time.Now().Add(-11 * time.Minute).Unix()
		token := fmt.Sprintf("%s|%d", userId, createdAt)
		token = base64.StdEncoding.EncodeToString([]byte(token))

		auth, err := ExtractAuthToken(token)

		assert.Nil(t, err)
		assert.Equal(t, userId, auth.UserId)
		assert.Equal(t, createdAt, auth.CreatedAt)

		valid := CheckTokenValidity(auth)

		assert.Equal(t, false, valid)
	})
}
