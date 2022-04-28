package domain

type PostRepository interface {
	AddPost(p *Post) uint
	RetrievePosts(arr []uint) ([]Post, error)
}
