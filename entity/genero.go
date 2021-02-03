package entity

import (
	"github.com/go-playground/validator/v10"
)

// Genero entidade Genero
type Genero struct {
	Nome     string    `validate:"required,gte=1"`
	Estilo   string    `validate:"required,oneof=blues rock mpb samba sertanejo jazz classica"`
	Artistas []Artista `validate:""`
}

// NewGenero cria um novo Genero
func NewGenero(nome, estilo string) (*Genero, error) {
	g := &Genero{
		Nome:   nome,
		Estilo: estilo,
	}
	err := g.Validate()
	if err != nil {
		return nil, err
	}
	return g, nil
}

// Validate valida os dados do Genero
func (g *Genero) Validate() error {
	vld := validator.New()
	if err := vld.Struct(g); err != nil {
		return ErrInvalidEntity
	}
	return nil
}

// AddArtista adiciona um Artista
func (g *Genero) AddArtista(artista Artista) error {
	_, err := g.GetArtista(artista)
	if err == nil {
		return ErrArtistaRegistered
	}
	g.Artistas = append(g.Artistas, artista)
	return nil
}

// RemoveArtista remove um Artista
func (g *Genero) RemoveArtista(artista Artista) error {
	for i, j := range g.Artistas {
		if j.Usuario.Email == artista.Usuario.Email {
			g.Artistas = append(g.Artistas[:i], g.Artistas[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// GetArtista get a Artista
func (g *Genero) GetArtista(artista Artista) (Artista, error) {
	for _, v := range g.Artistas {
		if v.Usuario.Email == artista.Usuario.Email {
			return artista, nil
		}
	}
	return artista, ErrNotFound
}
