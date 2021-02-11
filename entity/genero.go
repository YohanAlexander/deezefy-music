package entity

import (
	"github.com/go-playground/validator/v10"
)

// Genero entidade Genero
type Genero struct {
	Nome     string    `json:"nome" validate:"required,gte=1"`
	Estilo   string    `json:"estilo" validate:"required,oneof=blues rock mpb samba sertanejo jazz classica"`
	Artistas []Artista `json:"artistas" validate:""`
	Musicas  []Musica  `json:"musicas" validate:""`
	Perfis   []Perfil  `json:"perfis" validate:""`
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

// AddMusica adiciona um Musica
func (g *Genero) AddMusica(musica Musica) error {
	_, err := g.GetMusica(musica)
	if err == nil {
		return ErrMusicaRegistered
	}
	g.Musicas = append(g.Musicas, musica)
	return nil
}

// RemoveMusica remove um Musica
func (g *Genero) RemoveMusica(musica Musica) error {
	for i, j := range g.Musicas {
		if j.ID == musica.ID {
			g.Musicas = append(g.Musicas[:i], g.Musicas[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// GetMusica get a Musica
func (g *Genero) GetMusica(musica Musica) (Musica, error) {
	for _, v := range g.Musicas {
		if v.ID == musica.ID {
			return musica, nil
		}
	}
	return musica, ErrNotFound
}

// AddPerfil adiciona um Perfil
func (g *Genero) AddPerfil(perfil Perfil) error {
	_, err := g.GetPerfil(perfil)
	if err == nil {
		return ErrPerfilRegistered
	}
	g.Perfis = append(g.Perfis, perfil)
	return nil
}

// RemovePerfil remove um Perfil
func (g *Genero) RemovePerfil(perfil Perfil) error {
	for i, j := range g.Perfis {
		if j.ID == perfil.ID {
			g.Perfis = append(g.Perfis[:i], g.Perfis[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// GetPerfil get a Perfil
func (g *Genero) GetPerfil(perfil Perfil) (Perfil, error) {
	for _, v := range g.Perfis {
		if v.ID == perfil.ID {
			return perfil, nil
		}
	}
	return perfil, ErrNotFound
}
