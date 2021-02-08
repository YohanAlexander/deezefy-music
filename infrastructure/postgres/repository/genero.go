package repository

import (
	"database/sql"
	"fmt"

	"github.com/yohanalexander/deezefy-music/entity"
)

// GeneroPSQL postgres repo
type GeneroPSQL struct {
	db *sql.DB
}

// NewGeneroPSQL create new repository
func NewGeneroPSQL(db *sql.DB) *GeneroPSQL {
	return &GeneroPSQL{
		db: db,
	}
}

// Create an Genero
func (r *GeneroPSQL) Create(e *entity.Genero) (string, error) {
	stmt, err := r.db.Prepare(`
		insert into deezefy.Genero (nome, estilo)
		values($1,$2)`)
	if err != nil {
		return e.Nome, err
	}
	_, err = stmt.Exec(
		e.Nome,
		e.Estilo,
	)
	if err != nil {
		return e.Nome, err
	}
	err = stmt.Close()
	if err != nil {
		return e.Nome, err
	}
	return e.Nome, nil
}

// Get an Genero
func (r *GeneroPSQL) Get(nome string) (*entity.Genero, error) {
	return getGenero(nome, r.db)
}

func getGenero(nome string, db *sql.DB) (*entity.Genero, error) {
	stmt, err := db.Prepare(`
		select nome, estilo from deezefy.Genero
		where nome = $1`)
	if err != nil {
		return nil, err
	}
	var u entity.Genero
	rows, err := stmt.Query(nome)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&u.Nome, &u.Estilo)
	}
	// select related artista
	stmt, err = db.Prepare(`
		select email, senha, data_nascimento, nome_artistico, biografia, ano_formacao from deezefy.Genero
		join deezefy.Artista_Possui_Genero on(Artista_Possui_Genero.fk_genero = Genero.nome)
		join deezefy.Artista on(Artista.fk_usuario = Artista_Possui_Genero.fk_artista)
		join deezefy.Usuario on(Usuario.email = Artista.fk_usuario)
		where Genero.nome = $1`)
	if err != nil {
		return nil, err
	}
	rows, err = stmt.Query(nome)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		j := &entity.Artista{}
		err = rows.Scan(&j.Usuario.Email, &j.Usuario.Password, &j.Usuario.Birthday, &j.NomeArtistico, &j.Biografia, &j.AnoFormacao)
		u.Artistas = append(u.Artistas, *j)
	}
	// select related musica
	stmt, err = db.Prepare(`
		select id, nome, duracao from deezefy.Musica
		join deezefy.Musica_Possui_Genero on(Musica_Possui_Genero.fk_musica = Musica.id)
		where Genero.nome = $1`)
	if err != nil {
		return nil, err
	}
	rows, err = stmt.Query(nome)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		j := &entity.Musica{}
		err = rows.Scan(&j.ID, &j.Nome, &j.Duracao)
		u.Musicas = append(u.Musicas, *j)
	}
	// select related perfil
	stmt, err = db.Prepare(`
		select email, senha, data_nascimento, primeiro_nome, sobrenome, id, informacoes_relevantes from deezefy.Genero
		join deezefy.Generos_Favoritos on(Generos_Favoritos.fk_genero = Genero.nome)
		where Genero.nome = $1`)
	if err != nil {
		return nil, err
	}
	rows, err = stmt.Query(nome)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		j := &entity.Perfil{}
		err = rows.Scan(&j.Ouvinte.Usuario.Email, &j.Ouvinte.Usuario.Password, &j.Ouvinte.Usuario.Birthday, &j.Ouvinte.PrimeiroNome, &j.Ouvinte.Sobrenome, &j.ID, &j.InformacoesRelevantes)
		u.Perfis = append(u.Perfis, *j)
	}
	return &u, nil
}

// Update an Genero
func (r *GeneroPSQL) Update(e *entity.Genero) error {
	_, err := r.db.Exec(`
		update deezefy.Genero set nome = $1, status = $2
		where id = ?`, e.Nome, e.Estilo)
	if err != nil {
		return err
	}
	// update related artista
	_, err = r.db.Exec(`
		delete from deezefy.Artista_Possui_Genero
		where fk_genero = $1`, e.Nome)
	if err != nil {
		return err
	}
	for _, b := range e.Artistas {
		_, err := r.db.Exec(`
		insert into deezefy.Artista_Possui_Genero (fk_artista, fk_genero)
		values($1,$2)`, b.Usuario.Email, e.Nome)
		if err != nil {
			return err
		}
	}
	// update related musica
	_, err = r.db.Exec(`
		delete from deezefy.Musica_Possui_Genero
		where fk_genero = $1`, e.Nome)
	if err != nil {
		return err
	}
	for _, b := range e.Musicas {
		_, err := r.db.Exec(`
		insert into deezefy.Musica_Possui_Genero (fk_musica, fk_genero)
		values($1,$2)`, b.ID, e.Nome)
		if err != nil {
			return err
		}
	}
	// update related perfil
	_, err = r.db.Exec(`
		delete from deezefy.Generos_Favoritos
		where fk_genero = $1`, e.Nome)
	if err != nil {
		return err
	}
	for _, b := range e.Perfis {
		_, err := r.db.Exec(`
		insert into deezefy.Generos_Favoritos (fk_perfil, fk_ouvinte, fk_genero)
		values($1,$2,$3)`, b.ID, b.Ouvinte.Usuario.Email, e.Nome)
		if err != nil {
			return err
		}
	}
	return nil
}

// Search Genero
func (r *GeneroPSQL) Search(query string) ([]*entity.Genero, error) {
	stmt, err := r.db.Prepare(`
		select nome from deezefy.Genero
		where nome like $1`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var ids []string
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
		ids = append(ids, i)
	}
	if len(ids) == 0 {
		return nil, fmt.Errorf("not found")
	}
	var Generos []*entity.Genero
	for _, id := range ids {
		u, err := getGenero(id, r.db)
		if err != nil {
			return nil, err
		}
		Generos = append(Generos, u)
	}
	return Generos, nil
}

// List Generos
func (r *GeneroPSQL) List() ([]*entity.Genero, error) {
	stmt, err := r.db.Prepare(`
		select nome from deezefy.Genero`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var ids []string
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
		ids = append(ids, i)
	}
	if len(ids) == 0 {
		return nil, fmt.Errorf("not found")
	}
	var Generos []*entity.Genero
	for _, id := range ids {
		u, err := getGenero(id, r.db)
		if err != nil {
			return nil, err
		}
		Generos = append(Generos, u)
	}
	return Generos, nil
}

// Delete an Genero
func (r *GeneroPSQL) Delete(nome string) error {
	_, err := r.db.Exec(`
		delete from deezefy.Genero
		where nome = $1`, nome)
	if err != nil {
		return err
	}
	return nil
}
