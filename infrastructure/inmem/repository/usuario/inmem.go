package usuario

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/usuario"
)

// Inmem in memory repo
type Inmem struct {
	m map[string]*der.Usuario
}

// NewInmem create new repository
func NewInmem() *Inmem {
	var m = map[string]*der.Usuario{}
	return &Inmem{
		m: m,
	}
}

// Create Usuario
func (r *Inmem) Create(e *der.Usuario) (string, error) {
	r.m[e.Email] = e
	return e.Email, nil
}

// Get Usuario
func (r *Inmem) Get(email string) (*der.Usuario, error) {
	if r.m[email] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[email], nil
}

// Update Usuario
func (r *Inmem) Update(e *der.Usuario) error {
	_, err := r.Get(e.Email)
	if err != nil {
		return err
	}
	r.m[e.Email] = e
	return nil
}

// Search Usuarios
func (r *Inmem) Search(query string) ([]*der.Usuario, error) {
	var d []*der.Usuario
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Email), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, entity.ErrNotFound
	}

	return d, nil
}

// List Usuarios
func (r *Inmem) List() ([]*der.Usuario, error) {
	var d []*der.Usuario
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete Usuario
func (r *Inmem) Delete(email string) error {
	if r.m[email] == nil {
		return entity.ErrNotFound
	}
	r.m[email] = nil
	return nil
}
