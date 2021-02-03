package entity

import (
	"github.com/go-playground/validator/v10"
)

// Artista entidade Artista
type Artista struct {
	Usuario       Usuario   `validate:"required"`
	NomeArtistico string    `validate:"required,gte=1"`
	Biografia     string    `validate:"gte=1"`
	AnoFormacao   int       `validate:"gte=1000"`
	Seguidores    []Ouvinte `validate:""`
	Grava         []Musica  `validate:""`
	Perfis        []Perfil  `validate:""`
}

// NewArtista cria um novo Artista
func NewArtista(email, password, birthday, nomeartistico, biografia string, anoformacao int) (*Artista, error) {
	u, err := NewUsuario(email, password, birthday)
	if err != nil {
		return nil, err
	}
	a := &Artista{
		Usuario:       *u,
		NomeArtistico: nomeartistico,
		Biografia:     biografia,
		AnoFormacao:   anoformacao,
	}
	err = a.Validate()
	if err != nil {
		return nil, err
	}
	return a, nil
}

// Validate valida os dados do Artista
func (a *Artista) Validate() error {
	vld := validator.New()
	if err := vld.Struct(a); err != nil {
		return ErrInvalidEntity
	}
	return nil
}

// AddOuvinte adiciona um Ouvinte
func (a *Artista) AddOuvinte(ouvinte Ouvinte) error {
	_, err := a.GetOuvinte(ouvinte)
	if err == nil {
		return ErrOuvinteRegistered
	}
	a.Seguidores = append(a.Seguidores, ouvinte)
	return nil
}

// RemoveOuvinte remove um Ouvinte
func (a *Artista) RemoveOuvinte(ouvinte Ouvinte) error {
	for i, j := range a.Seguidores {
		if j.Usuario.Email == ouvinte.Usuario.Email {
			a.Seguidores = append(a.Seguidores[:i], a.Seguidores[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// GetOuvinte get a Ouvinte
func (a *Artista) GetOuvinte(ouvinte Ouvinte) (Ouvinte, error) {
	for _, v := range a.Seguidores {
		if v.Usuario.Email == ouvinte.Usuario.Email {
			return ouvinte, nil
		}
	}
	return ouvinte, ErrNotFound
}

// AddMusica adiciona um Musica
func (a *Artista) AddMusica(musica Musica) error {
	_, err := a.GetMusica(musica)
	if err == nil {
		return ErrMusicaRegistered
	}
	a.Grava = append(a.Grava, musica)
	return nil
}

// RemoveMusica remove um Musica
func (a *Artista) RemoveMusica(musica Musica) error {
	for i, j := range a.Grava {
		if j.ID == musica.ID {
			a.Grava = append(a.Grava[:i], a.Grava[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// GetMusica get a Musica
func (a *Artista) GetMusica(musica Musica) (Musica, error) {
	for _, v := range a.Grava {
		if v.ID == musica.ID {
			return musica, nil
		}
	}
	return musica, ErrNotFound
}

// AddPerfil adiciona um Perfil
func (a *Artista) AddPerfil(perfil Perfil) error {
	_, err := a.GetPerfil(perfil)
	if err == nil {
		return ErrPerfilRegistered
	}
	a.Perfis = append(a.Perfis, perfil)
	return nil
}

// RemovePerfil remove um Perfil
func (a *Artista) RemovePerfil(perfil Perfil) error {
	for i, j := range a.Perfis {
		if j.ID == perfil.ID {
			a.Perfis = append(a.Perfis[:i], a.Perfis[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// GetPerfil get a Perfil
func (a *Artista) GetPerfil(perfil Perfil) (Perfil, error) {
	for _, v := range a.Perfis {
		if v.ID == perfil.ID {
			return perfil, nil
		}
	}
	return perfil, ErrNotFound
}
