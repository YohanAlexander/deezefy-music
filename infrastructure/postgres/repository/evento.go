package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/yohanalexander/deezefy-music/entity"
)

// EventoPSQL postgres repo
type EventoPSQL struct {
	db *sql.DB
}

// NewEventoPSQL create new repository
func NewEventoPSQL(db *sql.DB) *EventoPSQL {
	return &EventoPSQL{
		db: db,
	}
}

// Create an Evento
func (r *EventoPSQL) Create(e *entity.Evento) (int, error) {
	stmt, err := r.db.Prepare(`
		insert into deezefy.Evento (id, nome, fk_usuario)
		values(?,?,?)`)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.Nome,
		e.Usuario.Email,
	)
	if err != nil {
		return e.ID, err
	}
	stmt, err = r.db.Prepare(`
		insert into deezefy.Ocorre (data, fk_artista, fk_local, fk_evento, fk_usuario)
		values(?,?,?,?,?)`)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		time.Now().Format("2006-01-02"),
		e.Usuario.Email,
		e.Local.ID,
		e.ID,
		e.Usuario.Email,
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

// Get an Evento
func (r *EventoPSQL) Get(id int) (*entity.Evento, error) {
	return getEvento(id, r.db)
}

func getEvento(id int, db *sql.DB) (*entity.Evento, error) {
	stmt, err := db.Prepare(`select email, senha, data_nascimento, id, cidade, pais, id, nome, data from deezefy.Evento
	join deezefy.Ocorre on(Ocorre.fk_evento = Evento.id)
	join deezefy.Local on(Local.id = Ocorre.fk_local)
	join deezefy.Usuario on(Usuario.email = Ocorre.fk_usuario)
	where Evento.id = ?`)
	if err != nil {
		return nil, err
	}
	var u entity.Evento
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&u.Usuario.Email, &u.Usuario.Password, &u.Usuario.Birthday, &u.Local.ID, &u.Local.Cidade, &u.Local.Pais, &u.ID, &u.Nome, &u.Data)
	}
	return &u, nil
}

// Update an Evento
func (r *EventoPSQL) Update(e *entity.Evento) error {
	_, err := r.db.Exec(`update deezefy.Evento set id = ?, nome = ?
	where id = ?`, e.ID, e.Nome)
	if err != nil {
		return err
	}
	return nil
}

// Search Evento
func (r *EventoPSQL) Search(query string) ([]*entity.Evento, error) {
	stmt, err := r.db.Prepare(`select nome from deezefy.Evento where nome like ?`)
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
	var Eventos []*entity.Evento
	for _, id := range ids {
		u, err := getEvento(id, r.db)
		if err != nil {
			return nil, err
		}
		Eventos = append(Eventos, u)
	}
	return Eventos, nil
}

// List Eventos
func (r *EventoPSQL) List() ([]*entity.Evento, error) {
	stmt, err := r.db.Prepare(`select id from deezefy.Evento`)
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
	var Eventos []*entity.Evento
	for _, id := range ids {
		u, err := getEvento(id, r.db)
		if err != nil {
			return nil, err
		}
		Eventos = append(Eventos, u)
	}
	return Eventos, nil
}

// Delete an Evento
func (r *EventoPSQL) Delete(id int) error {
	_, err := r.db.Exec(`delete from deezefy.Evento where id = ?`, id)
	if err != nil {
		return err
	}
	return nil
}
