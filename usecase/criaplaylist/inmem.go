package criaplaylist

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/criaplaylist"
)

// inmem in memory repo
type inmem struct {
	m map[string]*der.CriaPlaylist
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[string]*der.CriaPlaylist{}
	return &inmem{
		m: m,
	}
}

// Create CriaPlaylist
func (r *inmem) Create(e *der.CriaPlaylist) (string, string, error) {
	r.m[e.Playlist+e.Usuario] = e
	return e.Playlist, e.Usuario, nil
}

// Get CriaPlaylist
func (r *inmem) Get(playlist, usuario string) (*der.CriaPlaylist, error) {
	if r.m[playlist+usuario] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[playlist+usuario], nil
}

// GetByPlaylist CriaPlaylist
func (r *inmem) GetByPlaylist(playlist string) (*der.CriaPlaylist, error) {
	if r.m[playlist] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[playlist], nil
}

// GetByUsuario CriaPlaylist
func (r *inmem) GetByUsuario(usuario string) (*der.CriaPlaylist, error) {
	if r.m[usuario] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[usuario], nil
}

// Update CriaPlaylist
func (r *inmem) Update(e *der.CriaPlaylist) error {
	_, err := r.Get(e.Playlist, e.Usuario)
	if err != nil {
		return err
	}
	r.m[e.Playlist+e.Usuario] = e
	return nil
}

// Search CriaPlaylists
func (r *inmem) Search(query string) ([]*der.CriaPlaylist, error) {
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
func (r *inmem) List() ([]*der.CriaPlaylist, error) {
	var d []*der.CriaPlaylist
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete CriaPlaylist
func (r *inmem) Delete(playlist, usuario string) error {
	if r.m[playlist+usuario] == nil {
		return entity.ErrNotFound
	}
	r.m[playlist+usuario] = nil
	return nil
}
