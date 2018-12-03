package security

import (
	"bytes"
	"crypto/rand"
	"golang.org/x/crypto/scrypt"
)

//PasswordHelper struct for password helper
type PasswordHelper struct {
	pepper []byte
}

//NewPasswordHelper return
func NewPasswordHelper(pepper []byte) *PasswordHelper {
	return &PasswordHelper{pepper: pepper}
}

func (p PasswordHelper) getNewSalt() ([]byte, error) {
	salt := make([]byte, 32, 64)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func (p PasswordHelper) getSaltNPepper(salt []byte) []byte {
	return append(salt, p.pepper...)
}

// CryptPassword returns a hash and the salt
func (p PasswordHelper) CryptPassword(password string) ([]byte, []byte, error) {
	salt, err := p.getNewSalt()
	if err != nil {
		return nil, nil, err
	}

	hp, err := p.cryptPasswordWithSalt(password, salt)
	if err != nil {
		return nil, nil, err
	}
	return salt, hp, nil
}

func (p PasswordHelper) cryptPasswordWithSalt(password string, salt []byte) ([]byte, error) {
	hp, err := scrypt.Key([]byte(password), p.getSaltNPepper(salt), 16384, 8, 1, 64)
	if err != nil {
		return nil, err
	}
	return hp, nil
}

// ComparePassword return true if the hashed password with salt is equal to hash
func (p PasswordHelper) ComparePassword(password string, salt, hash []byte) (bool, error) {
	hp, err := p.cryptPasswordWithSalt(password, salt)
	if err != nil {
		return false, err
	}
	return (bytes.Compare(hp, hash) == 0), nil
}
