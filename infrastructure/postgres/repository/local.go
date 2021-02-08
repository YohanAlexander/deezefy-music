package repository

import (
	"database/sql"
	"fmt"

	"github.com/yohanalexander/deezefy-music/entity"
)

// LocalPSQL postgres repo
type LocalPSQL struct {
	db *sql.DB
}

// NewLocalPSQL create new repository
func NewLocalPSQL(db *sql.DB) *LocalPSQL {
	return &LocalPSQL{
		db: db,
	}
}

// Create an Local
func (r *LocalPSQL) Create(e *entity.Local) (int, error) {
	stmt, err := r.db.Prepare(`
		insert into deezefy.Local (id, cidade, pais)
		values($1,$2,$3)`)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.Cidade,
		e.Pais,
	)
	if err != nil {
		return e.ID, err
	}
	err = stmt.Close()
	if err != nil {
		return e.ID, err
	}
	return e.ID, nil
}

// Get an Local
func (r *LocalPSQL) Get(id int) (*entity.Local, error) {
	return getLocal(id, r.db)
}

func getLocal(id int, db *sql.DB) (*entity.Local, error) {
	stmt, err := db.Prepare(`
		select id, cidade, pais from deezefy.Local
		where id = $1`)
	if err != nil {
		return nil, err
	}
	var u entity.Local
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Cidade, &u.Pais)
	}
	return &u, nil
}

// Update an Local
func (r *LocalPSQL) Update(e *entity.Local) error {
	_, err := r.db.Exec(`
		update deezefy.Local set id = $1, cidade = $2, pais = $3
		where id = $4`, e.ID, e.Cidade, e.Pais, e.ID)
	if err != nil {
		return err
	}
	return nil
}

// Search Local
func (r *LocalPSQL) Search(query string) ([]*entity.Local, error) {
	stmt, err := r.db.Prepare(`
		select id from deezefy.Local
		where cidade like $1`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var ids []int
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var i int
		err = rows.Scan(&i)
		if err != nil {
			return nil, err
		}
		ids = append(ids, i)
	}
	if len(ids) == 0 {
		return nil, fmt.Errorf("not found")
	}
	var Locals []*entity.Local
	for _, id := range ids {
		u, err := getLocal(id, r.db)
		if err != nil {
			return nil, err
		}
		Locals = append(Locals, u)
	}
	return Locals, nil
}

// List Locals
func (r *LocalPSQL) List() ([]*entity.Local, error) {
	stmt, err := r.db.Prepare(`
		select id from deezefy.Local`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var ids []int
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var i int
		err = rows.Scan(&i)
		if err != nil {
			return nil, err
		}
		ids = append(ids, i)
	}
	if len(ids) == 0 {
		return nil, fmt.Errorf("not found")
	}
	var Locals []*entity.Local
	for _, id := range ids {
		u, err := getLocal(id, r.db)
		if err != nil {
			return nil, err
		}
		Locals = append(Locals, u)
	}
	return Locals, nil
}

// Delete an Local
func (r *LocalPSQL) Delete(id int) error {
	_, err := r.db.Exec(`
		delete from deezefy.Local
		where id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
