package entity

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

// Usuario entidade usuario
type Usuario struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"senha" validate:"required,gte=8"`
	Birthday string `json:"data_nascimento" validate:"datetime=2006-01-02"`
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
