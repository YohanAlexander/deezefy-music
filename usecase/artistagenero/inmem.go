package artistagenero

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/artistagenero"
)

// inmem in memory repo
type inmem struct {
	m map[string]*der.ArtistaGenero
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[string]*der.ArtistaGenero{}
	return &inmem{
		m: m,
	}
}

// Create ArtistaGenero
func (r *inmem) Create(e *der.ArtistaGenero) (string, string, error) {
	r.m[e.Artista+e.Genero] = e
	return e.Artista, e.Genero, nil
}

// Get ArtistaGenero
func (r *inmem) Get(artista, genero string) (*der.ArtistaGenero, error) {
	if r.m[artista+genero] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[artista+genero], nil
}

// GetByArtista ArtistaGenero
func (r *inmem) GetByArtista(artista string) (*der.ArtistaGenero, error) {
	if r.m[artista] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[artista], nil
}

// GetByGenero ArtistaGenero
func (r *inmem) GetByGenero(genero string) (*der.ArtistaGenero, error) {
	if r.m[genero] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[genero], nil
}

// Update ArtistaGenero
func (r *inmem) Update(e *der.ArtistaGenero) error {
	_, err := r.Get(e.Artista, e.Genero)
	if err != nil {
		return err
	}
	r.m[e.Artista+e.Genero] = e
	return nil
}

// Search ArtistaGeneros
func (r *inmem) Search(query string) ([]*der.ArtistaGenero, error) {
	var d []*der.ArtistaGenero
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

// List ArtistaGeneros
func (r *inmem) List() ([]*der.ArtistaGenero, error) {
	var d []*der.ArtistaGenero
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete ArtistaGenero
func (r *inmem) Delete(artista, genero string) error {
	if r.m[artista+genero] == nil {
		return entity.ErrNotFound
	}
	r.m[artista+genero] = nil
	return nil
}
