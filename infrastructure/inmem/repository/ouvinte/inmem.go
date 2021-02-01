package ouvinte

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/ouvinte"
)

// Inmem in memory repo
type Inmem struct {
	m map[string]*der.Ouvinte
}

// NewInmem create new repository
func NewInmem() *Inmem {
	var m = map[string]*der.Ouvinte{}
	return &Inmem{
		m: m,
	}
}

// Create Ouvinte
func (r *Inmem) Create(e *der.Ouvinte) (string, error) {
	r.m[e.Usuario] = e
	return e.Usuario, nil
}

// Get Ouvinte
func (r *Inmem) Get(email string) (*der.Ouvinte, error) {
	if r.m[email] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[email], nil
}

// Update Ouvinte
func (r *Inmem) Update(e *der.Ouvinte) error {
	_, err := r.Get(e.Usuario)
	if err != nil {
		return err
	}
	r.m[e.Usuario] = e
	return nil
}

// Search Ouvintes
func (r *Inmem) Search(query string) ([]*der.Ouvinte, error) {
	var d []*der.Ouvinte
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
func (r *Inmem) List() ([]*der.Ouvinte, error) {
	var d []*der.Ouvinte
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete Ouvinte
func (r *Inmem) Delete(email string) error {
	if r.m[email] == nil {
		return entity.ErrNotFound
	}
	r.m[email] = nil
	return nil
}
