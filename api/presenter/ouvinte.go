package presenter

import (
	"time"

	"github.com/yohanalexander/deezefy-music/entity"
)

// Ouvinte entidade Ouvinte
type Ouvinte struct {
	Usuario      entity.Usuario    `json:"usuario"`
	PrimeiroNome string            `json:"primeironome"`
	Sobrenome    string            `json:"sobrenome"`
	Telefones    []string          `json:"telefones"`
	Seguindo     []entity.Artista  `json:"seguindo"`
	Curtidas     []entity.Musica   `json:"curtidas"`
	Playlists    []entity.Playlist `json:"playlists"`
	Albums       []entity.Album    `json:"albums"`
	Idade        int               `json:"idade"`
}

// GetOuvinte seta os valores a partir da entidade
func (o *Ouvinte) GetOuvinte(ouvinte entity.Ouvinte) {
	o.Usuario = ouvinte.Usuario
	o.PrimeiroNome = ouvinte.PrimeiroNome
	o.Sobrenome = ouvinte.Sobrenome
	o.Telefones = ouvinte.Telefones
	o.Seguindo = ouvinte.Seguindo
	o.Curtidas = ouvinte.Curtidas
	o.Playlists = ouvinte.Playlists
	o.Albums = ouvinte.Albums
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
