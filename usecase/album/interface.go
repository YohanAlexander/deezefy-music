package album

import "github.com/yohanalexander/deezefy-music/entity"

// Album interface
type Album interface {
	Get(id int) (*entity.Album, error)
	Search(query string) ([]*entity.Album, error)
	List() ([]*entity.Album, error)
	Create(e *entity.Album) (int, error)
	Update(e *entity.Album) error
	Delete(id int) error
}

// Repository interface
type Repository interface {
	Album
}

// UseCase interface
type UseCase interface {
	GetAlbum(id int) (*entity.Album, error)
	SearchAlbums(query string) ([]*entity.Album, error)
	ListAlbums() ([]*entity.Album, error)
	CreateAlbum(id, anolancamento int, titulo, artista string) (int, error)
	UpdateAlbum(e *entity.Album) error
	DeleteAlbum(id int) error
}
