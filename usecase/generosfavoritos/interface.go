package generosfavoritos

import der "github.com/yohanalexander/deezefy-music/entity/generosfavoritos"

// GenerosFavoritos interface
type GenerosFavoritos interface {
	Get(perfil int, genero, ouvinte string) (*der.GenerosFavoritos, error)
	GetByPerfil(perfil int) (*der.GenerosFavoritos, error)
	GetByGenero(genero string) (*der.GenerosFavoritos, error)
	Search(query string) ([]*der.GenerosFavoritos, error)
	List() ([]*der.GenerosFavoritos, error)
	Create(e *der.GenerosFavoritos) (int, string, string, error)
	Update(e *der.GenerosFavoritos) error
	Delete(perfil int, genero, ouvinte string) error
}

// Repository interface
type Repository interface {
	GenerosFavoritos
}

// UseCase interface
type UseCase interface {
	GetGenerosFavoritos(perfil int, genero, ouvinte string) (*der.GenerosFavoritos, error)
	GetGenerosFavoritosByPerfil(perfil int) (*der.GenerosFavoritos, error)
	GetGenerosFavoritosByGenero(genero string) (*der.GenerosFavoritos, error)
	SearchGenerosFavoritoss(query string) ([]*der.GenerosFavoritos, error)
	ListGenerosFavoritoss() ([]*der.GenerosFavoritos, error)
	CreateGenerosFavoritos(perfil int, genero, ouvinte string) (int, string, string, error)
	UpdateGenerosFavoritos(e *der.GenerosFavoritos) error
	DeleteGenerosFavoritos(perfil int, genero, ouvinte string) error
}
