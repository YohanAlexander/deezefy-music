package repository

import (
	"database/sql"
	"fmt"

	"github.com/yohanalexander/deezefy-music/entity"
)

// MusicaPSQL postgres repo
type MusicaPSQL struct {
	db *sql.DB
}

// NewMusicaPSQL create new repository
func NewMusicaPSQL(db *sql.DB) *MusicaPSQL {
	return &MusicaPSQL{
		db: db,
	}
}

// Create an Musica
func (r *MusicaPSQL) Create(e *entity.Musica) (int, error) {
	stmt, err := r.db.Prepare(`
		insert into deezefy.Musica (id, nome, duracao)
		values($1,$2,$3)`)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.Nome,
		e.Duracao,
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

// Get an Musica
func (r *MusicaPSQL) Get(id int) (*entity.Musica, error) {
	return getMusica(id, r.db)
}

func getMusica(id int, db *sql.DB) (*entity.Musica, error) {
	stmt, err := db.Prepare(`
		select id, nome, duracao from deezefy.Musica
		where id = $1`)
	if err != nil {
		return nil, err
	}
	var u entity.Musica
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Nome, &u.Duracao)
	}
	// select related ouvinte
	stmt, err = db.Prepare(`
		select email, senha, data_nascimento, primeiro_nome, sobrenome from deezefy.Musica
		join deezefy.Curte on(Curte.fk_musica = Musica.id)
		join deezefy.Ouvinte on(Ouvinte.fk_usuario = Curte.fk_ouvinte)
		join deezefy.Usuario on(Usuario.email = Ouvinte.fk_usuario)
		where Musica.id = $1`)
	if err != nil {
		return nil, err
	}
	rows, err = stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		j := &entity.Ouvinte{}
		err = rows.Scan(&j.Usuario.Email, &j.Usuario.Password, &j.Usuario.Birthday, &j.PrimeiroNome, &j.Sobrenome)
		u.Curtiu = append(u.Curtiu, *j)
	}
	// select related artista
	stmt, err = db.Prepare(`
		select email, senha, data_nascimento, nome_artistico, biografia, ano_formacao from deezefy.Musica
		join deezefy.Grava on(Grava.fk_musica = Musica.id)
		join deezefy.Artista on(Artista.fk_usuario = Grava.fk_artista)
		join deezefy.Usuario on(Usuario.email = Artista.fk_usuario)
		where Musica.id = $1`)
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
		u.Gravou = append(u.Gravou, *j)
	}
	// select related playlist
	stmt, err = db.Prepare(`
		select nome, status, data_criacao from deezefy.Playlist
		join deezefy.Musica_em_Playlist on(Playlist.nome = Musica_em_Playlist.fk_playlist)
		where Musica.id = $1`)
	if err != nil {
		return nil, err
	}
	rows, err = stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		j := &entity.Playlist{}
		err = rows.Scan(&j.Nome, &j.Status, &j.DataCriacao)
		u.Playlists = append(u.Playlists, *j)
	}
	// select related album
	stmt, err = db.Prepare(`
		select id, titulo, ano_lancamento from deezefy.Album
		join deezefy.Album_Contem_Musica on(Album_Contem_Musica.fk_musica = Musica.id)
		where Musica.id = $1`)
	if err != nil {
		return nil, err
	}
	rows, err = stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		j := &entity.Album{}
		err = rows.Scan(&j.ID, &j.Titulo, &j.AnoLancamento)
		u.Albums = append(u.Albums, *j)
	}
	// select related genero
	stmt, err = db.Prepare(`
		select nome, estilo from deezefy.Musica
		join deezefy.Musica_Possui_Genero on(Musica_Possui_Genero.fk_musica = Musica.id)
		where Musica.id = $1`)
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
		u.Generos = append(u.Generos, *j)
	}
	return &u, nil
}

// Update an Musica
func (r *MusicaPSQL) Update(e *entity.Musica) error {
	_, err := r.db.Exec(`
		update deezefy.Musica set id = $1, nome = $2, duracao = $3
		where id = $4`, e.ID, e.Nome, e.Duracao, e.ID)
	if err != nil {
		return err
	}
	// update related ouvinte
	_, err = r.db.Exec(`
		delete from deezefy.Curte
		where fk_musica = $1`, e.ID)
	if err != nil {
		return err
	}
	for _, b := range e.Curtiu {
		_, err := r.db.Exec(`
		insert into deezefy.Curte (fk_ouvinte, fk_musica)
		values($1,$2)`, b.Usuario.Email, e.ID)
		if err != nil {
			return err
		}
	}
	// update related artista
	_, err = r.db.Exec(`
		delete from deezefy.Grava
		where fk_musica = $1`, e.ID)
	if err != nil {
		return err
	}
	for _, b := range e.Gravou {
		_, err := r.db.Exec(`
		insert into deezefy.Grava (fk_artista, fk_musica)
		values($1,$2)`, b.Usuario.Email, e.ID)
		if err != nil {
			return err
		}
	}
	// update related playlist
	_, err = r.db.Exec(`
		delete from deezefy.Musica_em_Playlist
		where fk_musica = $1`, e.ID)
	if err != nil {
		return err
	}
	for _, b := range e.Playlists {
		_, err := r.db.Exec(`
		insert into deezefy.Musica_em_Playlist (fk_playlist, fk_musica)
		values($1,$2)`, b.Nome, e.ID)
		if err != nil {
			return err
		}
	}
	// update related album
	_, err = r.db.Exec(`
		delete from deezefy.Album_Contem_Musica
		where fk_musica = $1`, e.ID)
	if err != nil {
		return err
	}
	for _, b := range e.Albums {
		_, err := r.db.Exec(`
		insert into deezefy.Album_Contem_Musica (fk_album, fk_artista, fk_musica)
		values($1,$2,$3)`, b.ID, b.Artista.Usuario.Email, e.ID)
		if err != nil {
			return err
		}
	}
	// update related genero
	_, err = r.db.Exec(`
		delete from deezefy.Musica_Possui_Genero
		where fk_musica = $1`, e.ID)
	if err != nil {
		return err
	}
	for _, b := range e.Generos {
		_, err := r.db.Exec(`
		insert into deezefy.Musica_Possui_Genero (fk_genero, fk_musica)
		values($1,$2)`, b.Nome, e.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

// Search Musica
func (r *MusicaPSQL) Search(query string) ([]*entity.Musica, error) {
	stmt, err := r.db.Prepare(`
		select id from deezefy.Musica
		where nome like $1`)
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
	var Musicas []*entity.Musica
	for _, id := range ids {
		u, err := getMusica(id, r.db)
		if err != nil {
			return nil, err
		}
		Musicas = append(Musicas, u)
	}
	return Musicas, nil
}

// List Musicas
func (r *MusicaPSQL) List() ([]*entity.Musica, error) {
	stmt, err := r.db.Prepare(`
		select id from deezefy.Musica`)
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
	var Musicas []*entity.Musica
	for _, id := range ids {
		u, err := getMusica(id, r.db)
		if err != nil {
			return nil, err
		}
		Musicas = append(Musicas, u)
	}
	return Musicas, nil
}

// Delete an Musica
func (r *MusicaPSQL) Delete(id int) error {
	_, err := r.db.Exec(`
		delete from deezefy.Musica
		where id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
