package repository

import (
	"database/sql"
	"fmt"

	"github.com/yohanalexander/deezefy-music/entity"
)

// PerfilPSQL postgres repo
type PerfilPSQL struct {
	db *sql.DB
}

// NewPerfilPSQL create new repository
func NewPerfilPSQL(db *sql.DB) *PerfilPSQL {
	return &PerfilPSQL{
		db: db,
	}
}

// Create an Perfil
func (r *PerfilPSQL) Create(e *entity.Perfil) (int, error) {
	stmt, err := r.db.Prepare(`
		insert into deezefy.Perfil (id, informacoes_relevantes)
		values($1,$2)`)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.InformacoesRelevantes,
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

// Get an Perfil
func (r *PerfilPSQL) Get(id int) (*entity.Perfil, error) {
	return getPerfil(id, r.db)
}

func getPerfil(id int, db *sql.DB) (*entity.Perfil, error) {
	stmt, err := db.Prepare(`
		select id, informacoes_relevantes from deezefy.Perfil
		where id = $1`)
	if err != nil {
		return nil, err
	}
	var u entity.Perfil
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&u.ID, &u.InformacoesRelevantes)
	}
	// select related artista
	stmt, err = db.Prepare(`
		select email, senha, data_nascimento, nome_artistico, biografia, ano_formacao from deezefy.Perfil
		join deezefy.Artistas_Favoritos on(Artistas_Favoritos.fk_perfil = Perfil.id)
		join deezefy.Artista on(Artista.fk_usuario = Artistas_Favoritos.fk_artista)
		join deezefy.Usuario on(Usuario.email = Artista.fk_usuario)
		where Perfil.id = $1`)
	if err != nil {
		return nil, err
	}
	rows, err = stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		j := &entity.Artista{}
		err = rows.Scan(&j.Usuario.Email, &j.Usuario.Password, &j.Usuario.Birthday, &j.NomeArtistico, &j.Biografia, &j.AnoFormacao)
		u.ArtistasFavoritos = append(u.ArtistasFavoritos, *j)
	}
	// select related genero
	stmt, err = db.Prepare(`
		select nome, estilo from deezefy.Perfil
		join deezefy.Generos_Favoritos on(Generos_Favoritos.fk_perfil = Perfil.id)
		where Perfil.id = $1`)
	if err != nil {
		return nil, err
	}
	rows, err = stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		j := &entity.Genero{}
		err = rows.Scan(&j.Nome, &j.Estilo)
		u.GenerosFavoritos = append(u.GenerosFavoritos, *j)
	}
	return &u, nil
}

// Update an Perfil
func (r *PerfilPSQL) Update(e *entity.Perfil) error {
	_, err := r.db.Exec(`
		update deezefy.Perfil set id = $1, informacoes_relevantes = $2
		where id = $3`, e.ID, e.InformacoesRelevantes, e.ID)
	if err != nil {
		return err
	}
	// update related artista
	_, err = r.db.Exec(`
		delete from deezefy.Artistas_Favoritos
		where fk_perfil = $1`, e.ID)
	if err != nil {
		return err
	}
	for _, b := range e.ArtistasFavoritos {
		_, err := r.db.Exec(`
		insert into deezefy.Artistas_Favoritos (fk_artista, fk_ouvinte, fk_perfil)
		values($1,$2,$3)`, b.Usuario.Email, e.Ouvinte.Usuario.Email, e.ID)
		if err != nil {
			return err
		}
	}
	// update related genero
	_, err = r.db.Exec(`
		delete from deezefy.Generos_Favoritos
		where fk_perfil = $1`, e.ID)
	if err != nil {
		return err
	}
	for _, b := range e.GenerosFavoritos {
		_, err := r.db.Exec(`
		insert into deezefy.Generos_Favoritos (fk_genero, fk_ouvinte, fk_perfil)
		values($1,$2,$3)`, b.Nome, e.Ouvinte.Usuario.Email, e.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

// Search Perfil
func (r *PerfilPSQL) Search(query string) ([]*entity.Perfil, error) {
	stmt, err := r.db.Prepare(`
		select id from deezefy.Perfil
		join deezefy.Ouvinte on(Ouvinte.fk_usuario = Perfil.fk_ouvinte)
		join deezefy.Usuario on(Usuario.email = Ouvinte.fk_usuario)
		where email like $1`)
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
	var Perfils []*entity.Perfil
	for _, id := range ids {
		u, err := getPerfil(id, r.db)
		if err != nil {
			return nil, err
		}
		Perfils = append(Perfils, u)
	}
	return Perfils, nil
}

// List Perfils
func (r *PerfilPSQL) List() ([]*entity.Perfil, error) {
	stmt, err := r.db.Prepare(`
		select id from deezefy.Perfil`)
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
	var Perfils []*entity.Perfil
	for _, id := range ids {
		u, err := getPerfil(id, r.db)
		if err != nil {
			return nil, err
		}
		Perfils = append(Perfils, u)
	}
	return Perfils, nil
}

// Delete an Perfil
func (r *PerfilPSQL) Delete(id int) error {
	_, err := r.db.Exec(`
		delete from deezefy.Perfil
		where id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
