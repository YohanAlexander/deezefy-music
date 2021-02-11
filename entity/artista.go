package entity

import (
	"github.com/go-playground/validator/v10"
)

// Artista entidade Artista
type Artista struct {
	Usuario       Usuario   `json:"usuario" validate:"required"`
	NomeArtistico string    `json:"nome_artistico" validate:"required,gte=1"`
	Biografia     string    `json:"biografia" validate:"gte=1"`
	AnoFormacao   int       `json:"ano_formacao" validate:"gte=1000"`
	Organizador   []Evento  `json:"eventos" validate:""`
	Seguidores    []Ouvinte `json:"seguidores" validate:""`
	Grava         []Musica  `json:"musicas" validate:""`
	Perfis        []Perfil  `json:"perfis" validate:""`
	Generos       []Genero  `json:"generos" validate:""`
	Albums        []Album   `json:"albums" validate:""`
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

// AddEvento adiciona um Evento
func (a *Artista) AddEvento(evento Evento) error {
	_, err := a.GetEvento(evento)
	if err == nil {
		return ErrEventoRegistered
	}
	a.Organizador = append(a.Organizador, evento)
	return nil
}

// RemoveEvento remove um Evento
func (a *Artista) RemoveEvento(evento Evento) error {
	for i, j := range a.Organizador {
		if j.ID == evento.ID {
			a.Organizador = append(a.Organizador[:i], a.Organizador[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// GetEvento get a Evento
func (a *Artista) GetEvento(evento Evento) (Evento, error) {
	for _, v := range a.Organizador {
		if v.ID == evento.ID {
			return evento, nil
		}
	}
	return evento, ErrNotFound
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

// AddGenero adiciona um Genero
func (a *Artista) AddGenero(genero Genero) error {
	_, err := a.GetGenero(genero)
	if err == nil {
		return ErrGeneroRegistered
	}
	a.Generos = append(a.Generos, genero)
	return nil
}

// RemoveGenero remove um Genero
func (a *Artista) RemoveGenero(genero Genero) error {
	for i, j := range a.Generos {
		if j.Nome == genero.Nome {
			a.Generos = append(a.Generos[:i], a.Generos[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// GetGenero get a Genero
func (a *Artista) GetGenero(genero Genero) (Genero, error) {
	for _, v := range a.Generos {
		if v.Nome == genero.Nome {
			return genero, nil
		}
	}
	return genero, ErrNotFound
}

// AddAlbum adiciona um Album
func (a *Artista) AddAlbum(album Album) error {
	_, err := a.GetAlbum(album)
	if err == nil {
		return ErrArtistaRegistered
	}
	a.Albums = append(a.Albums, album)
	return nil
}

// RemoveAlbum remove um Album
func (a *Artista) RemoveAlbum(album Album) error {
	for i, j := range a.Albums {
		if j.ID == album.ID {
			a.Albums = append(a.Albums[:i], a.Albums[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// GetAlbum get a Album
func (a *Artista) GetAlbum(album Album) (Album, error) {
	for _, v := range a.Albums {
		if v.ID == album.ID {
			return album, nil
		}
	}
	return album, ErrNotFound
}
