package util

import "golang.org/x/crypto/bcrypt"

func SetPassword(origin string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(origin), 12)
	if err != nil {
		return "", err
	}
	return string(password), nil
}

func ComparePassword(password string, passwordDigest string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordDigest), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
