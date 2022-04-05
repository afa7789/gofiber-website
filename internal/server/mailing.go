package server

import (
	"afa7789/site/internal/mail"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type MailerController struct {
	mailer *mail.SMTP
}

func NewMailerController() *MailerController {
	mailer := mail.NewSmtpServer()
	return &MailerController{
		mailer: mailer,
	}
}

func (mc *MailerController) send() fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := struct {
			Name         string `json:"name"`
			ContactEmail string `json:"contactEmail"`
			Subject      string `json:"subject"`
			Message      string `json:"message"`
		}{}

		if err := c.BodyParser(&body); err != nil {
			print("body parse error")
			return err
		}

		str := "From: " + body.Name + " <" + body.ContactEmail + ">\r\n" +
			"To: afa7789@gmail.com\r\n" +
			"Sender: " + body.Name + "\r\n" +
			"Subject: " + body.Subject + "\r\n\r\n" +
			body.Message + "\r\n"

		if err := mc.mailer.Send([]string{"afa7789@gmail.com"}, str); err != nil {
			print("send error")
			return err
		}

		return c.Status(http.StatusOK).JSON(true)
	}
}
