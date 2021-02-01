package curte

import (
	"strconv"
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/curte"
)

// Inmem in memory repo
type Inmem struct {
	m map[string]*der.Curte
}

// NewInmem create new repository
func NewInmem() *Inmem {
	var m = map[string]*der.Curte{}
	return &Inmem{
		m: m,
	}
}

// Create Curte
func (r *Inmem) Create(e *der.Curte) (int, string, error) {
	r.m[strconv.Itoa(e.Musica)+e.Ouvinte] = e
	return e.Musica, e.Ouvinte, nil
}

// Get Curte
func (r *Inmem) Get(musica int, ouvinte string) (*der.Curte, error) {
	if r.m[strconv.Itoa(musica)+ouvinte] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(musica)+ouvinte], nil
}

// GetByMusica Curte
func (r *Inmem) GetByMusica(musica int) (*der.Curte, error) {
	if r.m[strconv.Itoa(musica)] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(musica)], nil
}

// GetByOuvinte Curte
func (r *Inmem) GetByOuvinte(ouvinte string) (*der.Curte, error) {
	if r.m[ouvinte] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[ouvinte], nil
}

// Update Curte
func (r *Inmem) Update(e *der.Curte) error {
	_, err := r.Get(e.Musica, e.Ouvinte)
	if err != nil {
		return err
	}
	r.m[strconv.Itoa(e.Musica)+e.Ouvinte] = e
	return nil
}

// Search Curtes
func (r *Inmem) Search(query string) ([]*der.Curte, error) {
	var d []*der.Curte
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

// List Curtes
func (r *Inmem) List() ([]*der.Curte, error) {
	var d []*der.Curte
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete Curte
func (r *Inmem) Delete(musica int, ouvinte string) error {
	if r.m[strconv.Itoa(musica)+ouvinte] == nil {
		return entity.ErrNotFound
	}
	r.m[strconv.Itoa(musica)+ouvinte] = nil
	return nil
}
