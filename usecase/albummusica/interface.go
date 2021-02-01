package albummusica

import der "github.com/yohanalexander/deezefy-music/entity/albummusica"

// AlbumMusica interface
type AlbumMusica interface {
	Get(album, musica int, artista string) (*der.AlbumMusica, error)
	GetByAlbum(album int) (*der.AlbumMusica, error)
	GetByMusica(musica int) (*der.AlbumMusica, error)
	Search(query string) ([]*der.AlbumMusica, error)
	List() ([]*der.AlbumMusica, error)
	Create(e *der.AlbumMusica) (int, int, string, error)
	Update(e *der.AlbumMusica) error
	Delete(album, musica int, artista string) error
}

// Repository interface
type Repository interface {
	AlbumMusica
}

// UseCase interface
type UseCase interface {
	GetAlbumMusica(album, musica int, artista string) (*der.AlbumMusica, error)
	GetAlbumMusicaByAlbum(album int) (*der.AlbumMusica, error)
	GetAlbumMusicaByMusica(musica int) (*der.AlbumMusica, error)
	SearchAlbumMusicas(query string) ([]*der.AlbumMusica, error)
	ListAlbumMusicas() ([]*der.AlbumMusica, error)
	CreateAlbumMusica(album, musica int, artista string) (int, int, string, error)
	UpdateAlbumMusica(e *der.AlbumMusica) error
	DeleteAlbumMusica(album, musica int, artista string) error
}
