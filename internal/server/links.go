package server

import (
	"afa7789/site/internal/domain"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// validateLink validates the post data
func validateLink(l *domain.Link) bool {
	if l.HREF != "" && l.Title != "" {
		return true
	}
	return false
}

// Receive link from a multi-form from the add link page.
func (s *Server) receiveLink() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var link domain.Link
		// parsinsg the link that's in the form coming from the request
		if err := c.BodyParser(&link); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(struct {
				Message string `json:"message"`
			}{
				Message: "Error at content parsing " + err.Error(),
			})
		}

		if !validateLink(&link) {
			return c.Status(fiber.StatusBadRequest).JSON(link)
		}

		// upload or create the link
		// if it's create will be sent with link id 0.
		s.reps.LinkRep.AddLink(&link)

		return c.Status(fiber.StatusOK).JSON(link)
	}
}

// linksView renders a lists of links in the link page.
func (s *Server) linksView() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// mount the stuff to show
		links, _ := s.reps.LinkRep.RetrieveLinks()

		LinksIds := []uint{}
		LinksTitles := []string{}
		LinksImages := []template.HTML{}
		LinksDescriptions := []string{}
		LinksHREFs := []string{}

		for _, l := range links {
			LinksIds = append(LinksIds, l.ID)
			LinksTitles = append(LinksTitles, l.Title)
			LinksImages = append(LinksImages, template.HTML(l.Image))

			LinksDescriptions = append(LinksDescriptions, l.Description)
			LinksHREFs = append(LinksHREFs, l.HREF)
		}

		fmt.Printf("links len %d", len(LinksImages))

		// link list
		return c.Status(http.StatusOK).Render("links.html", fiber.Map{
			"Title":             "Links - afa7789 ",
			"LinksID":           LinksIds,
			"LinksHREFs":        LinksHREFs,
			"LinksImages":       LinksImages,
			"LinksTitles":       LinksTitles,
			"LinksDescriptions": LinksDescriptions,
		})
	}
}

// linkEditor opens the template
// this func returns a page to edit an old link or create a newer one
func (s *Server) linkEditor() fiber.Handler {
	return func(c *fiber.Ctx) error {
		linkID := c.Params("link_id")
		// from string to uint
		IDLink, err := strconv.ParseUint(linkID, 10, 32)
		// cast uint64 to uint
		if err != nil {
			log.Default().Printf("Error with link ID = %s : %s", linkID, err.Error())
			linkID = ""
		}

		link, err := s.reps.LinkRep.RetrieveLink(uint(IDLink))
		if err != nil {
			log.Default().Printf("Error with link ID = %s : %s", linkID, err.Error())
			linkID = ""
		}

		if linkID != "" {
			// retrieve link data
			return c.Status(http.StatusOK).Render("link.html", fiber.Map{
				"Title":           "Link Editor - " + linkID + " - afa7789 ",
				"LinkID":          linkID,
				"LinkImage":       link.Image,
				"LinkTitle":       link.Title,
				"LinkHREF":        link.HREF,
				"LinkDescription": link.Description,
			})
		}

		return c.Status(http.StatusOK).Render("link.html", fiber.Map{
			"Title": "Link Creator - afa7789 ",
		})
	}
}

// swapIndexes route
func (s *Server) swapIndexes() fiber.Handler {
	return func(c *fiber.Ctx) error {
		targetID := c.Params("target_id")
		// from string to uint
		tID, err := strconv.ParseUint(targetID, 10, 32)
		if err != nil {
			log.Default().Printf("Error with target ID = %s : %s", targetID, err.Error())
			return c.Status(http.StatusBadRequest).JSON("wrong targetID: " + err.Error())
		}

		sourceID := c.Params("source_id")
		// from string to uint
		sID, err := strconv.ParseUint(sourceID, 10, 32)
		if err != nil {
			log.Default().Printf("Error with source ID = %s : %s", sourceID, err.Error())
			return c.Status(http.StatusBadRequest).JSON("wrong sourceID: " + err.Error())
		}

		// call bd function to swap
		err = s.reps.LinkRep.SwapOrder(uint(sID), uint(tID))
		if err != nil {
			log.Default().Printf("Error on swapping", err.Error())
			return c.Status(http.StatusInternalServerError).JSON("not swapped: " + err.Error())
		}

		return c.Status(http.StatusOK).JSON("swapped")
	}
}

// deliteLink route
func (s *Server) deleteLink() fiber.Handler {
	return func(c *fiber.Ctx) error {
		linkID := c.Params("link_id")
		// from string to uint
		IDLink, err := strconv.ParseUint(linkID, 10, 32)
		// cast uint64 to uint
		if err != nil {
			log.Default().Printf("Error with link ID = %s : %s", linkID, err.Error())
			linkID = ""
			return c.Status(http.StatusBadRequest).JSON("wrong link id")
		}

		err = s.reps.LinkRep.DeleteLink(uint(IDLink))
		if err != nil {
			log.Default().Printf("Error on deleting link ID = %s : %s", linkID, err.Error())
			return c.Status(http.StatusInternalServerError).JSON("not deleted")
		}

		return c.Status(http.StatusOK).JSON("deleted")
	}
}
