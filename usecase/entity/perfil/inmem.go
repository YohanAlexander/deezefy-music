package perfil

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
)

// inmem in memory repo
type inmem struct {
	m map[string]*entity.Perfil
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[string]*entity.Perfil{}
	return &inmem{
		m: m,
	}
}

// Create Perfil
func (r *inmem) Create(e *entity.Perfil) (string, error) {
	r.m[e.Ouvinte.Usuario.Email] = e
	return e.Ouvinte.Usuario.Email, nil
}

// Get Perfil
func (r *inmem) Get(email string) (*entity.Perfil, error) {
	if r.m[email] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[email], nil
}

// Update Perfil
func (r *inmem) Update(e *entity.Perfil) error {
	_, err := r.Get(e.Ouvinte.Usuario.Email)
	if err != nil {
		return err
	}
	r.m[e.Ouvinte.Usuario.Email] = e
	return nil
}

// Search Perfils
func (r *inmem) Search(query string) ([]*entity.Perfil, error) {
	var d []*entity.Perfil
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Ouvinte.Usuario.Email), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, entity.ErrNotFound
	}

	return d, nil
}

// List Perfils
func (r *inmem) List() ([]*entity.Perfil, error) {
	var d []*entity.Perfil
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete Perfil
func (r *inmem) Delete(email string) error {
	if r.m[email] == nil {
		return entity.ErrNotFound
	}
	r.m[email] = nil
	return nil
}
