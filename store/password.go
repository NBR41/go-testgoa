package store

import (
	"bytes"
	"crypto/rand"
	"golang.org/x/crypto/scrypt"
)

var pepper = []byte("0UArPJLVC3h667sQ")

func getNewSalt() ([]byte, error) {
	salt := make([]byte, 32, 64)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func getSaltNPepper(salt []byte) []byte {
	return append(salt, pepper...)
}

func cryptPassword(password string) ([]byte, []byte, error) {
	salt, err := getNewSalt()
	if err != nil {
		return nil, nil, err
	}

	hp, err := cryptPasswordWithSalt(password, salt)
	if err != nil {
		return nil, nil, err
	}
	return salt, hp, nil
}

func cryptPasswordWithSalt(password string, salt []byte) ([]byte, error) {
	hp, err := scrypt.Key([]byte(password), getSaltNPepper(salt), 16384, 8, 1, 64)
	if err != nil {
		return nil, err
	}
	return hp, nil
}

func comparePassword(password string, salt, hash []byte) (bool, error) {
	hp, err := cryptPasswordWithSalt(password, salt)
	if err != nil {
		return false, err
	}
	return (bytes.Compare(hp, hash) == 0), nil
}
