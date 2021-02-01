package playlist

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/playlist"
)

// Inmem in memory repo
type Inmem struct {
	m map[string]*der.Playlist
}

// NewInmem create new repository
func NewInmem() *Inmem {
	var m = map[string]*der.Playlist{}
	return &Inmem{
		m: m,
	}
}

// Create Playlist
func (r *Inmem) Create(e *der.Playlist) (string, error) {
	r.m[e.Nome] = e
	return e.Nome, nil
}

// Get Playlist
func (r *Inmem) Get(email string) (*der.Playlist, error) {
	if r.m[email] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[email], nil
}

// Update Playlist
func (r *Inmem) Update(e *der.Playlist) error {
	_, err := r.Get(e.Nome)
	if err != nil {
		return err
	}
	r.m[e.Nome] = e
	return nil
}

// Search Playlists
func (r *Inmem) Search(query string) ([]*der.Playlist, error) {
	var d []*der.Playlist
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Nome), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, entity.ErrNotFound
	}

	return d, nil
}

// List Playlists
func (r *Inmem) List() ([]*der.Playlist, error) {
	var d []*der.Playlist
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete Playlist
func (r *Inmem) Delete(email string) error {
	if r.m[email] == nil {
		return entity.ErrNotFound
	}
	r.m[email] = nil
	return nil
}
