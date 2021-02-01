package local

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/local"
)

// Inmem in memory repo
type Inmem struct {
	m map[int]*der.Local
}

// NewInmem create new repository
func NewInmem() *Inmem {
	var m = map[int]*der.Local{}
	return &Inmem{
		m: m,
	}
}

// Create Local
func (r *Inmem) Create(e *der.Local) (int, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

// Get Local
func (r *Inmem) Get(id int) (*der.Local, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

// Update Local
func (r *Inmem) Update(e *der.Local) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

// Search Locals
func (r *Inmem) Search(query string) ([]*der.Local, error) {
	var d []*der.Local
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Cidade), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, entity.ErrNotFound
	}

	return d, nil
}

// List Locals
func (r *Inmem) List() ([]*der.Local, error) {
	var d []*der.Local
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete Local
func (r *Inmem) Delete(id int) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}
