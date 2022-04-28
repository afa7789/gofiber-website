package database

import "afa7789/site/internal/domain"

func NewRepositories() *domain.Repositories {
	db := NewDatabase()
	postRepository := NewPostRepository(db)
	return &domain.Repositories{
		PostRepository: postRepository,
	}
}
