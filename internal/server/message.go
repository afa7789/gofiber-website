package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// messagesView renders the lists of messages page
func (s *Server) messagesView() fiber.Handler {
	return func(c *fiber.Ctx) error {

		// mount the stuff to show
		mesg, _ := s.reps.MessageRep.AllMessages()

		MessageIds := []uint{}
		MessageSubject := []string{}
		MessageName := []string{}
		MessageEmail := []string{}
		MessageText := []string{}

		for _, p := range mesg {
			MessageIds = append(MessageIds, p.ID)
			MessageSubject = append(MessageSubject, p.Subject)
			MessageName = append(MessageName, p.Name)
			MessageEmail = append(MessageEmail, p.Email)
			MessageText = append(MessageText, p.Text)
		}

		// blog posts
		return c.Status(http.StatusOK).Render("messages.html", fiber.Map{
			"Title":           "Messages - afa7789 ",
			"MessageIds":      MessageIds,
			"MessageSubjects": MessageSubject,
			"MessageNames":    MessageName,
			"MessageTexts":    MessageText,
			"MessageEmails":   MessageEmail,
		})
	}
}
