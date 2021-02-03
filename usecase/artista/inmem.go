package artista

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
)

// inmem in memory repo
type inmem struct {
	m map[string]*entity.Artista
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[string]*entity.Artista{}
	return &inmem{
		m: m,
	}
}

// Create Artista
func (r *inmem) Create(e *entity.Artista) (string, error) {
	r.m[e.Usuario] = e
	return e.Usuario, nil
}

// Get Artista
func (r *inmem) Get(email string) (*entity.Artista, error) {
	if r.m[email] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[email], nil
}

// Update Artista
func (r *inmem) Update(e *entity.Artista) error {
	_, err := r.Get(e.Usuario)
	if err != nil {
		return err
	}
	r.m[e.Usuario] = e
	return nil
}

// Search Artistas
func (r *inmem) Search(query string) ([]*entity.Artista, error) {
	var d []*entity.Artista
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
func (r *inmem) List() ([]*entity.Artista, error) {
	var d []*entity.Artista
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
