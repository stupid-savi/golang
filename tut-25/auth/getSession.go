package auth

import (
	"crypto/rand"
	"encoding/base64"
)

func session(length int) (string, error) {

	byteLength := (length * 3) / 4
	randomBytes := make([]byte, byteLength)
	_, err := rand.Read(randomBytes)

	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(randomBytes)[:byteLength], nil
}

func GetUserSession() string {
	sessionToken, err := session(128)
	if err != nil {
		return "401 Error!"
	}

	return sessionToken
}
