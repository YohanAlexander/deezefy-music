package ocorre

import (
	"strconv"
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/ocorre"
)

// Inmem in memory repo
type Inmem struct {
	m map[string]*der.Ocorre
}

// NewInmem create new repository
func NewInmem() *Inmem {
	var m = map[string]*der.Ocorre{}
	return &Inmem{
		m: m,
	}
}

// Create Ocorre
func (r *Inmem) Create(e *der.Ocorre) (string, string, int, int, error) {
	r.m[e.Artista+e.Usuario+strconv.Itoa(e.Local)+strconv.Itoa(e.Evento)] = e
	return e.Artista, e.Usuario, e.Local, e.Evento, nil
}

// Get Ocorre
func (r *Inmem) Get(artista, usuario string, local, evento int) (*der.Ocorre, error) {
	if r.m[artista+usuario+strconv.Itoa(local)+strconv.Itoa(evento)] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[artista+usuario+strconv.Itoa(local)+strconv.Itoa(evento)], nil
}

// GetByLocal Ocorre
func (r *Inmem) GetByLocal(local int) (*der.Ocorre, error) {
	if r.m[strconv.Itoa(local)] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(local)], nil
}

// GetByEvento Ocorre
func (r *Inmem) GetByEvento(evento int) (*der.Ocorre, error) {
	if r.m[strconv.Itoa(evento)] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[strconv.Itoa(evento)], nil
}

// GetByArtista Ocorre
func (r *Inmem) GetByArtista(artista string) (*der.Ocorre, error) {
	if r.m[artista] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[artista], nil
}

// Update Ocorre
func (r *Inmem) Update(e *der.Ocorre) error {
	_, err := r.Get(e.Artista, e.Usuario, e.Local, e.Evento)
	if err != nil {
		return err
	}
	r.m[e.Artista+e.Usuario+strconv.Itoa(e.Local)+strconv.Itoa(e.Evento)] = e
	return nil
}

// Search Ocorres
func (r *Inmem) Search(query string) ([]*der.Ocorre, error) {
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
func (r *Inmem) List() ([]*der.Ocorre, error) {
	var d []*der.Ocorre
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete Ocorre
func (r *Inmem) Delete(artista, usuario string, local, evento int) error {
	if r.m[artista+usuario+strconv.Itoa(local)+strconv.Itoa(evento)] == nil {
		return entity.ErrNotFound
	}
	r.m[artista+usuario+strconv.Itoa(local)+strconv.Itoa(evento)] = nil
	return nil
}
