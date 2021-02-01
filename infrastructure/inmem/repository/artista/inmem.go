package artista

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/artista"
)

// Inmem in memory repo
type Inmem struct {
	m map[string]*der.Artista
}

// NewInmem create new repository
func NewInmem() *Inmem {
	var m = map[string]*der.Artista{}
	return &Inmem{
		m: m,
	}
}

// Create Artista
func (r *Inmem) Create(e *der.Artista) (string, error) {
	r.m[e.Usuario] = e
	return e.Usuario, nil
}

// Get Artista
func (r *Inmem) Get(email string) (*der.Artista, error) {
	if r.m[email] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[email], nil
}

// Update Artista
func (r *Inmem) Update(e *der.Artista) error {
	_, err := r.Get(e.Usuario)
	if err != nil {
		return err
	}
	r.m[e.Usuario] = e
	return nil
}

// Search Artistas
func (r *Inmem) Search(query string) ([]*der.Artista, error) {
	var d []*der.Artista
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

// List Artistas
func (r *Inmem) List() ([]*der.Artista, error) {
	var d []*der.Artista
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete Artista
func (r *Inmem) Delete(email string) error {
	if r.m[email] == nil {
		return entity.ErrNotFound
	}
	r.m[email] = nil
	return nil
}
