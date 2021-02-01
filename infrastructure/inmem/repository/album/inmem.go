package album

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/album"
)

// Inmem in memory repo
type Inmem struct {
	m map[int]*der.Album
}

// NewInmem create new repository
func NewInmem() *Inmem {
	var m = map[int]*der.Album{}
	return &Inmem{
		m: m,
	}
}

// Create Album
func (r *Inmem) Create(e *der.Album) (int, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

// Get Album
func (r *Inmem) Get(id int) (*der.Album, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

// Update Album
func (r *Inmem) Update(e *der.Album) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

// Search Albums
func (r *Inmem) Search(query string) ([]*der.Album, error) {
	var d []*der.Album
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Titulo), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, entity.ErrNotFound
	}

	return d, nil
}

// List Albums
func (r *Inmem) List() ([]*der.Album, error) {
	var d []*der.Album
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete Album
func (r *Inmem) Delete(id int) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}
