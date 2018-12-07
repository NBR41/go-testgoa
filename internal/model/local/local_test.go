package local

import (
	"testing"
)

func expectingError(t *testing.T, err, exp error) {
	if err == nil {
		t.Fatal("expecting error got nil")
	} else {
		if err.Error() != exp.Error() {
			t.Fatal("unexpected error", err)
		}
	}
}
