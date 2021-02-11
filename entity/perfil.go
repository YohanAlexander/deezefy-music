package entity

import (
	"github.com/go-playground/validator/v10"
)

// Perfil entidade Perfil
type Perfil struct {
	Ouvinte               Ouvinte   `json:"ouvinte" validate:"required"`
	ID                    int       `json:"id" validate:"required,gte=1"`
	InformacoesRelevantes string    `json:"informacoes_relevantes" validate:"required,gte=1"`
	ArtistasFavoritos     []Artista `json:"artistas_favoritos" validate:""`
	GenerosFavoritos      []Genero  `json:"generos_favoritos" validate:""`
}

// NewPerfil cria um novo Perfil
func NewPerfil(email, password, birthday, primeironome, sobrenome, informacoesrelevantes string, id int) (*Perfil, error) {
	o, err := NewOuvinte(email, password, birthday, primeironome, sobrenome)
	if err != nil {
		return nil, err
	}
	p := &Perfil{
		ID:                    id,
		Ouvinte:               *o,
		InformacoesRelevantes: informacoesrelevantes,
	}
	err = p.Validate()
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Validate valida os dados do Perfil
func (p *Perfil) Validate() error {
	vld := validator.New()
	if err := vld.Struct(p); err != nil {
		return ErrInvalidEntity
	}
	return nil
}

// AddArtista adiciona um Artista
func (p *Perfil) AddArtista(artista Artista) error {
	_, err := p.GetArtista(artista)
	if err == nil {
		return ErrArtistaRegistered
	}
	p.ArtistasFavoritos = append(p.ArtistasFavoritos, artista)
	return nil
}

// RemoveArtista remove um Artista
func (p *Perfil) RemoveArtista(artista Artista) error {
	for i, j := range p.ArtistasFavoritos {
		if j.Usuario.Email == artista.Usuario.Email {
			p.ArtistasFavoritos = append(p.ArtistasFavoritos[:i], p.ArtistasFavoritos[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// GetArtista get a Artista
func (p *Perfil) GetArtista(artista Artista) (Artista, error) {
	for _, v := range p.ArtistasFavoritos {
		if v.Usuario.Email == artista.Usuario.Email {
			return artista, nil
		}
	}
	return artista, ErrNotFound
}

// AddGenero adiciona um Genero
func (p *Perfil) AddGenero(genero Genero) error {
	_, err := p.GetGenero(genero)
	if err == nil {
		return ErrGeneroRegistered
	}
	p.GenerosFavoritos = append(p.GenerosFavoritos, genero)
	return nil
}

// RemoveGenero remove um Genero
func (p *Perfil) RemoveGenero(genero Genero) error {
	for i, j := range p.GenerosFavoritos {
		if j.Nome == genero.Nome {
			p.GenerosFavoritos = append(p.GenerosFavoritos[:i], p.GenerosFavoritos[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// GetGenero get a Genero
func (p *Perfil) GetGenero(genero Genero) (Genero, error) {
	for _, v := range p.GenerosFavoritos {
		if v.Nome == genero.Nome {
			return genero, nil
		}
	}
	return genero, ErrNotFound
}
