package domain

import "time"

type Flags struct {
	Port *int
}

type Post struct {
	ID           uint      `form:"id" json:"id"`
	Title        string    `form:"title" json:"title"`
	Slug         string    `form:"slug" json:"slug"`
	Synopsis     string    `form:"synopsis" json:"synopsis"`
	Image        string    `form:"image" json:"image"` // This is the image path + filename.
	Content      string    `form:"content" json:"content"`
	RelatedPosts string    `form:"related_posts" json:"related_posts"` // array of int of related posts.
	CreatedAt    time.Time `form:"created_at" json:"created_at"`
}

type Repositories struct {
	PostRep PostRepository
}

type ServerInput struct {
	Reps *Repositories
}
