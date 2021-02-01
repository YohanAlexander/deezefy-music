package artistagenero

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/artistagenero"
)

// Inmem in memory repo
type Inmem struct {
	m map[string]*der.ArtistaGenero
}

// NewInmem create new repository
func NewInmem() *Inmem {
	var m = map[string]*der.ArtistaGenero{}
	return &Inmem{
		m: m,
	}
}

// Create ArtistaGenero
func (r *Inmem) Create(e *der.ArtistaGenero) (string, string, error) {
	r.m[e.Artista+e.Genero] = e
	return e.Artista, e.Genero, nil
}

// Get ArtistaGenero
func (r *Inmem) Get(artista, genero string) (*der.ArtistaGenero, error) {
	if r.m[artista+genero] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[artista+genero], nil
}

// GetByArtista ArtistaGenero
func (r *Inmem) GetByArtista(artista string) (*der.ArtistaGenero, error) {
	if r.m[artista] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[artista], nil
}

// GetByGenero ArtistaGenero
func (r *Inmem) GetByGenero(genero string) (*der.ArtistaGenero, error) {
	if r.m[genero] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[genero], nil
}

// Update ArtistaGenero
func (r *Inmem) Update(e *der.ArtistaGenero) error {
	_, err := r.Get(e.Artista, e.Genero)
	if err != nil {
		return err
	}
	r.m[e.Artista+e.Genero] = e
	return nil
}

// Search ArtistaGeneros
func (r *Inmem) Search(query string) ([]*der.ArtistaGenero, error) {
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
func (r *Inmem) List() ([]*der.ArtistaGenero, error) {
	var d []*der.ArtistaGenero
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete ArtistaGenero
func (r *Inmem) Delete(artista, genero string) error {
	if r.m[artista+genero] == nil {
		return entity.ErrNotFound
	}
	r.m[artista+genero] = nil
	return nil
}
