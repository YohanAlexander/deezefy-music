package musicagenero

import (
	"strconv"
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/musicagenero"
)

// inmem in memory repo
type inmem struct {
	m map[string]*der.MusicaGenero
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[string]*der.MusicaGenero{}
	return &inmem{
		m: m,
	}
}

// Create MusicaGenero
func (r *inmem) Create(e *der.MusicaGenero) (int, string, error) {
	r.m[strconv.Itoa(e.Musica)+e.Genero] = e
	return e.Musica, e.Genero, nil
}

// Get MusicaGenero
func (r *inmem) Get(musica int, genero string) (*der.MusicaGenero, error) {
	if r.m[strconv.Itoa(musica)+genero] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(musica)+genero], nil
}

// GetByMusica MusicaGenero
func (r *inmem) GetByMusica(musica int) (*der.MusicaGenero, error) {
	if r.m[strconv.Itoa(musica)] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(musica)], nil
}

// GetByGenero MusicaGenero
func (r *inmem) GetByGenero(genero string) (*der.MusicaGenero, error) {
	if r.m[genero] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[genero], nil
}

// Update MusicaGenero
func (r *inmem) Update(e *der.MusicaGenero) error {
	_, err := r.Get(e.Musica, e.Genero)
	if err != nil {
		return err
	}
	r.m[strconv.Itoa(e.Musica)+e.Genero] = e
	return nil
}

// Search MusicaGeneros
func (r *inmem) Search(query string) ([]*der.MusicaGenero, error) {
	var d []*der.MusicaGenero
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Genero), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, entity.ErrNotFound
	}

	return d, nil
}

// List MusicaGeneros
func (r *inmem) List() ([]*der.MusicaGenero, error) {
	var d []*der.MusicaGenero
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete MusicaGenero
func (r *inmem) Delete(musica int, genero string) error {
	if r.m[strconv.Itoa(musica)+genero] == nil {
		return entity.ErrNotFound
	}
	r.m[strconv.Itoa(musica)+genero] = nil
	return nil
}
