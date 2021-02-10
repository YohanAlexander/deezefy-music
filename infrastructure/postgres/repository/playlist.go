package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/yohanalexander/deezefy-music/entity"
)

// PlaylistPSQL postgres repo
type PlaylistPSQL struct {
	db *sql.DB
}

// NewPlaylistPSQL create new repository
func NewPlaylistPSQL(db *sql.DB) *PlaylistPSQL {
	return &PlaylistPSQL{
		db: db,
	}
}

// Create an Playlist
func (r *PlaylistPSQL) Create(e *entity.Playlist) (string, error) {
	stmt, err := r.db.Prepare(`
		insert into deezefy.Playlist (nome, status)
		values($1,$2)`)
	if err != nil {
		return e.Nome, err
	}
	_, err = stmt.Exec(
		e.Nome,
		e.Status,
	)
	if err != nil {
		return e.Nome, err
	}
	stmt, err = r.db.Prepare(`
		insert into deezefy.Cria (data_criacao, fk_playlist, fk_usuario)
		values($1,$2,$3)`)
	if err != nil {
		return e.Nome, err
	}
	_, err = stmt.Exec(
		time.Now().Format("2006-01-02"),
		e.Nome,
		e.Usuario.Email,
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

// Get an Playlist
func (r *PlaylistPSQL) Get(nome string) (*entity.Playlist, error) {
	return getPlaylist(nome, r.db)
}

func getPlaylist(nome string, db *sql.DB) (*entity.Playlist, error) {
	stmt, err := db.Prepare(`
		select email, senha, data_nascimento, nome, status, data_criacao from deezefy.Playlist
		join deezefy.Cria on(Cria.fk_playlist = Playlist.nome)
		join deezefy.Usuario on(Cria.fk_usuario = Usuario.email)
		where nome = $1`)
	if err != nil {
		return nil, err
	}
	var u entity.Playlist
	rows, err := stmt.Query(nome)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&u.Usuario.Email, &u.Usuario.Password, &u.Usuario.Birthday, &u.Nome, &u.Status, &u.DataCriacao)
	}
	// select related ouvinte
	stmt, err = db.Prepare(`
		select email, senha, data_nascimento, primeiro_nome, sobrenome from deezefy.Playlist
		join deezefy.Ouvinte_Salva_Playlist on(Ouvinte_Salva_Playlist.fk_playlist = Playlist.nome)
		join deezefy.Ouvinte on(Ouvinte.fk_usuario = Ouvinte_Salva_Playlist.fk_ouvinte)
		join deezefy.Usuario on(Usuario.email = Ouvinte.fk_usuario)
		where Playlist.nome = $1`)
	if err != nil {
		return nil, err
	}
	rows, err = stmt.Query(nome)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		j := &entity.Ouvinte{}
		err = rows.Scan(&j.Usuario.Email, &j.Usuario.Password, &j.Usuario.Birthday, &j.PrimeiroNome, &j.Sobrenome)
		u.Salvou = append(u.Salvou, *j)
	}
	// select related musica
	stmt, err = db.Prepare(`
		select id, Musica.nome, duracao from deezefy.Musica
		join deezefy.Musica_em_Playlist on(Musica_em_Playlist.fk_musica = Musica.id)
		join deezefy.Playlist on(Musica_em_Playlist.fk_playlist = Playlist.nome)
		where Playlist.nome = $1`)
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
	return &u, nil
}

// Update an Playlist
func (r *PlaylistPSQL) Update(e *entity.Playlist) error {
	_, err := r.db.Exec(`
		update deezefy.Playlist set nome = $1, status = $2
		where id = $3`, e.Nome, e.Status, e.Nome)
	if err != nil {
		return err
	}
	// update related ouvinte
	_, err = r.db.Exec(`
		delete from deezefy.Ouvinte_Salva_Playlist
		where fk_playlist = $1`, e.Nome)
	if err != nil {
		return err
	}
	for _, b := range e.Salvou {
		_, err := r.db.Exec(`
		insert into deezefy.Ouvinte_Salva_Playlist (fk_ouvinte, fk_playlist)
		values($1,$2)`, b.Usuario.Email, e.Nome)
		if err != nil {
			return err
		}
	}
	// update related musica
	_, err = r.db.Exec(`
		delete from deezefy.Musica_em_Playlist
		where fk_playlist = $1`, e.Nome)
	if err != nil {
		return err
	}
	for _, b := range e.Musicas {
		_, err := r.db.Exec(`
		insert into deezefy.Musica_em_Playlist (fk_musica, fk_playlist)
		values($1,$2)`, b.ID, e.Nome)
		if err != nil {
			return err
		}
	}
	return nil
}

// Search Playlist
func (r *PlaylistPSQL) Search(query string) ([]*entity.Playlist, error) {
	stmt, err := r.db.Prepare(`
		select nome from deezefy.Playlist
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
	var Playlists []*entity.Playlist
	for _, id := range ids {
		u, err := getPlaylist(id, r.db)
		if err != nil {
			return nil, err
		}
		Playlists = append(Playlists, u)
	}
	return Playlists, nil
}

// List Playlists
func (r *PlaylistPSQL) List() ([]*entity.Playlist, error) {
	stmt, err := r.db.Prepare(`
		select nome from deezefy.Playlist`)
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
	var Playlists []*entity.Playlist
	for _, id := range ids {
		u, err := getPlaylist(id, r.db)
		if err != nil {
			return nil, err
		}
		Playlists = append(Playlists, u)
	}
	return Playlists, nil
}

// Delete an Playlist
func (r *PlaylistPSQL) Delete(nome string) error {
	_, err := r.db.Exec(`
		delete from deezefy.Playlist
		where nome = $1`, nome)
	if err != nil {
		return err
	}
	return nil
}
