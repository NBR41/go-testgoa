package security

import (
	"testing"
)

func TestCryptPassword(t *testing.T) {
	p := NewPasswordHelper([]byte("0UArPJLVC3h667sQ"))
	salt, hash, err := p.CryptPassword("foo")
	if err != nil {
		t.Errorf("unexpected error, %v", err)
	} else {
		valid, err := p.ComparePassword("foo", salt, hash)
		if err != nil {
			t.Errorf("unexpected error, %v", err)
		} else {
			if !valid {
				t.Errorf("unexpected value, password is not the same")
			}
		}
	}
}
