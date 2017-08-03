package appmail

import (
	"fmt"
	"net/mail"
	"net/smtp"
	"strings"

	"github.com/NBR41/go-testgoa/appmodel"
)

var (
	defaultemail = "fabien@localhost"
	baseURL      = "http://localhost"

	fmtResetPasswordBody    = "You have forgotten your password ?\r\nCopy the link below in your favorite browser:\r\n\r\n%s/password/reset?t=%s"
	fmtModifiedPasswordBody = "Your password has been successfully updated.\r\n Visit: %s"
	fmtNewUserBody          = "Hello %s,\r\nYour account has been successfully created.\r\n Visit: %s"
	fmtValidateUserBody     = "Hello %s,\r\nYour account need to be validated.\r\nCopy the link below in your favorite browser:\r\n\r\n%s/user/activate?t=%s"
	fmtValidatedUserBody    = "Your account has been successfully validated.\r\n Visit: %s"
)

// SendResetPasswordMail send reset password link mail
func SendResetPasswordMail(email, token string) error {
	return send(email, "MyInventory: Reset Password", fmt.Sprintf(fmtResetPasswordBody, baseURL, token))
}

// SendPasswordUpdatedMail send reset password notification mail
func SendPasswordUpdatedMail(email string) error {
	return send(email, "MyInventory: Password update", fmt.Sprintf(fmtResetPasswordBody, baseURL))
}

// SendNewUserMail send user creation mail
func SendNewUserMail(u *appmodel.User, token string) error {
	return send(u.Email, "MyInventory: New Account", fmt.Sprintf(fmtNewUserBody, u.Nickname, baseURL))
}

// SendActivationMail send user activation mail
func SendActivationMail(u *appmodel.User, token string) error {
	return send(u.Email, "MyInventory: Validate your Account", fmt.Sprintf(fmtValidateUserBody, u.Nickname, baseURL, token))
}

// SendUserActivatedMail send activated user notification mail
func SendUserActivatedMail(email string) error {
	return send(email, "MyInventory: Your account is validated", fmt.Sprintf(fmtValidatedUserBody, baseURL))
}

func send(email, subject, body string) error {
	addr := "localhost:25"
	//toNames := []string{}
	toEmails := []string{defaultemail}
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
	header["Subject"] = subject
	header["Content-Type"] = `text/html; charset="UTF-8"`
	msg := ""
	for k, v := range header {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	msg += "\r\n" + body
	bMsg := []byte(msg)
	// Send using local postfix service
	c, err := smtp.Dial(addr)
	if err != nil {
		return err
	}

	defer c.Close()
	if err = c.Mail(fromHeader); err != nil {
		return err
	}
	for _, addr := range toEmails {
		fmt.Println(addr)
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
