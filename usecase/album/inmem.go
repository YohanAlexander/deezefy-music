package album

import (
	"strings"

	"github.com/yohanalexander/deezefy-music/entity"
	der "github.com/yohanalexander/deezefy-music/entity/album"
)

// inmem in memory repo
type inmem struct {
	m map[int]*der.Album
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[int]*der.Album{}
	return &inmem{
		m: m,
	}
}

// Create Album
func (r *inmem) Create(e *der.Album) (int, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

// Get Album
func (r *inmem) Get(id int) (*der.Album, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

// Update Album
func (r *inmem) Update(e *der.Album) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

// Search Albums
func (r *inmem) Search(query string) ([]*der.Album, error) {
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
func (r *inmem) List() ([]*der.Album, error) {
	var d []*der.Album
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete Album
func (r *inmem) Delete(id int) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}
