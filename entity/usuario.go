package entity

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

// Usuario entidade usuario
type Usuario struct {
	Email    string     `validate:"required,email"`
	Password string     `validate:"required,gte=8"`
	Birthday string     `validate:"datetime=2006-01-02"`
	Cria     []Playlist `validate:""`
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

// AddPlaylist adiciona um Playlist
func (u *Usuario) AddPlaylist(playlist Playlist) error {
	_, err := u.GetPlaylist(playlist)
	if err == nil {
		return ErrPlaylistRegistered
	}
	u.Cria = append(u.Cria, playlist)
	return nil
}

// RemovePlaylist remove um Playlist
func (u *Usuario) RemovePlaylist(playlist Playlist) error {
	for i, j := range u.Cria {
		if j.Nome == playlist.Nome {
			u.Cria = append(u.Cria[:i], u.Cria[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// GetPlaylist get a Playlist
func (u *Usuario) GetPlaylist(playlist Playlist) (Playlist, error) {
	for _, v := range u.Cria {
		if v.Nome == playlist.Nome {
			return playlist, nil
		}
	}
	return playlist, ErrNotFound
}
