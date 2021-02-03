package entity

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

// Usuario entidade usuario
type Usuario struct {
	Email       string   `validate:"required,email"`
	Password    string   `validate:"required,gte=8"`
	Birthday    string   `validate:"datetime=2006-01-02"`
	Organizador []Evento `validate:""`
}

// NewUsuario cria um novo usuario
func NewUsuario(email, password, birthday string) (*Usuario, error) {
	u := &Usuario{
		Email:    email,
		Password: password,
		Birthday: birthday,
	}
	err := u.Validate()
	if err != nil {
		return nil, err
	}
	pwd, err := generatePassword(password)
	if err != nil {
		return nil, err
	}
	u.Password = pwd
	return u, nil
}

// Validate valida os dados do usuario
func (u *Usuario) Validate() error {
	vld := validator.New()
	if err := vld.Struct(u); err != nil {
		return ErrInvalidEntity
	}
	return nil
}

// ValidatePassword valida a senha do usuario
func (u *Usuario) ValidatePassword(p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	if err != nil {
		return err
	}
	return nil
}

func generatePassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// AddEvento adiciona um Evento
func (u *Usuario) AddEvento(evento Evento) error {
	_, err := u.GetEvento(evento)
	if err == nil {
		return ErrEventoRegistered
	}
	u.Organizador = append(u.Organizador, evento)
	return nil
}

// RemoveEvento remove um Evento
func (u *Usuario) RemoveEvento(evento Evento) error {
	for i, j := range u.Organizador {
		if j.Usuario.Email == evento.Usuario.Email {
			u.Organizador = append(u.Organizador[:i], u.Organizador[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// GetEvento get a Evento
func (u *Usuario) GetEvento(evento Evento) (Evento, error) {
	for _, v := range u.Organizador {
		if v.Usuario.Email == evento.Usuario.Email {
			return evento, nil
		}
	}
	return evento, ErrNotFound
}
