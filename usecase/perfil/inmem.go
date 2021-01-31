package perfil

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/perfil"
)

// inmem in memory repo
type inmem struct {
	m map[int]*der.Perfil
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[int]*der.Perfil{}
	return &inmem{
		m: m,
	}
}

// Create Perfil
func (r *inmem) Create(e *der.Perfil) (int, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

// Get Perfil
func (r *inmem) Get(id int) (*der.Perfil, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

// Update Perfil
func (r *inmem) Update(e *der.Perfil) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

// Search Perfils
func (r *inmem) Search(query string) ([]*der.Perfil, error) {
	var d []*der.Perfil
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.InformacoesRelevantes), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, entity.ErrNotFound
	}

	return d, nil
}

// List Perfils
func (r *inmem) List() ([]*der.Perfil, error) {
	var d []*der.Perfil
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete Perfil
func (r *inmem) Delete(id int) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}
