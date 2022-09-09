package database

import (
	"afa7789/site/internal/domain"
	"log"

	"gorm.io/gorm/clause"
)

type MessageRepository struct {
	db *Database
}

// NewMessageRepository creates a new Message repository.
func NewMessageRepository(db *Database) *MessageRepository {
	if db == nil {
		return nil
	}
	return &MessageRepository{
		db: db,
	}
}

func (pr *MessageRepository) DeleteMessage(id uint) error {
	message := &domain.Message{}

	// select id
	result := pr.db.client.First(message, id)
	if result.Error != nil {
		return result.Error
	}

	// delete
	result = pr.db.client.Delete(message)
	if result.Error != nil {
		return result.Error
	}

	return nil

}

// AddMessage is an upsert, will update or insert
func (pr *MessageRepository) AddMessage(p *domain.Message) uint {
	// Update all columns, except primary keys, to new value on conflict
	// upsert
	pr.db.client.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(p)

	return p.ID
}

// RetrieveMessages returns all Messages from the database in accord to their ids
func (pr *MessageRepository) RetrieveMessages(arr []uint) ([]domain.Message, error) {
	var Messages []domain.Message

	// select multiple ids
	result := pr.db.client.Find(&Messages, arr)
	if result.Error != nil {
		return Messages, result.Error
	}

	return Messages, nil
}

// RetrieveMessageByID returns a Message by its ID
func (pr *MessageRepository) RetrieveMessage(id uint) (*domain.Message, error) {
	Message := &domain.Message{}

	// select id
	result := pr.db.client.First(Message, id)
	if result.Error != nil {
		log.Printf("erro aqui ?:%s", result.Error.Error())
		return nil, result.Error
	}

	return Message, nil
}

func (pr *MessageRepository) AllMessages() ([]domain.Message, int64) {
	var Messages []domain.Message
	result := pr.db.client.Find(&Messages)
	if result.Error != nil {
		log.Printf("erro aqui ?:%s", result.Error.Error())
		return nil, 0
	}

	count := int64(0)
	pr.db.client.Model(&domain.Message{}).Count(&count)
	return Messages, count
}
