package entity

import (
	"github.com/go-playground/validator/v10"
)

// Ouvinte entidade Ouvinte
type Ouvinte struct {
	Usuario      Usuario  `validate:"required"`
	PrimeiroNome string   `validate:"required"`
	Sobrenome    string   `validate:"required"`
	Telefones    []string `validate:""`
}

// NewOuvinte cria um novo Ouvinte
func NewOuvinte(email, password, birthday, primeironome, sobrenome string) (*Ouvinte, error) {
	u, err := NewUsuario(email, password, birthday)
	if err != nil {
		return nil, err
	}
	o := &Ouvinte{
		Usuario:      *u,
		PrimeiroNome: primeironome,
		Sobrenome:    sobrenome,
	}
	err = o.Validate()
	if err != nil {
		return nil, err
	}
	return o, nil
}

// Validate valida os dados do Ouvinte
func (o *Ouvinte) Validate() error {
	vld := validator.New()
	if err := vld.Struct(o); err != nil {
		return ErrInvalidEntity
	}
	return nil
}

// AddTelefone adiciona um Telefone
func (o *Ouvinte) AddTelefone(telefone string) error {
	_, err := o.GetTelefone(telefone)
	if err == nil {
		return ErrPhoneRegistered
	}
	o.Telefones = append(o.Telefones, telefone)
	return nil
}

// RemoveTelefone remove um Telefone
func (o *Ouvinte) RemoveTelefone(telefone string) error {
	for i, j := range o.Telefones {
		if j == telefone {
			o.Telefones = append(o.Telefones[:i], o.Telefones[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// GetTelefone get a Telefone
func (o *Ouvinte) GetTelefone(telefone string) (string, error) {
	for _, v := range o.Telefones {
		if v == telefone {
			return telefone, nil
		}
	}
	return telefone, ErrNotFound
}
