package domain

type PostRepository interface {
	AddPost(p *Post) (uint, error)
	RetrievePosts(arr []uint) ([]Post, error)
}
