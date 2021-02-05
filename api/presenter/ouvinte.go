package presenter

import (
	"time"
)

// Ouvinte entidade Ouvinte
type Ouvinte struct {
	Usuario      Usuario    `json:"usuario"`
	PrimeiroNome string     `json:"primeironome"`
	Sobrenome    string     `json:"sobrenome"`
	Telefones    []string   `json:"telefones"`
	Seguindo     []Artista  `json:"seguindo"`
	Curtidas     []Musica   `json:"curtidas"`
	Playlists    []Playlist `json:"playlists"`
	Albums       []Album    `json:"albums"`
	Idade        int        `json:"idade"`
}

// GetIdade calcula a idade do usuario
func (o *Ouvinte) GetIdade(now time.Time) {

	birthDate, err := parseBirthday(o.Usuario.Birthday)

	if err != nil {
		o.Idade = 0
	}

	age := now.Year() - birthDate.Year()

	birthDay := getAdjustedBirthDay(birthDate, now)
	if now.YearDay() < birthDay {
		age--
	}

	o.Idade = age

}
