package mail

import (
	"net/smtp"
	"os"
)

type SMTP struct {
	Auth smtp.Auth
	Port string
	Host string
	From string
}

// NewSmtpServer
func NewSmtpServer() *SMTP {
	// from is senders email address

	// we used environment variables to load the
	// email address and the password from the shell
	// you can also directly assign the email address
	// and the password
	from := os.Getenv("SMTP_MAIL")
	username := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASS")

	// host is address of server that the
	// sender's email address belongs,
	// in this case its gmail.
	// For e.g if your are using yahoo
	// mail change the address as smtp.mail.yahoo.com
	host := os.Getenv("SMTP_HOST")

	// PlainAuth uses the given username and password to
	// authenticate to host and act as identity.
	// Usually identity should be the empty string,
	// to act as username.
	auth := smtp.PlainAuth("", username, password, host)
	// Its the default port of smtp server
	port := os.Getenv("SMTP_PORT")

	// Create a new SMTP server
	return &SMTP{
		Auth: auth,
		Port: port,
		Host: host,
		From: from,
	}
}

// Send will send a msg to every recipient in to array
func (s *SMTP) Send(to []string, msg string) error {

	// This is the message to send in the mail
	// msg := "Hello geeks!!!"

	// We can't send strings directly in mail,
	// strings need to be converted into slice bytes
	body := []byte(msg)

	// SendMail uses TLS connection to send the mail
	// The email is sent to all address in the toList,
	// the body should be of type bytes, not strings
	// This returns error if any occurred.
	err := smtp.SendMail(s.Host+":"+s.Port, s.Auth, s.From, to, body)

	// handling the errors
	if err != nil {
		return err
		// os.Exit(1)
	}

	// fmt.Println("Successfully sent mail to all users in toList")
	return nil
}
