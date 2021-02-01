package artistasfavoritos

import (
	"strconv"
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/artistasfavoritos"
)

// Inmem in memory repo
type Inmem struct {
	m map[string]*der.ArtistasFavoritos
}

// NewInmem create new repository
func NewInmem() *Inmem {
	var m = map[string]*der.ArtistasFavoritos{}
	return &Inmem{
		m: m,
	}
}

// Create ArtistasFavoritos
func (r *Inmem) Create(e *der.ArtistasFavoritos) (int, string, string, error) {
	r.m[strconv.Itoa(e.Perfil)+e.Ouvinte] = e
	return e.Perfil, e.Ouvinte, e.Artista, nil
}

// Get ArtistasFavoritos
func (r *Inmem) Get(perfil int, ouvinte, artista string) (*der.ArtistasFavoritos, error) {
	if r.m[strconv.Itoa(perfil)+ouvinte] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(perfil)+ouvinte], nil
}

// GetByPerfil ArtistasFavoritos
func (r *Inmem) GetByPerfil(perfil int) (*der.ArtistasFavoritos, error) {
	if r.m[strconv.Itoa(perfil)] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(perfil)], nil
}

// GetByOuvinte ArtistasFavoritos
func (r *Inmem) GetByOuvinte(ouvinte string) (*der.ArtistasFavoritos, error) {
	if r.m[ouvinte] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[ouvinte], nil
}

// Update ArtistasFavoritos
func (r *Inmem) Update(e *der.ArtistasFavoritos) error {
	_, err := r.Get(e.Perfil, e.Ouvinte, e.Artista)
	if err != nil {
		return err
	}
	r.m[strconv.Itoa(e.Perfil)+e.Ouvinte] = e
	return nil
}

// Search ArtistasFavoritoss
func (r *Inmem) Search(query string) ([]*der.ArtistasFavoritos, error) {
	var d []*der.ArtistasFavoritos
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Ouvinte), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, entity.ErrNotFound
	}

	return d, nil
}

// List ArtistasFavoritoss
func (r *Inmem) List() ([]*der.ArtistasFavoritos, error) {
	var d []*der.ArtistasFavoritos
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete ArtistasFavoritos
func (r *Inmem) Delete(perfil int, ouvinte, artista string) error {
	if r.m[strconv.Itoa(perfil)+ouvinte] == nil {
		return entity.ErrNotFound
	}
	r.m[strconv.Itoa(perfil)+ouvinte] = nil
	return nil
}
