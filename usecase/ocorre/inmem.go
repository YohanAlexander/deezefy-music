package ocorre

import (
	"strconv"
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/ocorre"
)

// inmem in memory repo
type inmem struct {
	m map[string]*der.Ocorre
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[string]*der.Ocorre{}
	return &inmem{
		m: m,
	}
}

// Create Ocorre
func (r *inmem) Create(e *der.Ocorre) (string, string, int, int, error) {
	r.m[e.Artista+e.Usuario+strconv.Itoa(e.Local)+strconv.Itoa(e.Evento)] = e
	return e.Artista, e.Usuario, e.Local, e.Evento, nil
}

// Get Ocorre
func (r *inmem) Get(artista, usuario string, local, evento int) (*der.Ocorre, error) {
	if r.m[artista+usuario+strconv.Itoa(local)+strconv.Itoa(evento)] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[artista+usuario+strconv.Itoa(local)+strconv.Itoa(evento)], nil
}

// GetByLocal Ocorre
func (r *inmem) GetByLocal(local int) (*der.Ocorre, error) {
	if r.m[strconv.Itoa(local)] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(local)], nil
}

// GetByEvento Ocorre
func (r *inmem) GetByEvento(evento int) (*der.Ocorre, error) {
	if r.m[strconv.Itoa(evento)] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(evento)], nil
}

// GetByArtista Ocorre
func (r *inmem) GetByArtista(artista string) (*der.Ocorre, error) {
	if r.m[artista] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[artista], nil
}

// Update Ocorre
func (r *inmem) Update(e *der.Ocorre) error {
	_, err := r.Get(e.Artista, e.Usuario, e.Local, e.Evento)
	if err != nil {
		return err
	}
	r.m[e.Artista+e.Usuario+strconv.Itoa(e.Local)+strconv.Itoa(e.Evento)] = e
	return nil
}

// Search Ocorres
func (r *inmem) Search(query string) ([]*der.Ocorre, error) {
	var d []*der.Ocorre
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

// List Ocorres
func (r *inmem) List() ([]*der.Ocorre, error) {
	var d []*der.Ocorre
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete Ocorre
func (r *inmem) Delete(artista, usuario string, local, evento int) error {
	if r.m[artista+usuario+strconv.Itoa(local)+strconv.Itoa(evento)] == nil {
		return entity.ErrNotFound
	}
	r.m[artista+usuario+strconv.Itoa(local)+strconv.Itoa(evento)] = nil
	return nil
}
