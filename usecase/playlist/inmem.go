package playlist

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/playlist"
)

// inmem in memory repo
type inmem struct {
	m map[string]*der.Playlist
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[string]*der.Playlist{}
	return &inmem{
		m: m,
	}
}

// Create Playlist
func (r *inmem) Create(e *der.Playlist) (string, error) {
	r.m[e.Nome] = e
	return e.Nome, nil
}

// Get Playlist
func (r *inmem) Get(email string) (*der.Playlist, error) {
	if r.m[email] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[email], nil
}

// Update Playlist
func (r *inmem) Update(e *der.Playlist) error {
	_, err := r.Get(e.Nome)
	if err != nil {
		return err
	}
	r.m[e.Nome] = e
	return nil
}

// Search Playlists
func (r *inmem) Search(query string) ([]*der.Playlist, error) {
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
func (r *inmem) List() ([]*der.Playlist, error) {
	var d []*der.Playlist
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete Playlist
func (r *inmem) Delete(email string) error {
	if r.m[email] == nil {
		return entity.ErrNotFound
	}
	r.m[email] = nil
	return nil
}
