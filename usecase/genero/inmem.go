package genero

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
)

// inmem in memory repo
type inmem struct {
	m map[string]*entity.Genero
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[string]*entity.Genero{}
	return &inmem{
		m: m,
	}
}

// Create Genero
func (r *inmem) Create(e *entity.Genero) (string, error) {
	r.m[e.Nome] = e
	return e.Nome, nil
}

// Get Genero
func (r *inmem) Get(email string) (*entity.Genero, error) {
	if r.m[email] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[email], nil
}

// Update Genero
func (r *inmem) Update(e *entity.Genero) error {
	_, err := r.Get(e.Nome)
	if err != nil {
		return err
	}
	r.m[e.Nome] = e
	return nil
}

// Search Generos
func (r *inmem) Search(query string) ([]*entity.Genero, error) {
	var d []*entity.Genero
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
func (r *inmem) List() ([]*entity.Genero, error) {
	var d []*entity.Genero
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete Genero
func (r *inmem) Delete(email string) error {
	if r.m[email] == nil {
		return entity.ErrNotFound
	}
	r.m[email] = nil
	return nil
}
