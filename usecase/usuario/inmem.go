package usuario

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/usuario"
)

// inmem in memory repo
type inmem struct {
	m map[string]*der.Usuario
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[string]*der.Usuario{}
	return &inmem{
		m: m,
	}
}

// Create Usuario
func (r *inmem) Create(e *der.Usuario) (string, error) {
	r.m[e.Email] = e
	return e.Email, nil
}

// Get Usuario
func (r *inmem) Get(email string) (*der.Usuario, error) {
	if r.m[email] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[email], nil
}

// Update Usuario
func (r *inmem) Update(e *der.Usuario) error {
	_, err := r.Get(e.Email)
	if err != nil {
		return err
	}
	r.m[e.Email] = e
	return nil
}

// Search Usuarios
func (r *inmem) Search(query string) ([]*der.Usuario, error) {
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
func (r *inmem) List() ([]*der.Usuario, error) {
	var d []*der.Usuario
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete Usuario
func (r *inmem) Delete(email string) error {
	if r.m[email] == nil {
		return entity.ErrNotFound
	}
	r.m[email] = nil
	return nil
}
