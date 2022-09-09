package server

import (
	"log"
	"net/http"
	"strconv"

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

func (s *Server) deleteMessage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//
		id := c.Params("delete_id")
		// from string to uint
		idd, err := strconv.ParseUint(id, 10, 32)
		// cast uint64 to uint
		if err != nil {
			log.Default().Printf("Error with link ID = %s : %s", id, err.Error())
			id = ""
		}
		log.Printf("id %s, idd %d", id, uint(idd))

		err = s.reps.MessageRep.DeleteMessage(uint(idd))

		// blog posts
		if err != nil {
			log.Printf("Error at delete: %s", err.Error())
			return c.Status(http.StatusInternalServerError).JSON(struct {
				Message string `json:"message"`
			}{
				Message: "Error at deleting Message: " + err.Error(),
			})
		}
		return c.Status(http.StatusOK).Status(http.StatusInternalServerError).JSON(struct {
			Message string `json:"message"`
		}{
			Message: "Okay",
		})
	}
}
