package database

import (
	"afa7789/site/internal/domain"

	"gorm.io/gorm/clause"
)

type PostRepository struct {
	db *Database
}

func NewPostRepository(db *Database) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (pr *PostRepository) AddPost(p *domain.Post) uint {
	// Update all columns, except primary keys, to new value on conflict
	pr.db.client.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(p)

	// upsert
	return p.ID
}

func (pr *PostRepository) RetrievePosts(arr []uint) ([]domain.Post, error) {
	var posts []domain.Post

	result := pr.db.client.Find(&posts, arr)
	if result.Error != nil {
		return posts, result.Error
	}
	// select
	return posts, nil
}
