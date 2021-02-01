package genero

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/genero"
)

// Inmem in memory repo
type Inmem struct {
	m map[string]*der.Genero
}

// NewInmem create new repository
func NewInmem() *Inmem {
	var m = map[string]*der.Genero{}
	return &Inmem{
		m: m,
	}
}

// Create Genero
func (r *Inmem) Create(e *der.Genero) (string, error) {
	r.m[e.Nome] = e
	return e.Nome, nil
}

// Get Genero
func (r *Inmem) Get(email string) (*der.Genero, error) {
	if r.m[email] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[email], nil
}

// Update Genero
func (r *Inmem) Update(e *der.Genero) error {
	_, err := r.Get(e.Nome)
	if err != nil {
		return err
	}
	r.m[e.Nome] = e
	return nil
}

// Search Generos
func (r *Inmem) Search(query string) ([]*der.Genero, error) {
	var d []*der.Genero
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Nome), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, entity.ErrNotFound
	}

	return d, nil
}

// List Generos
func (r *Inmem) List() ([]*der.Genero, error) {
	var d []*der.Genero
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete Genero
func (r *Inmem) Delete(email string) error {
	if r.m[email] == nil {
		return entity.ErrNotFound
	}
	r.m[email] = nil
	return nil
}
