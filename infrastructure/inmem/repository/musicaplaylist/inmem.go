package musicaplaylist

import (
	"strconv"
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/musicaplaylist"
)

// Inmem in memory repo
type Inmem struct {
	m map[string]*der.MusicaPlaylist
}

// NewInmem create new repository
func NewInmem() *Inmem {
	var m = map[string]*der.MusicaPlaylist{}
	return &Inmem{
		m: m,
	}
}

// Create MusicaPlaylist
func (r *Inmem) Create(e *der.MusicaPlaylist) (int, string, error) {
	r.m[strconv.Itoa(e.Musica)+e.Playlist] = e
	return e.Musica, e.Playlist, nil
}

// Get MusicaPlaylist
func (r *Inmem) Get(musica int, playlist string) (*der.MusicaPlaylist, error) {
	if r.m[strconv.Itoa(musica)+playlist] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(musica)+playlist], nil
}

// GetByMusica MusicaPlaylist
func (r *Inmem) GetByMusica(musica int) (*der.MusicaPlaylist, error) {
	if r.m[strconv.Itoa(musica)] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(musica)], nil
}

// GetByPlaylist MusicaPlaylist
func (r *Inmem) GetByPlaylist(playlist string) (*der.MusicaPlaylist, error) {
	if r.m[playlist] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[playlist], nil
}

// Update MusicaPlaylist
func (r *Inmem) Update(e *der.MusicaPlaylist) error {
	_, err := r.Get(e.Musica, e.Playlist)
	if err != nil {
		return err
	}
	r.m[strconv.Itoa(e.Musica)+e.Playlist] = e
	return nil
}

// Search MusicaPlaylists
func (r *Inmem) Search(query string) ([]*der.MusicaPlaylist, error) {
	var d []*der.MusicaPlaylist
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

// List MusicaPlaylists
func (r *Inmem) List() ([]*der.MusicaPlaylist, error) {
	var d []*der.MusicaPlaylist
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete MusicaPlaylist
func (r *Inmem) Delete(musica int, playlist string) error {
	if r.m[strconv.Itoa(musica)+playlist] == nil {
		return entity.ErrNotFound
	}
	r.m[strconv.Itoa(musica)+playlist] = nil
	return nil
}
