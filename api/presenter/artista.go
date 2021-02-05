package presenter

import (
	"time"

	"github.com/yohanalexander/deezefy-music/entity"
)

// Artista entidade Artista
type Artista struct {
	Usuario       entity.Usuario   `json:"usuario"`
	NomeArtistico string           `json:"nomeartistico"`
	Biografia     string           `json:"biografia"`
	AnoFormacao   int              `json:"anoformacao"`
	Seguidores    []entity.Ouvinte `json:"seguidores"`
	Musicas       []entity.Musica  `json:"musicas"`
	Perfis        []entity.Perfil  `json:"perfis"`
	Generos       []entity.Genero  `json:"generos"`
	Albums        []entity.Album   `json:"albums"`
	Idade         int              `json:"idade"`
}

// GetArtista seta os valores a partir da entidade
func (a *Artista) GetArtista(artista entity.Artista) {
	a.Usuario = artista.Usuario
	a.NomeArtistico = artista.NomeArtistico
	a.Biografia = artista.Biografia
	a.AnoFormacao = artista.AnoFormacao
	a.Seguidores = artista.Seguidores
	a.Musicas = artista.Grava
	a.Perfis = artista.Perfis
	a.Generos = artista.Generos
	a.Albums = artista.Albums
	a.Idade = a.GetIdade(time.Now())
}

// GetIdade calcula a idade do usuario
func (a *Artista) GetIdade(now time.Time) int {

	birthDate, err := parseBirthday(a.Usuario.Birthday)

	if err != nil {
		return 1
	}

	age := now.Year() - birthDate.Year()

	birthDay := getAdjustedBirthDay(birthDate, now)
	if now.YearDay() < birthDay {
		age--
	}

	return age

}
