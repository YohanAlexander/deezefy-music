package ouvinte

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
)

// inmem in memory repo
type inmem struct {
	m map[string]*entity.Ouvinte
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[string]*entity.Ouvinte{}
	return &inmem{
		m: m,
	}
}

// Create Ouvinte
func (r *inmem) Create(e *entity.Ouvinte) (string, error) {
	r.m[e.Usuario] = e
	return e.Usuario, nil
}

// Get Ouvinte
func (r *inmem) Get(email string) (*entity.Ouvinte, error) {
	if r.m[email] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[email], nil
}

// Update Ouvinte
func (r *inmem) Update(e *entity.Ouvinte) error {
	_, err := r.Get(e.Usuario)
	if err != nil {
		return err
	}
	r.m[e.Usuario] = e
	return nil
}

// Search Ouvintes
func (r *inmem) Search(query string) ([]*entity.Ouvinte, error) {
	var d []*entity.Ouvinte
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Usuario), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, entity.ErrNotFound
	}

	return d, nil
}

// List Ouvintes
func (r *inmem) List() ([]*entity.Ouvinte, error) {
	var d []*entity.Ouvinte
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete Ouvinte
func (r *inmem) Delete(email string) error {
	if r.m[email] == nil {
		return entity.ErrNotFound
	}
	r.m[email] = nil
	return nil
}
