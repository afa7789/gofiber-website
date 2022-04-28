package domain

type Flags struct {
	Port *int
}

type Post struct {
	ID           uint   `form:"id" json:"id"`
	Title        string `form:"title" json:"title"`
	Synopsis     string `form:"synopsis" json:"synopsis"`
	Image        string `form:"image" json:"image"` // This is the image path + filename.
	Content      string `form:"content" json:"content"`
	RelatedPosts string `form:"related_posts" json:"related_posts"` // array of int of related posts.
}

type Repositories struct {
	PostRepository PostRepository
}

type ServerInput struct {
	repositories *Repositories
}
