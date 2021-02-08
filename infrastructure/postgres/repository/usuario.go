package repository

import (
	"database/sql"
	"fmt"

	"github.com/yohanalexander/deezefy-music/entity"
)

// UsuarioPSQL postgres repo
type UsuarioPSQL struct {
	db *sql.DB
}

// NewUsuarioPSQL create new repository
func NewUsuarioPSQL(db *sql.DB) *UsuarioPSQL {
	return &UsuarioPSQL{
		db: db,
	}
}

// Create an usuario
func (r *UsuarioPSQL) Create(e *entity.Usuario) (string, error) {
	stmt, err := r.db.Prepare(`
		insert into deezefy.Usuario (email, senha, data_nascimento)
		values(?,?,?)`)
	if err != nil {
		return e.Email, err
	}
	_, err = stmt.Exec(
		e.Email,
		e.Password,
		e.Birthday,
	)
	if err != nil {
		return e.Email, err
	}
	err = stmt.Close()
	if err != nil {
		return e.Email, err
	}
	return e.Email, nil
}

// Get an usuario
func (r *UsuarioPSQL) Get(email string) (*entity.Usuario, error) {
	return getUsuario(email, r.db)
}

func getUsuario(email string, db *sql.DB) (*entity.Usuario, error) {
	stmt, err := db.Prepare(`select email, senha, data_nascimento from deezefy.Usuario where email = ?`)
	if err != nil {
		return nil, err
	}
	var u entity.Usuario
	rows, err := stmt.Query(email)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&u.Email, &u.Password, &u.Birthday)
	}
	return &u, nil
}

// Update an usuario
func (r *UsuarioPSQL) Update(e *entity.Usuario) error {
	_, err := r.db.Exec(`update deezefy.Usuario set email = ?, password = ?, data_nascimento = ? where email = ?`, e.Email, e.Password, e.Birthday, e.Email)
	if err != nil {
		return err
	}
	return nil
}

// Search usuario
func (r *UsuarioPSQL) Search(query string) ([]*entity.Usuario, error) {
	stmt, err := r.db.Prepare(`select email from deezefy.Usuario where email like ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var emails []string
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var i string
		err = rows.Scan(&i)
		if err != nil {
			return nil, err
		}
		emails = append(emails, i)
	}
	if len(emails) == 0 {
		return nil, fmt.Errorf("not found")
	}
	var usuarios []*entity.Usuario
	for _, email := range emails {
		u, err := getUsuario(email, r.db)
		if err != nil {
			return nil, err
		}
		usuarios = append(usuarios, u)
	}
	return usuarios, nil
}

// List usuarios
func (r *UsuarioPSQL) List() ([]*entity.Usuario, error) {
	stmt, err := r.db.Prepare(`select email from deezefy.Usuario`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var emails []string
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var i string
		err = rows.Scan(&i)
		if err != nil {
			return nil, err
		}
		emails = append(emails, i)
	}
	if len(emails) == 0 {
		return nil, fmt.Errorf("not found")
	}
	var usuarios []*entity.Usuario
	for _, email := range emails {
		u, err := getUsuario(email, r.db)
		if err != nil {
			return nil, err
		}
		usuarios = append(usuarios, u)
	}
	return usuarios, nil
}

// Delete an usuario
func (r *UsuarioPSQL) Delete(email string) error {
	_, err := r.db.Exec(`delete from deezefy.Usuario where email = ?`, email)
	if err != nil {
		return err
	}
	return nil
}
