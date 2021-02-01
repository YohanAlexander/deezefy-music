package grava

import der "github.com/yohanalexander/deezefy-music/entity/grava"

// Grava interface
type Grava interface {
	Get(musica int, artista string) (*der.Grava, error)
	GetByMusica(musica int) (*der.Grava, error)
	GetByArtista(artista string) (*der.Grava, error)
	Search(query string) ([]*der.Grava, error)
	List() ([]*der.Grava, error)
	Create(e *der.Grava) (int, string, error)
	Update(e *der.Grava) error
	Delete(musica int, artista string) error
}

// Repository interface
type Repository interface {
	Grava
}

// UseCase interface
type UseCase interface {
	GetGrava(musica int, artista string) (*der.Grava, error)
	GetGravaByMusica(Musica int) (*der.Grava, error)
	GetGravaByartista(artista string) (*der.Grava, error)
	SearchGravas(query string) ([]*der.Grava, error)
	ListGravas() ([]*der.Grava, error)
	CreateGrava(musica int, artista string) (int, string, error)
	UpdateGrava(e *der.Grava) error
	DeleteGrava(musica int, artista string) error
}
