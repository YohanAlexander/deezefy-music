package musicagenero

import der "github.com/yohanalexander/deezefy-music/entity/musicagenero"

// MusicaGenero interface
type MusicaGenero interface {
	Get(musica int, genero string) (*der.MusicaGenero, error)
	GetByMusica(musica int) (*der.MusicaGenero, error)
	GetByGenero(genero string) (*der.MusicaGenero, error)
	Search(query string) ([]*der.MusicaGenero, error)
	List() ([]*der.MusicaGenero, error)
	Create(e *der.MusicaGenero) (int, string, error)
	Update(e *der.MusicaGenero) error
	Delete(musica int, genero string) error
}

// Repository interface
type Repository interface {
	MusicaGenero
}

// UseCase interface
type UseCase interface {
	GetMusicaGenero(musica int, genero string) (*der.MusicaGenero, error)
	GetMusicaGeneroByMusica(Musica int) (*der.MusicaGenero, error)
	GetMusicaGeneroByGenero(genero string) (*der.MusicaGenero, error)
	SearchMusicaGeneros(query string) ([]*der.MusicaGenero, error)
	ListMusicaGeneros() ([]*der.MusicaGenero, error)
	CreateMusicaGenero(musica int, genero string) (int, string, error)
	UpdateMusicaGenero(e *der.MusicaGenero) error
	DeleteMusicaGenero(musica int, genero string) error
}
