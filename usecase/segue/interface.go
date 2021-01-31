package segue

import der "github.com/yohanalexander/deezefy-music/entity/segue"

// Segue interface
type Segue interface {
	Get(artista, ouvinte string) (*der.Segue, error)
	GetByArtista(artista string) (*der.Segue, error)
	GetByOuvinte(ouvinte string) (*der.Segue, error)
	Search(query string) ([]*der.Segue, error)
	List() ([]*der.Segue, error)
	Create(e *der.Segue) (string, string, error)
	Update(e *der.Segue) error
	Delete(artista, ouvinte string) error
}

// Repository interface
type Repository interface {
	Segue
}

// UseCase interface
type UseCase interface {
	GetSegue(artista, ouvinte string) (*der.Segue, error)
	GetSegueByArtista(artista string) (*der.Segue, error)
	GetSegueByOuvinte(ouvinte string) (*der.Segue, error)
	SearchSegues(query string) ([]*der.Segue, error)
	ListSegues() ([]*der.Segue, error)
	CreateSegue(artista, ouvinte string) (string, string, error)
	UpdateSegue(e *der.Segue) error
	DeleteSegue(artista, ouvinte string) error
}
