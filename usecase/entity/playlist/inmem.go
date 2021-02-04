package playlist

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
)

// inmem in memory repo
type inmem struct {
	m map[string]*entity.Playlist
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[string]*entity.Playlist{}
	return &inmem{
		m: m,
	}
}

// Create Playlist
func (r *inmem) Create(e *entity.Playlist) (string, error) {
	r.m[e.Nome] = e
	return e.Nome, nil
}

// Get Playlist
func (r *inmem) Get(nome string) (*entity.Playlist, error) {
	if r.m[nome] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[nome], nil
}

// Update Playlist
func (r *inmem) Update(e *entity.Playlist) error {
	_, err := r.Get(e.Nome)
	if err != nil {
		return err
	}
	r.m[e.Nome] = e
	return nil
}

// Search Playlists
func (r *inmem) Search(query string) ([]*entity.Playlist, error) {
	var d []*entity.Playlist
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
func (r *inmem) List() ([]*entity.Playlist, error) {
	var d []*entity.Playlist
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete Playlist
func (r *inmem) Delete(nome string) error {
	if r.m[nome] == nil {
		return entity.ErrNotFound
	}
	r.m[nome] = nil
	return nil
}
