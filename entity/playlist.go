package entity

import (
	"github.com/go-playground/validator/v10"
)

// Playlist entidade Playlist
type Playlist struct {
	Usuario     Usuario   `json:"usuario" validate:"required"`
	Nome        string    `json:"nome" validate:"required,gte=1"`
	Status      string    `json:"status" validate:"required,oneof=ativo inativo"`
	DataCriacao string    `json:"data_criacao" validate:"required,datetime=2006-01-02"`
	Salvou      []Ouvinte `json:"ouvintes" validate:""`
	Musicas     []Musica  `json:"musicas" validate:""`
}

// NewPlaylist cria um novo Playlist
func NewPlaylist(email, password, birthday, nome, status, datacriacao string) (*Playlist, error) {
	u, err := NewUsuario(email, password, birthday)
	if err != nil {
		return nil, err
	}
	p := &Playlist{
		Usuario:     *u,
		Nome:        nome,
		Status:      status,
		DataCriacao: datacriacao,
	}
	err = p.Validate()
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

// AddMusica adiciona um Musica
func (p *Playlist) AddMusica(musica Musica) error {
	_, err := p.GetMusica(musica)
	if err == nil {
		return ErrMusicaRegistered
	}
	p.Musicas = append(p.Musicas, musica)
	return nil
}

// RemoveMusica remove um Musica
func (p *Playlist) RemoveMusica(musica Musica) error {
	for i, j := range p.Musicas {
		if j.ID == musica.ID {
			p.Musicas = append(p.Musicas[:i], p.Musicas[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// GetMusica get a Musica
func (p *Playlist) GetMusica(musica Musica) (Musica, error) {
	for _, v := range p.Musicas {
		if v.ID == musica.ID {
			return musica, nil
		}
	}
	return musica, ErrNotFound
}
