package presenter

import (
	"time"

	"github.com/yohanalexander/deezefy-music/entity"
)

// PostUsuario entidade Usuario
type PostUsuario struct {
	Email    string `json:"email"`
	Password string `json:"senha"`
	Birthday string `json:"data_nascimento"`
}

// MakeUsuario seta os valores a partir da entidade
func (u *PostUsuario) MakeUsuario(usuario entity.Usuario) {
	u.Email = usuario.Email
	u.Password = usuario.Password
	u.Birthday = usuario.Birthday
}

// Usuario presenter Usuario
type Usuario struct {
	Email    string `json:"email"`
	Password string `json:"senha"`
	Birthday string `json:"data_nascimento"`
	Idade    int    `json:"idade"`
}

// AppendUsuario adiciona presenter na lista
func AppendUsuario(usuario entity.Usuario, usuarios []Usuario) []Usuario {
	u := &Usuario{}
	u.MakeUsuario(usuario)
	usuarios = append(usuarios, *u)
	return usuarios
}

// MakeUsuario seta os valores a partir da entidade
func (u *Usuario) MakeUsuario(usuario entity.Usuario) {
	u.Email = usuario.Email
	u.Password = usuario.Password
	u.Birthday = usuario.Birthday
	u.Idade = u.GetIdade(time.Now())
}

// GetIdade calcula a idade do usuario
func (u *Usuario) GetIdade(now time.Time) int {

	birthDate, err := parseBirthday(u.Birthday)

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

// parseBirthday cast da string birthday para time.Time
func parseBirthday(birthday string) (time.Time, error) {

	layout := "2006-01-02T00:00:00Z"

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
