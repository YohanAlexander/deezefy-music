package artista

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/artista"
)

// inmem in memory repo
type inmem struct {
	m map[string]*der.Artista
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[string]*der.Artista{}
	return &inmem{
		m: m,
	}
}

// Create Artista
func (r *inmem) Create(e *der.Artista) (string, error) {
	r.m[e.Usuario] = e
	return e.Usuario, nil
}

// Get Artista
func (r *inmem) Get(email string) (*der.Artista, error) {
	if r.m[email] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[email], nil
}

// Update Artista
func (r *inmem) Update(e *der.Artista) error {
	_, err := r.Get(e.Usuario)
	if err != nil {
		return err
	}
	r.m[e.Usuario] = e
	return nil
}

// Search Artistas
func (r *inmem) Search(query string) ([]*der.Artista, error) {
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
func (r *inmem) List() ([]*der.Artista, error) {
	var d []*der.Artista
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete Artista
func (r *inmem) Delete(email string) error {
	if r.m[email] == nil {
		return entity.ErrNotFound
	}
	r.m[email] = nil
	return nil
}
