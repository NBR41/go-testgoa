package display

import (
	"github.com/NBR41/go-testgoa/internal/mail/content"
)

func ExampleDo() {
	a := Actionner{}
	err := a.Do("foo@bar.com", &content.Mail{Subject: "baz", Body: "qux"})
	if err != nil {
		panic(err)
	}
	// Output: to: foo@bar.com
	//Subject: baz
	//Body: qux
}
