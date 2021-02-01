package artistagenero

import der "github.com/yohanalexander/deezefy-music/entity/artistagenero"

// ArtistaGenero interface
type ArtistaGenero interface {
	Get(artista, genero string) (*der.ArtistaGenero, error)
	GetByArtista(artista string) (*der.ArtistaGenero, error)
	GetByGenero(genero string) (*der.ArtistaGenero, error)
	Search(query string) ([]*der.ArtistaGenero, error)
	List() ([]*der.ArtistaGenero, error)
	Create(e *der.ArtistaGenero) (string, string, error)
	Update(e *der.ArtistaGenero) error
	Delete(artista, genero string) error
}

// Repository interface
type Repository interface {
	ArtistaGenero
}

// UseCase interface
type UseCase interface {
	GetArtistaGenero(artista, genero string) (*der.ArtistaGenero, error)
	GetArtistaGeneroByArtista(artista string) (*der.ArtistaGenero, error)
	GetArtistaGeneroBygenero(genero string) (*der.ArtistaGenero, error)
	SearchArtistaGeneros(query string) ([]*der.ArtistaGenero, error)
	ListArtistaGeneros() ([]*der.ArtistaGenero, error)
	CreateArtistaGenero(artista, genero string) (string, string, error)
	UpdateArtistaGenero(e *der.ArtistaGenero) error
	DeleteArtistaGenero(artista, genero string) error
}
