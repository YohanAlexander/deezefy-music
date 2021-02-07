package presenter

import (
	"time"

	"github.com/yohanalexander/deezefy-music/entity"
)

// Usuario entidade usuario
type Usuario struct {
	Email       string            `json:"email"`
	Password    string            `json:"-"`
	Birthday    string            `json:"datanascimento"`
	Organizador []entity.Evento   `json:"eventos"`
	Cria        []entity.Playlist `json:"playlists"`
	Idade       int               `json:"idade"`
}

// GetUsuario seta os valores a partir da entidade
func (u *Usuario) GetUsuario(usuario entity.Usuario) {
	u.Email = usuario.Email
	u.Password = usuario.Password
	u.Birthday = usuario.Birthday
	u.Organizador = usuario.Organizador
	u.Cria = usuario.Cria
}

// GetIdade calcula a idade do usuario
func (u *Usuario) GetIdade(now time.Time) {

	birthDate, err := parseBirthday(u.Birthday)

	if err != nil {
		u.Idade = 0
	}

	age := now.Year() - birthDate.Year()

	birthDay := getAdjustedBirthDay(birthDate, now)
	if now.YearDay() < birthDay {
		age--
	}

	u.Idade = age

}

// parseBirthday cast da string birthday para time.Time
func parseBirthday(birthday string) (time.Time, error) {

	layout := "2006-01-02"

	date, err := time.Parse(layout, birthday)

	if err != nil {
		return date, err
	}

	return date, nil

}

// getAdjustedBirthDay verifica as diferenças em ano bissexto
func getAdjustedBirthDay(birthDate time.Time, now time.Time) int {
	birthDay := birthDate.YearDay()
	currentDay := now.YearDay()
	if isLeap(birthDate) && !isLeap(now) && birthDay >= 60 {
		return birthDay - 1
	}
	if isLeap(now) && !isLeap(birthDate) && currentDay >= 60 {
		return birthDay + 1
	}
	return birthDay
}

// isLeap verifica se o ano é bissexto
func isLeap(date time.Time) bool {
	year := date.Year()
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}
