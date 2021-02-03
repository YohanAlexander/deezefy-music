package entity

import (
	"github.com/go-playground/validator/v10"
)

// Musica entidade Musica
type Musica struct {
	ID      int       `validate:"required,gte=1"`
	Duracao int       `validate:"required,gte=100"`
	Nome    string    `validate:"required,gte=1"`
	Curtiu  []Ouvinte `validate:""`
	Gravou  []Artista `validate:""`
}

// NewMusica cria um novo Musica
func NewMusica(id, duracao int, nome string) (*Musica, error) {
	m := &Musica{
		ID:      id,
		Duracao: duracao,
		Nome:    nome,
	}
	err := m.Validate()
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Validate valida os dados do Musica
func (m *Musica) Validate() error {
	vld := validator.New()
	if err := vld.Struct(m); err != nil {
		return ErrInvalidEntity
	}
	return nil
}

// AddOuvinte adiciona um Ouvinte
func (m *Musica) AddOuvinte(ouvinte Ouvinte) error {
	_, err := m.GetOuvinte(ouvinte)
	if err == nil {
		return ErrOuvinteRegistered
	}
	m.Curtiu = append(m.Curtiu, ouvinte)
	return nil
}

// RemoveOuvinte remove um Ouvinte
func (m *Musica) RemoveOuvinte(ouvinte Ouvinte) error {
	for i, j := range m.Curtiu {
		if j.Usuario.Email == ouvinte.Usuario.Email {
			m.Curtiu = append(m.Curtiu[:i], m.Curtiu[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// GetOuvinte get a Ouvinte
func (m *Musica) GetOuvinte(ouvinte Ouvinte) (Ouvinte, error) {
	for _, v := range m.Curtiu {
		if v.Usuario.Email == ouvinte.Usuario.Email {
			return ouvinte, nil
		}
	}
	return ouvinte, ErrNotFound
}

// AddArtista adiciona um Artista
func (m *Musica) AddArtista(artista Artista) error {
	_, err := m.GetArtista(artista)
	if err == nil {
		return ErrArtistaRegistered
	}
	m.Gravou = append(m.Gravou, artista)
	return nil
}

// RemoveArtista remove um Artista
func (m *Musica) RemoveArtista(artista Artista) error {
	for i, j := range m.Gravou {
		if j.Usuario.Email == artista.Usuario.Email {
			m.Gravou = append(m.Gravou[:i], m.Gravou[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// GetArtista get a Artista
func (m *Musica) GetArtista(artista Artista) (Artista, error) {
	for _, v := range m.Gravou {
		if v.Usuario.Email == artista.Usuario.Email {
			return artista, nil
		}
	}
	return artista, ErrNotFound
}
