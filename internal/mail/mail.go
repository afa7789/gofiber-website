package mail

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"os"
)

type SMTP struct {
	tls          *tls.Config
	Auth         smtp.Auth
	Port         string
	Host         string
	From         string
	CompanyEmail string
}

// NewSMTPServer
func NewSMTPServer() *SMTP {
	// from is senders email address
	// we used environment variables to load the
	// email address and the password from the shell
	// you can also directly assign the email address
	// and the password
	from := os.Getenv("SMTP_MAIL")
	username := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASS")

	// company email
	companyEmail := os.Getenv("SMTP_COMPANY_EMAIL")

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

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// // from the very beginning (no starttls)
	// conn, err := tls.Dial("tcp", host+":"+port, tlsconfig)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// println("SMTP_HOST:", host)
	// println("SMTP_PORT:", port)
	// println("SMTP_MAIL:", from)
	// println("SMTP_USER:", username)
	// println("SMTP_PASS:", password)
	// println("SMTP_AUTH:", auth)
	// Create a new SMTP server
	return &SMTP{
		tls:          tlsconfig,
		Auth:         auth,
		Port:         port,
		Host:         host,
		From:         from,
		CompanyEmail: companyEmail,
	}
}

// Send will send a msg to every recipient in to array
func (s *SMTP) Send(to []string, msg string) error {
	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", s.Host+":"+s.Port, s.tls)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, s.Host)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		err := c.Quit()
		if err != nil {
			log.Panic(err)
		}
	}()

	// Auth
	if err := c.Auth(s.Auth); err != nil {
		log.Panic(err)
	}

	// This is the message to send in the mail
	// msg := "Hello geeks!!!"

	// We can't send strings directly in mail,
	// strings need to be converted into slice bytes
	body := []byte(msg)

	// To && From
	if err = c.Mail(s.From); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(to[0]); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write(body)
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	// handling the errors
	if err != nil {
		print(err.Error())
		return err
	}

	fmt.Println("Sent mail")
	return nil
}
