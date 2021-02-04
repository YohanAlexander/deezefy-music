package album

import "github.com/yohanalexander/deezefy-music/entity"

// Read interface
type Read interface {
	Get(id int) (*entity.Album, error)
	Search(query string) ([]*entity.Album, error)
	List() ([]*entity.Album, error)
}

// Write interface
type Write interface {
	Create(e *entity.Album) (int, error)
	Update(e *entity.Album) error
	Delete(id int) error
}

// Repository interface
type Repository interface {
	Read
	Write
}

// UseCase interface
type UseCase interface {
	GetAlbum(id int) (*entity.Album, error)
	SearchAlbums(query string) ([]*entity.Album, error)
	ListAlbums() ([]*entity.Album, error)
	CreateAlbum(email, password, birthday, nomeartistico, biografia, titulo string, anoformacao, anolancamento, id int) (int, error)
	UpdateAlbum(e *entity.Album) error
	DeleteAlbum(id int) error
}
