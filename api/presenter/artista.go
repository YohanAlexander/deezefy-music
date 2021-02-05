package presenter

import (
	"time"
)

// Artista entidade Artista
type Artista struct {
	Usuario       Usuario   `json:"usuario"`
	NomeArtistico string    `json:"nomeartistico"`
	Biografia     string    `json:"biografia"`
	AnoFormacao   int       `json:"anoformacao"`
	Seguidores    []Ouvinte `json:"seguidores"`
	Musicas       []Musica  `json:"musicas"`
	Perfis        []Perfil  `json:"perfis"`
	Generos       []Genero  `json:"generos"`
	Albums        []Album   `json:"albums"`
	Idade         int       `json:"idade"`
}

// GetIdade calcula a idade do usuario
func (a *Artista) GetIdade(now time.Time) {

	birthDate, err := parseBirthday(a.Usuario.Birthday)

	if err != nil {
		a.Idade = 0
	}

	age := now.Year() - birthDate.Year()

	birthDay := getAdjustedBirthDay(birthDate, now)
	if now.YearDay() < birthDay {
		age--
	}

	a.Idade = age

}
