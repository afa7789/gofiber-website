package domain

type PostRepository interface {
	AddPost(p *Post) uint
	RetrievePosts(arr []uint) ([]Post, error)
	RetrievePost(id uint) (*Post, error)
	LastThreePosts() ([]Post, error)
	PageResult(page int) ([]Post, int64)
}

type LinkRepository interface {
	AddLink(l *Link) uint
	RetrieveLinks() ([]Link, error)
	RetrieveLink(id uint) (*Link, error)
	DeleteLink(id uint) error
}
