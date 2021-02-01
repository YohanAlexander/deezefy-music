package artistasfavoritos

import der "github.com/yohanalexander/deezefy-music/entity/artistasfavoritos"

// ArtistasFavoritos interface
type ArtistasFavoritos interface {
	Get(perfil int, ouvinte, artista string) (*der.ArtistasFavoritos, error)
	GetByPerfil(perfil int) (*der.ArtistasFavoritos, error)
	GetByOuvinte(ouvinte string) (*der.ArtistasFavoritos, error)
	Search(query string) ([]*der.ArtistasFavoritos, error)
	List() ([]*der.ArtistasFavoritos, error)
	Create(e *der.ArtistasFavoritos) (int, string, string, error)
	Update(e *der.ArtistasFavoritos) error
	Delete(perfil int, ouvinte, artista string) error
}

// Repository interface
type Repository interface {
	ArtistasFavoritos
}

// UseCase interface
type UseCase interface {
	GetArtistasFavoritos(perfil int, ouvinte, artista string) (*der.ArtistasFavoritos, error)
	GetArtistasFavoritosByPerfil(perfil int) (*der.ArtistasFavoritos, error)
	GetArtistasFavoritosByOuvinte(ouvinte string) (*der.ArtistasFavoritos, error)
	SearchArtistasFavoritoss(query string) ([]*der.ArtistasFavoritos, error)
	ListArtistasFavoritoss() ([]*der.ArtistasFavoritos, error)
	CreateArtistasFavoritos(perfil int, ouvinte, artista string) (int, string, string, error)
	UpdateArtistasFavoritos(e *der.ArtistasFavoritos) error
	DeleteArtistasFavoritos(perfil int, ouvinte, artista string) error
}
