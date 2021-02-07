package segue

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/segue"
)

// inmem in memory repo
type inmem struct {
	m map[string]*der.Segue
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[string]*der.Segue{}
	return &inmem{
		m: m,
	}
}

// Create Segue
func (r *inmem) Create(e *der.Segue) (string, string, error) {
	r.m[e.Artista+e.Ouvinte] = e
	return e.Artista, e.Ouvinte, nil
}

// Get Segue
func (r *inmem) Get(artista, ouvinte string) (*der.Segue, error) {
	if r.m[artista+ouvinte] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[artista+ouvinte], nil
}

// GetByArtista Segue
func (r *inmem) GetByArtista(artista string) (*der.Segue, error) {
	if r.m[artista] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[artista], nil
}

// GetByOuvinte Segue
func (r *inmem) GetByOuvinte(ouvinte string) (*der.Segue, error) {
	if r.m[ouvinte] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[ouvinte], nil
}

// Update Segue
func (r *inmem) Update(e *der.Segue) error {
	_, err := r.Get(e.Artista, e.Ouvinte)
	if err != nil {
		return err
	}
	r.m[e.Artista+e.Ouvinte] = e
	return nil
}

// Search Segues
func (r *inmem) Search(query string) ([]*der.Segue, error) {
	var d []*der.Segue
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Artista), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, entity.ErrNotFound
	}

	return d, nil
}

// List Segues
func (r *inmem) List() ([]*der.Segue, error) {
	var d []*der.Segue
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete Segue
func (r *inmem) Delete(artista, ouvinte string) error {
	if r.m[artista+ouvinte] == nil {
		return entity.ErrNotFound
	}
	r.m[artista+ouvinte] = nil
	return nil
}
