package evento

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/evento"
)

// Inmem in memory repo
type Inmem struct {
	m map[int]*der.Evento
}

// NewInmem create new repository
func NewInmem() *Inmem {
	var m = map[int]*der.Evento{}
	return &Inmem{
		m: m,
	}
}

// Create Evento
func (r *Inmem) Create(e *der.Evento) (int, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

// Get Evento
func (r *Inmem) Get(id int) (*der.Evento, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

// Update Evento
func (r *Inmem) Update(e *der.Evento) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

// Search Eventos
func (r *Inmem) Search(query string) ([]*der.Evento, error) {
	var d []*der.Evento
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

// List Eventos
func (r *Inmem) List() ([]*der.Evento, error) {
	var d []*der.Evento
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete Evento
func (r *Inmem) Delete(id int) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}
