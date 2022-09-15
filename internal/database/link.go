package database

import (
	"afa7789/site/internal/domain"
	"fmt"

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
	result := pr.db.client.Order("index_order asc").Find(&links)
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

// DeleteLink receive id as uint and uses gorm to delete it
func (pr *LinkRepository) DeleteLink(id uint) error {
	link := &domain.Link{}

	// select id
	result := pr.db.client.First(link, id)
	if result.Error != nil {
		return result.Error
	}

	// delete
	result = pr.db.client.Delete(link)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (pr *LinkRepository) SwapOrder(sourceIndex, targetIndex uint) error {

	// check if both ids exists
	var count int64
	result := pr.db.client.Table("links").
		Where(fmt.Sprintf("index_order = %d", sourceIndex)).
		Or(fmt.Sprintf("index_order = %d", targetIndex)).
		Count(&count)
	if result.Error != nil {
		return fmt.Errorf("at verification")
	}
	if count != 2 {
		return fmt.Errorf("either source or target index does not exist")
	}

	result = pr.db.client.Table("links").Count(&count)
	if result.Error != nil {
		return fmt.Errorf("at count part of swap: %w", result.Error)
	}

	//swap to an auxiliar number for now
	result = pr.db.client.Table("links").Where("index_order = ?", sourceIndex).Update("index_order", count+1)
	if result.Error != nil {
		return fmt.Errorf("at first part of swap: %w", result.Error)
	}

	//set ID
	result = pr.db.client.Table("links").Where("index_order = ?", targetIndex).Update("index_order", sourceIndex)
	if result.Error != nil {
		return fmt.Errorf("at second part of swap: %w", result.Error)
	}

	//set ID
	result = pr.db.client.Table("links").Where("index_order = ?", count+1).Update("index_order", targetIndex)
	if result.Error != nil {
		return fmt.Errorf("at third part of swap: %w", result.Error)
	}

	return nil
}
