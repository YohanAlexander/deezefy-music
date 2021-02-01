package musica

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/musica"
)

// Inmem in memory repo
type Inmem struct {
	m map[int]*der.Musica
}

// NewInmem create new repository
func NewInmem() *Inmem {
	var m = map[int]*der.Musica{}
	return &Inmem{
		m: m,
	}
}

// Create Musica
func (r *Inmem) Create(e *der.Musica) (int, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

// Get Musica
func (r *Inmem) Get(id int) (*der.Musica, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

// Update Musica
func (r *Inmem) Update(e *der.Musica) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

// Search Musicas
func (r *Inmem) Search(query string) ([]*der.Musica, error) {
	var d []*der.Musica
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

// List Musicas
func (r *Inmem) List() ([]*der.Musica, error) {
	var d []*der.Musica
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete Musica
func (r *Inmem) Delete(id int) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}
