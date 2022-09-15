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
	SwapOrder(sourceIndex, targetIndex uint) error
}

type MessageRepository interface {
	AddMessage(l *Message) uint
	RetrieveMessage(id uint) (*Message, error)
	RetrieveMessages(arr []uint) ([]Message, error)
	AllMessages() ([]Message, int64)
	DeleteMessage(id uint) error
}
