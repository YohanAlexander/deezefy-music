package artistasfavoritos

import (
	"strconv"
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/artistasfavoritos"
)

// inmem in memory repo
type inmem struct {
	m map[string]*der.ArtistasFavoritos
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[string]*der.ArtistasFavoritos{}
	return &inmem{
		m: m,
	}
}

// Create ArtistasFavoritos
func (r *inmem) Create(e *der.ArtistasFavoritos) (int, string, string, error) {
	r.m[strconv.Itoa(e.Perfil)+e.Ouvinte] = e
	return e.Perfil, e.Ouvinte, e.Artista, nil
}

// Get ArtistasFavoritos
func (r *inmem) Get(perfil int, ouvinte, artista string) (*der.ArtistasFavoritos, error) {
	if r.m[strconv.Itoa(perfil)+ouvinte] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(perfil)+ouvinte], nil
}

// GetByPerfil ArtistasFavoritos
func (r *inmem) GetByPerfil(perfil int) (*der.ArtistasFavoritos, error) {
	if r.m[strconv.Itoa(perfil)] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(perfil)], nil
}

// GetByOuvinte ArtistasFavoritos
func (r *inmem) GetByOuvinte(ouvinte string) (*der.ArtistasFavoritos, error) {
	if r.m[ouvinte] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[ouvinte], nil
}

// Update ArtistasFavoritos
func (r *inmem) Update(e *der.ArtistasFavoritos) error {
	_, err := r.Get(e.Perfil, e.Ouvinte, e.Artista)
	if err != nil {
		return err
	}
	r.m[strconv.Itoa(e.Perfil)+e.Ouvinte] = e
	return nil
}

// Search ArtistasFavoritoss
func (r *inmem) Search(query string) ([]*der.ArtistasFavoritos, error) {
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
func (r *inmem) List() ([]*der.ArtistasFavoritos, error) {
	var d []*der.ArtistasFavoritos
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete ArtistasFavoritos
func (r *inmem) Delete(perfil int, ouvinte, artista string) error {
	if r.m[strconv.Itoa(perfil)+ouvinte] == nil {
		return entity.ErrNotFound
	}
	r.m[strconv.Itoa(perfil)+ouvinte] = nil
	return nil
}
