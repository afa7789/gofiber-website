package server

import (
	"afa7789/site/internal/mail"
	"log"

	"github.com/gofiber/fiber/v2"
)

type MailerController struct {
	mailer *mail.SMTP
}

func newMailerController() *MailerController {
	mailer := mail.NewSMTPServer()
	return &MailerController{
		mailer: mailer,
	}
}

// send mail
func (mc *MailerController) send() fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := struct {
			Name         string `form:"name"`
			ContactEmail string `form:"contact"`
			Subject      string `form:"subject"`
			Message      string `form:"message"`
		}{}

		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(struct {
				Message string `json:"message"`
			}{
				Message: "Error at content parsing",
			})
		}

		if validator(body.Name, body.ContactEmail, body.Message) {
			return c.Status(fiber.StatusBadRequest).Redirect("/failed")
		}

		go func() {
			str := emailConstructor(body.Name, body.Subject, body.Message, body.ContactEmail, mc.mailer.CompanyEmail)
			if err := mc.mailer.Send([]string{"" + mc.mailer.CompanyEmail + ""}, str); err != nil {
				// TODO change this LOG to log to a file.
				log.Default().Print("Error sending email: ", err)
			}
		}()

		return c.Status(fiber.StatusOK).Redirect("/thanks")
	}
}

func validator(name, contactEmail, message string) bool {
	if (name == "" || contactEmail == "") && message == "" {
		return false
	}
	return true
}

// emailConstructor is the mail builder
func emailConstructor(name, subject, message, contactEmail, companyEmail string) string {
	return "From: " + name + " [" + contactEmail + "]\r\n" +
		"To: " + companyEmail + "\r\n" +
		"Sender: " + name + "\r\n" +
		"Subject: " + subject + "\r\n" +
		message + "\r\n\r\n" +
		"Sent from the afa7789 site on behalf of: " + contactEmail + "\r\n "
}
