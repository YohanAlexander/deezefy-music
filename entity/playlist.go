package entity

import (
	"github.com/go-playground/validator/v10"
)

// Playlist entidade Playlist
type Playlist struct {
	Nome   string    `validate:"required,gte=1"`
	Status string    `validate:"required,oneof=ativo inativo"`
	Salvou []Ouvinte `validate:""`
}

// NewPlaylist cria um novo Playlist
func NewPlaylist(nome, status string) (*Playlist, error) {
	p := &Playlist{
		Nome:   nome,
		Status: status,
	}
	err := p.Validate()
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Validate valida os dados do Playlist
func (p *Playlist) Validate() error {
	vld := validator.New()
	if err := vld.Struct(p); err != nil {
		return ErrInvalidEntity
	}
	return nil
}

// AddOuvinte adiciona um Ouvinte
func (p *Playlist) AddOuvinte(ouvinte Ouvinte) error {
	_, err := p.GetOuvinte(ouvinte)
	if err == nil {
		return ErrOuvinteRegistered
	}
	p.Salvou = append(p.Salvou, ouvinte)
	return nil
}

// RemoveOuvinte remove um Ouvinte
func (p *Playlist) RemoveOuvinte(ouvinte Ouvinte) error {
	for i, j := range p.Salvou {
		if j.Usuario.Email == ouvinte.Usuario.Email {
			p.Salvou = append(p.Salvou[:i], p.Salvou[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// GetOuvinte get a Ouvinte
func (p *Playlist) GetOuvinte(ouvinte Ouvinte) (Ouvinte, error) {
	for _, v := range p.Salvou {
		if v.Usuario.Email == ouvinte.Usuario.Email {
			return ouvinte, nil
		}
	}
	return ouvinte, ErrNotFound
}
