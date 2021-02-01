package salvaplaylist

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/salvaplaylist"
)

// Inmem in memory repo
type Inmem struct {
	m map[string]*der.SalvaPlaylist
}

// NewInmem create new repository
func NewInmem() *Inmem {
	var m = map[string]*der.SalvaPlaylist{}
	return &Inmem{
		m: m,
	}
}

// Create SalvaPlaylist
func (r *Inmem) Create(e *der.SalvaPlaylist) (string, string, error) {
	r.m[e.Playlist+e.Ouvinte] = e
	return e.Playlist, e.Ouvinte, nil
}

// Get SalvaPlaylist
func (r *Inmem) Get(playlist, ouvinte string) (*der.SalvaPlaylist, error) {
	if r.m[playlist+ouvinte] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[playlist+ouvinte], nil
}

// GetByPlaylist SalvaPlaylist
func (r *Inmem) GetByPlaylist(playlist string) (*der.SalvaPlaylist, error) {
	if r.m[playlist] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[playlist], nil
}

// GetByOuvinte SalvaPlaylist
func (r *Inmem) GetByOuvinte(ouvinte string) (*der.SalvaPlaylist, error) {
	if r.m[ouvinte] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[ouvinte], nil
}

// Update SalvaPlaylist
func (r *Inmem) Update(e *der.SalvaPlaylist) error {
	_, err := r.Get(e.Playlist, e.Ouvinte)
	if err != nil {
		return err
	}
	r.m[e.Playlist+e.Ouvinte] = e
	return nil
}

// Search SalvaPlaylists
func (r *Inmem) Search(query string) ([]*der.SalvaPlaylist, error) {
	var d []*der.SalvaPlaylist
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

// List SalvaPlaylists
func (r *Inmem) List() ([]*der.SalvaPlaylist, error) {
	var d []*der.SalvaPlaylist
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete SalvaPlaylist
func (r *Inmem) Delete(playlist, ouvinte string) error {
	if r.m[playlist+ouvinte] == nil {
		return entity.ErrNotFound
	}
	r.m[playlist+ouvinte] = nil
	return nil
}
