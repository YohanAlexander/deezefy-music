package album

import der "github.com/yohanalexander/deezefy-music/entity/album"

// Album interface
type Album interface {
	Get(id int) (*der.Album, error)
	Search(query string) ([]*der.Album, error)
	List() ([]*der.Album, error)
	Create(e *der.Album) (int, error)
	Update(e *der.Album) error
	Delete(id int) error
}

// Repository interface
type Repository interface {
	Album
}

// UseCase interface
type UseCase interface {
	GetAlbum(id int) (*der.Album, error)
	SearchAlbums(query string) ([]*der.Album, error)
	ListAlbums() ([]*der.Album, error)
	CreateAlbum(id, anolancamento int, titulo, artista string) (int, error)
	UpdateAlbum(e *der.Album) error
	DeleteAlbum(id int) error
}
