package evento

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/evento"
)

// inmem in memory repo
type inmem struct {
	m map[int]*der.Evento
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[int]*der.Evento{}
	return &inmem{
		m: m,
	}
}

// Create Evento
func (r *inmem) Create(e *der.Evento) (int, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

// Get Evento
func (r *inmem) Get(id int) (*der.Evento, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

// Update Evento
func (r *inmem) Update(e *der.Evento) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

// Search Eventos
func (r *inmem) Search(query string) ([]*der.Evento, error) {
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
func (r *inmem) List() ([]*der.Evento, error) {
	var d []*der.Evento
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete Evento
func (r *inmem) Delete(id int) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}
