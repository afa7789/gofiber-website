package server

import (
	"afa7789/site/internal/mail"
	"log"

	"github.com/gofiber/fiber/v2"
)

type MailerController struct {
	mailer *mail.SMTP
}

func NewMailerController() *MailerController {
	mailer := mail.NewSMTPServer()
	return &MailerController{
		mailer: mailer,
	}
}

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

		if (body.Name == "" || body.ContactEmail == "") && body.Message == "" {
			return c.Status(fiber.StatusBadRequest).Redirect("/failed")
		}

		go func() {
			str := "From: " + body.Name + " [" + body.ContactEmail + "]\r\n" +
				"To: " + mc.mailer.CompanyEmail + "\r\n" +
				"Sender: " + body.Name + "\r\n" +
				"Subject: " + body.Subject + "\r\n" +
				body.Message + "\r\n\r\n" +
				"Sent from the afa7789 site on behalf of: " + body.ContactEmail + "\r\n "

			if err := mc.mailer.Send([]string{"" + mc.mailer.CompanyEmail + ""}, str); err != nil {
				// TODO change this LOG to log to a file.
				log.Default().Print("Error sending email: ", err)
			}
		}()

		return c.Status(fiber.StatusOK).Redirect("/thanks")
	}
}
