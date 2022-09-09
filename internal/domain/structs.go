package domain

import "time"

type Flags struct {
	Port *int
	TLS  *bool
}

type Post struct {
	ID           uint      `form:"id" json:"id" gorm:"primarykey"`
	Title        string    `form:"title" json:"title"`
	Slug         string    `form:"slug" json:"slug"`
	Synopsis     string    `form:"synopsis" json:"synopsis"`
	Image        string    `form:"image" json:"image"` // This is the image path + filename.
	Content      string    `form:"content" json:"content"`
	RelatedPosts string    `form:"related_posts" json:"related_posts"` // array of int of related posts.
	CreatedAt    time.Time `form:"created_at" json:"created_at"`
}

type Link struct {
	ID          uint      `form:"id" json:"id" gorm:"primarykey"`
	Title       string    `form:"title" json:"title"`
	HREF        string    `form:"href" json:"href"`
	Description string    `form:"description" json:"description"`
	Image       string    `form:"image" json:"image"` // This is a emoji or an svg file.
	CreatedAt   time.Time `form:"created_at" json:"created_at"`
}

type Message struct {
	ID        uint      `form:"id" json:"id" gorm:"primarykey"`
	Subject   string    `form:"subject" json:"subject"`
	Name      string    `form:"slug" json:"slug"`
	Text      string    `form:"text" json:"text"`
	Email     string    `form:"email" json:"email"`
	CreatedAt time.Time `form:"created_at" json:"created_at"`
}

type Repositories struct {
	PostRep    PostRepository
	LinkRep    LinkRepository
	MessageRep MessageRepository
}

type ServerInput struct {
	Reps *Repositories
}
