package grava

import (
	"strconv"
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/grava"
)

// Inmem in memory repo
type Inmem struct {
	m map[string]*der.Grava
}

// NewInmem create new repository
func NewInmem() *Inmem {
	var m = map[string]*der.Grava{}
	return &Inmem{
		m: m,
	}
}

// Create Grava
func (r *Inmem) Create(e *der.Grava) (int, string, error) {
	r.m[strconv.Itoa(e.Musica)+e.Artista] = e
	return e.Musica, e.Artista, nil
}

// Get Grava
func (r *Inmem) Get(musica int, artista string) (*der.Grava, error) {
	if r.m[strconv.Itoa(musica)+artista] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(musica)+artista], nil
}

// GetByMusica Grava
func (r *Inmem) GetByMusica(musica int) (*der.Grava, error) {
	if r.m[strconv.Itoa(musica)] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(musica)], nil
}

// GetByArtista Grava
func (r *Inmem) GetByArtista(artista string) (*der.Grava, error) {
	if r.m[artista] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[artista], nil
}

// Update Grava
func (r *Inmem) Update(e *der.Grava) error {
	_, err := r.Get(e.Musica, e.Artista)
	if err != nil {
		return err
	}
	r.m[strconv.Itoa(e.Musica)+e.Artista] = e
	return nil
}

// Search Gravas
func (r *Inmem) Search(query string) ([]*der.Grava, error) {
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
func (r *Inmem) List() ([]*der.Grava, error) {
	var d []*der.Grava
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete Grava
func (r *Inmem) Delete(musica int, Artista string) error {
	if r.m[strconv.Itoa(musica)+Artista] == nil {
		return entity.ErrNotFound
	}
	r.m[strconv.Itoa(musica)+Artista] = nil
	return nil
}
