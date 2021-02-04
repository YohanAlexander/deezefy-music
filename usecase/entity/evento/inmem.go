package evento

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
)

// inmem in memory repo
type inmem struct {
	m map[int]*entity.Evento
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[int]*entity.Evento{}
	return &inmem{
		m: m,
	}
}

// Create Evento
func (r *inmem) Create(e *entity.Evento) (int, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

// Get Evento
func (r *inmem) Get(id int) (*entity.Evento, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

// Update Evento
func (r *inmem) Update(e *entity.Evento) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

// Search Eventos
func (r *inmem) Search(query string) ([]*entity.Evento, error) {
	var d []*entity.Evento
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
func (r *inmem) List() ([]*entity.Evento, error) {
	var d []*entity.Evento
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
