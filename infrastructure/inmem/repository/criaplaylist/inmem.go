package criaplaylist

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/criaplaylist"
)

// Inmem in memory repo
type Inmem struct {
	m map[string]*der.CriaPlaylist
}

// NewInmem create new repository
func NewInmem() *Inmem {
	var m = map[string]*der.CriaPlaylist{}
	return &Inmem{
		m: m,
	}
}

// Create CriaPlaylist
func (r *Inmem) Create(e *der.CriaPlaylist) (string, string, error) {
	r.m[e.Playlist+e.Usuario] = e
	return e.Playlist, e.Usuario, nil
}

// Get CriaPlaylist
func (r *Inmem) Get(playlist, usuario string) (*der.CriaPlaylist, error) {
	if r.m[playlist+usuario] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[playlist+usuario], nil
}

// GetByPlaylist CriaPlaylist
func (r *Inmem) GetByPlaylist(playlist string) (*der.CriaPlaylist, error) {
	if r.m[playlist] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[playlist], nil
}

// GetByUsuario CriaPlaylist
func (r *Inmem) GetByUsuario(usuario string) (*der.CriaPlaylist, error) {
	if r.m[usuario] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[usuario], nil
}

// Update CriaPlaylist
func (r *Inmem) Update(e *der.CriaPlaylist) error {
	_, err := r.Get(e.Playlist, e.Usuario)
	if err != nil {
		return err
	}
	r.m[e.Playlist+e.Usuario] = e
	return nil
}

// Search CriaPlaylists
func (r *Inmem) Search(query string) ([]*der.CriaPlaylist, error) {
	var d []*der.CriaPlaylist
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Playlist), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, entity.ErrNotFound
	}

	return d, nil
}

// List CriaPlaylists
func (r *Inmem) List() ([]*der.CriaPlaylist, error) {
	var d []*der.CriaPlaylist
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete CriaPlaylist
func (r *Inmem) Delete(playlist, usuario string) error {
	if r.m[playlist+usuario] == nil {
		return entity.ErrNotFound
	}
	r.m[playlist+usuario] = nil
	return nil
}
