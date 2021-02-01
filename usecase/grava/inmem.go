package grava

import (
	"strconv"
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/grava"
)

// inmem in memory repo
type inmem struct {
	m map[string]*der.Grava
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[string]*der.Grava{}
	return &inmem{
		m: m,
	}
}

// Create Grava
func (r *inmem) Create(e *der.Grava) (int, string, error) {
	r.m[strconv.Itoa(e.Musica)+e.Artista] = e
	return e.Musica, e.Artista, nil
}

// Get Grava
func (r *inmem) Get(musica int, artista string) (*der.Grava, error) {
	if r.m[strconv.Itoa(musica)+artista] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(musica)+artista], nil
}

// GetByMusica Grava
func (r *inmem) GetByMusica(musica int) (*der.Grava, error) {
	if r.m[strconv.Itoa(musica)] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(musica)], nil
}

// GetByArtista Grava
func (r *inmem) GetByArtista(artista string) (*der.Grava, error) {
	if r.m[artista] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[artista], nil
}

// Update Grava
func (r *inmem) Update(e *der.Grava) error {
	_, err := r.Get(e.Musica, e.Artista)
	if err != nil {
		return err
	}
	r.m[strconv.Itoa(e.Musica)+e.Artista] = e
	return nil
}

// Search Gravas
func (r *inmem) Search(query string) ([]*der.Grava, error) {
	var d []*der.Grava
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Artista), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, entity.ErrNotFound
	}

	return d, nil
}

// List Gravas
func (r *inmem) List() ([]*der.Grava, error) {
	var d []*der.Grava
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete Grava
func (r *inmem) Delete(musica int, Artista string) error {
	if r.m[strconv.Itoa(musica)+Artista] == nil {
		return entity.ErrNotFound
	}
	r.m[strconv.Itoa(musica)+Artista] = nil
	return nil
}
