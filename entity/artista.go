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
