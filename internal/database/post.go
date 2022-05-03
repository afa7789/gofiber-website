package database

import (
	"afa7789/site/internal/domain"

	"gorm.io/gorm/clause"
)

type PostRepository struct {
	db *Database
}

// NewPostRepository creates a new post repository.
func NewPostRepository(db *Database) *PostRepository {
	if db == nil {
		return nil
	}
	return &PostRepository{
		db: db,
	}
}

// AddPost is an upsert, will update or insert
func (pr *PostRepository) AddPost(p *domain.Post) uint {
	// Update all columns, except primary keys, to new value on conflict
	// upsert
	pr.db.client.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(p)

	return p.ID
}

// RetrievePosts returns all posts from the database in accord to their ids
func (pr *PostRepository) RetrievePosts(arr []uint) ([]domain.Post, error) {
	var posts []domain.Post

	// select multiple ids
	result := pr.db.client.Find(&posts, arr)
	if result.Error != nil {
		return posts, result.Error
	}

	return posts, nil
}

// RetrievePostByID returns a post by its ID
func (pr *PostRepository) RetrievePost(id uint) (*domain.Post, error) {
	post := &domain.Post{}

	// select id
	result := pr.db.client.First(post, id)
	if result.Error != nil {
		print("erro aqui ?:", result.Error.Error())
		print("\n")
		return nil, result.Error
	}

	return post, nil
}

// LastThreePosts returns the last three posts looking (created at) from the database for the home page.
func (pr *PostRepository) LastThreePosts() ([]domain.Post, error) {
	var posts []domain.Post

	pr.db.client.Order("created_at desc").Find(&posts).Limit(3)

	return posts, nil
}

// PageResult does the pagination of results of blog posts for better viewing
func (pr *PostRepository) PageResult(page int) ([]domain.Post, int64) {
	var posts []domain.Post

	pr.db.client.Offset(page * domain.PageLimit).Limit(domain.PageLimit).Find(&posts)

	count := int64(0)
	pr.db.client.Model(&domain.Post{}).Count(&count)

	return posts, count
}
