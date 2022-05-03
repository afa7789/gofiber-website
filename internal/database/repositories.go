package database

import "afa7789/site/internal/domain"

// NewRepositories creates a new repositories collection struct to be used.
func NewRepositories() *domain.Repositories {
	db := NewDatabase()
	postRepository := NewPostRepository(db)
	return &domain.Repositories{
		PostRep: postRepository,
	}
}
