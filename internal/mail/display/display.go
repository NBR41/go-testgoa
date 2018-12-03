package display

import (
	"fmt"

	"github.com/NBR41/go-testgoa/internal/mail/content"
)

type Actionner struct {
}

func (a *Actionner) Do(email string, cnt *content.Mail) error {
	fmt.Println(fmt.Sprintf("to: %s", email))
	fmt.Println(fmt.Sprintf("Subject: %s", cnt.Subject))
	fmt.Println(fmt.Sprintf("Body: %s", cnt.Body))
	return nil
}
