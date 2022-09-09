package database

import "afa7789/site/internal/domain"

// NewRepositories creates a new repositories collection struct to be used.
func NewRepositories() *domain.Repositories {
	db := NewDatabase()
	postRepository := NewPostRepository(db)
	linkRepository := NewLinkRepository(db)
	messageRepository := NewMessageRepository(db)
	return &domain.Repositories{
		PostRep:    postRepository,
		LinkRep:    linkRepository,
		MessageRep: messageRepository,
	}
}
