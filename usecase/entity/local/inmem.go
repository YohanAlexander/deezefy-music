package local

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
)

// inmem in memory repo
type inmem struct {
	m map[int]*entity.Local
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[int]*entity.Local{}
	return &inmem{
		m: m,
	}
}

// Create Local
func (r *inmem) Create(e *entity.Local) (int, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

// Get Local
func (r *inmem) Get(id int) (*entity.Local, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

// Update Local
func (r *inmem) Update(e *entity.Local) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

// Search Locals
func (r *inmem) Search(query string) ([]*entity.Local, error) {
	var d []*entity.Local
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
func (r *inmem) List() ([]*entity.Local, error) {
	var d []*entity.Local
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete Local
func (r *inmem) Delete(id int) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}
