package salvaalbum

import (
	"strconv"
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/salvaalbum"
)

// inmem in memory repo
type inmem struct {
	m map[string]*der.SalvaAlbum
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[string]*der.SalvaAlbum{}
	return &inmem{
		m: m,
	}
}

// Create SalvaAlbum
func (r *inmem) Create(e *der.SalvaAlbum) (int, string, string, error) {
	r.m[strconv.Itoa(e.Album)+e.Ouvinte] = e
	return e.Album, e.Ouvinte, e.Artista, nil
}

// Get SalvaAlbum
func (r *inmem) Get(album int, ouvinte, artista string) (*der.SalvaAlbum, error) {
	if r.m[strconv.Itoa(album)+ouvinte] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(album)+ouvinte], nil
}

// GetByAlbum SalvaAlbum
func (r *inmem) GetByAlbum(album int) (*der.SalvaAlbum, error) {
	if r.m[strconv.Itoa(album)] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(album)], nil
}

// GetByOuvinte SalvaAlbum
func (r *inmem) GetByOuvinte(ouvinte string) (*der.SalvaAlbum, error) {
	if r.m[ouvinte] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[ouvinte], nil
}

// Update SalvaAlbum
func (r *inmem) Update(e *der.SalvaAlbum) error {
	_, err := r.Get(e.Album, e.Ouvinte, e.Artista)
	if err != nil {
		return err
	}
	r.m[strconv.Itoa(e.Album)+e.Ouvinte] = e
	return nil
}

// Search SalvaAlbums
func (r *inmem) Search(query string) ([]*der.SalvaAlbum, error) {
	var d []*der.SalvaAlbum
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

// List SalvaAlbums
func (r *inmem) List() ([]*der.SalvaAlbum, error) {
	var d []*der.SalvaAlbum
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete SalvaAlbum
func (r *inmem) Delete(album int, ouvinte, artista string) error {
	if r.m[strconv.Itoa(album)+ouvinte] == nil {
		return entity.ErrNotFound
	}
	r.m[strconv.Itoa(album)+ouvinte] = nil
	return nil
}
