package database

import (
	"afa7789/site/internal/domain"

	"gorm.io/gorm/clause"
)

type LinkRepository struct {
	db *Database
}

// NewLinkRepository creates a new link repository.
func NewLinkRepository(db *Database) *LinkRepository {
	if db == nil {
		return nil
	}
	return &LinkRepository{
		db: db,
	}
}

// AddLink is an upsert, will update or insert
func (pr *LinkRepository) AddLink(l *domain.Link) uint {
	// Update all columns, except primary keys, to new value on conflict
	// upsert
	pr.db.client.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(l)

	return l.ID
}

// RetrieveLinks returns all links from the database
func (pr *LinkRepository) RetrieveLinks() ([]domain.Link, error) {
	var links []domain.Link

	// select all
	result := pr.db.client.Find(&links)
	if result.Error != nil {
		return links, result.Error
	}

	return links, nil
}

// RetrieveLinkByID returns a link by its ID
func (pr *LinkRepository) RetrieveLink(id uint) (*domain.Link, error) {
	link := &domain.Link{}

	// select id
	result := pr.db.client.First(link, id)
	if result.Error != nil {
		print("erro aqui ?:", result.Error.Error())
		print("\n")
		return nil, result.Error
	}

	return link, nil
}
