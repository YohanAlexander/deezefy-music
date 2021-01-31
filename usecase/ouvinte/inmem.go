package ouvinte

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/ouvinte"
)

// inmem in memory repo
type inmem struct {
	m map[string]*der.Ouvinte
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[string]*der.Ouvinte{}
	return &inmem{
		m: m,
	}
}

// Create Ouvinte
func (r *inmem) Create(e *der.Ouvinte) (string, error) {
	r.m[e.Usuario] = e
	return e.Usuario, nil
}

// Get Ouvinte
func (r *inmem) Get(email string) (*der.Ouvinte, error) {
	if r.m[email] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[email], nil
}

// Update Ouvinte
func (r *inmem) Update(e *der.Ouvinte) error {
	_, err := r.Get(e.Usuario)
	if err != nil {
		return err
	}
	r.m[e.Usuario] = e
	return nil
}

// Search Ouvintes
func (r *inmem) Search(query string) ([]*der.Ouvinte, error) {
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
func (r *inmem) List() ([]*der.Ouvinte, error) {
	var d []*der.Ouvinte
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
