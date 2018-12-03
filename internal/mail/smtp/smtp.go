package smtp

import (
	"fmt"
	"net/mail"
	gosmtp "net/smtp"
	"strings"

	"github.com/NBR41/go-testgoa/internal/mail/content"
)

type Actionner struct {
}

func (a *Actionner) Do(email string, cnt *content.Mail) error {
	addr := "localhost:25"
	//toNames := []string{}
	toEmails := []string{email}
	// Build RFC-2822 email
	toAddresses := []string{}
	for i := range toEmails {
		to := mail.Address{Address: toEmails[i]}
		toAddresses = append(toAddresses, to.String())
	}

	from := mail.Address{Name: "admin", Address: "admin@localhost.org"}
	fromHeader := from.String()

	header := make(map[string]string)
	header["To"] = strings.Join(toAddresses, ", ")
	header["From"] = fromHeader
	header["Subject"] = cnt.Subject
	header["Content-Type"] = `text/html; charset="UTF-8"`
	msg := ""
	for k, v := range header {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	msg += "\r\n" + cnt.Body
	bMsg := []byte(msg)
	// Send using local postfix service
	c, err := gosmtp.Dial(addr)
	if err != nil {
		return err
	}

	defer c.Close()
	if err = c.Mail(fromHeader); err != nil {
		return err
	}
	for _, addr := range toEmails {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(bMsg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	err = c.Quit()
	// Or alternatively, send with remote service like Amazon SES
	// err = smtp.SendMail(addr, auth, fromEmail, toEmails, bMsg)
	// Handle response from local postfix or remote service
	if err != nil {
		return err
	}
	return nil
}
