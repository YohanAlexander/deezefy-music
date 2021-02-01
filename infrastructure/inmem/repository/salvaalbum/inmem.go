package salvaalbum

import (
	"strconv"
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/salvaalbum"
)

// Inmem in memory repo
type Inmem struct {
	m map[string]*der.SalvaAlbum
}

// NewInmem create new repository
func NewInmem() *Inmem {
	var m = map[string]*der.SalvaAlbum{}
	return &Inmem{
		m: m,
	}
}

// Create SalvaAlbum
func (r *Inmem) Create(e *der.SalvaAlbum) (int, string, string, error) {
	r.m[strconv.Itoa(e.Album)+e.Ouvinte] = e
	return e.Album, e.Ouvinte, e.Artista, nil
}

// Get SalvaAlbum
func (r *Inmem) Get(album int, ouvinte, artista string) (*der.SalvaAlbum, error) {
	if r.m[strconv.Itoa(album)+ouvinte] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(album)+ouvinte], nil
}

// GetByAlbum SalvaAlbum
func (r *Inmem) GetByAlbum(album int) (*der.SalvaAlbum, error) {
	if r.m[strconv.Itoa(album)] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(album)], nil
}

// GetByOuvinte SalvaAlbum
func (r *Inmem) GetByOuvinte(ouvinte string) (*der.SalvaAlbum, error) {
	if r.m[ouvinte] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[ouvinte], nil
}

// Update SalvaAlbum
func (r *Inmem) Update(e *der.SalvaAlbum) error {
	_, err := r.Get(e.Album, e.Ouvinte, e.Artista)
	if err != nil {
		return err
	}
	r.m[strconv.Itoa(e.Album)+e.Ouvinte] = e
	return nil
}

// Search SalvaAlbums
func (r *Inmem) Search(query string) ([]*der.SalvaAlbum, error) {
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
func (r *Inmem) List() ([]*der.SalvaAlbum, error) {
	var d []*der.SalvaAlbum
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete SalvaAlbum
func (r *Inmem) Delete(album int, ouvinte, artista string) error {
	if r.m[strconv.Itoa(album)+ouvinte] == nil {
		return entity.ErrNotFound
	}
	r.m[strconv.Itoa(album)+ouvinte] = nil
	return nil
}
