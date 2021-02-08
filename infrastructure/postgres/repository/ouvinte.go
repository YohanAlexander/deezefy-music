package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/yohanalexander/deezefy-music/entity"
)

// OuvintePSQL postgres repo
type OuvintePSQL struct {
	db *sql.DB
}

// NewOuvintePSQL create new repository
func NewOuvintePSQL(db *sql.DB) *OuvintePSQL {
	return &OuvintePSQL{
		db: db,
	}
}

// Create an Ouvinte
func (r *OuvintePSQL) Create(e *entity.Ouvinte) (string, error) {
	stmt, err := r.db.Prepare(`
		insert into deezefy.Usuario (email, senha, data_nascimento)
		values(?,?,?)`)
	if err != nil {
		return e.Usuario.Email, err
	}
	_, err = stmt.Exec(
		e.Usuario.Email,
		e.Usuario.Password,
		e.Usuario.Birthday,
	)
	if err != nil {
		return e.Usuario.Email, err
	}
	stmt, err = r.db.Prepare(`
		insert into deezefy.Ouvinte (primeiro_nome, sobrenome)
		values(?,?)`)
	if err != nil {
		return e.Usuario.Email, err
	}
	_, err = stmt.Exec(
		e.PrimeiroNome,
		e.Sobrenome,
	)
	if err != nil {
		return e.Usuario.Email, err
	}
	err = stmt.Close()
	if err != nil {
		return e.Usuario.Email, err
	}
	return e.Usuario.Email, nil
}

// Get an Ouvinte
func (r *OuvintePSQL) Get(email string) (*entity.Ouvinte, error) {
	return getOuvinte(email, r.db)
}

func getOuvinte(email string, db *sql.DB) (*entity.Ouvinte, error) {
	stmt, err := db.Prepare(`select email, senha, data_nascimento, primeiro_nome, sobrenome from deezefy.Ouvinte
		join deezefy.Usuario on(Ouvinte.fk_usuario = Usuario.email)
		where email = ?`)
	if err != nil {
		return nil, err
	}
	var u entity.Ouvinte
	rows, err := stmt.Query(email)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&u.Usuario.Email, &u.Usuario.Password, &u.Usuario.Birthday, &u.PrimeiroNome, &u.Sobrenome)
	}
	// select related cria
	stmt, err = db.Prepare(`select nome, status, data_criacao from deezefy.Playlist
	join deezefy.Cria on(Playlist.nome = Cria.fk_playlist) where fk_usuario = ?`)
	if err != nil {
		return nil, err
	}
	rows, err = stmt.Query(email)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		j := &entity.Playlist{}
		err = rows.Scan(&j.Nome, &j.Status, &j.DataCriacao)
		u.Cria = append(u.Cria, *j)
	}
	// select related telefone
	stmt, err = db.Prepare(`select telefone from deezefy.Telefone
	join deezefy.Ouvinte on(Ouvinte.fk_usuario = Telefone.fk_ouvinte)
	where fk_ouvinte = ?`)
	if err != nil {
		return nil, err
	}
	rows, err = stmt.Query(email)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var j string
		err = rows.Scan(&j)
		u.Telefones = append(u.Telefones, j)
	}
	// select related artista
	stmt, err = db.Prepare(`select email, senha, data_nascimento, nome_artistico, biografia, ano_formacao from deezefy.Artista
	join deezefy.Usuario on(Usuario.email = Artista.fk_usuario)
	join deezefy.Segue on(Ouvinte.fk_usuario = Segue.fk_artista)
	where fk_ouvinte = ?`)
	if err != nil {
		return nil, err
	}
	rows, err = stmt.Query(email)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		j := &entity.Artista{}
		err = rows.Scan(&j.Usuario.Email, &j.Usuario.Password, &j.Usuario.Birthday, &j.NomeArtistico, &j.Biografia, &j.AnoFormacao)
		u.Seguindo = append(u.Seguindo, *j)
	}
	// select related musica
	stmt, err = db.Prepare(`select id, nome, duracao from deezefy.Musica
	join deezefy.Curte on(Curte.fk_musica = Musica.id)
	where fk_ouvinte = ?`)
	if err != nil {
		return nil, err
	}
	rows, err = stmt.Query(email)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		j := &entity.Musica{}
		err = rows.Scan(&j.ID, &j.Nome, &j.Duracao)
		u.Curtidas = append(u.Curtidas, *j)
	}
	// select related playlist
	stmt, err = db.Prepare(`select nome, status, data_criacao from deezefy.Playlist
	join deezefy.Ouvinte_Salva_Playlist on(Playlist.nome = Ouvinte_Salva_Playlist.fk_playlist) where fk_ouvinte = ?`)
	if err != nil {
		return nil, err
	}
	rows, err = stmt.Query(email)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		j := &entity.Playlist{}
		err = rows.Scan(&j.Nome, &j.Status, &j.DataCriacao)
		u.Playlists = append(u.Playlists, *j)
	}
	// select related album
	stmt, err = db.Prepare(`select id, titulo, ano_lancamento from deezefy.Album
	join deezefy.Ouvinte_Salva_Album on(Ouvinte_Salva_Album.fk_album = Album.id)
	where fk_ouvinte = ?`)
	if err != nil {
		return nil, err
	}
	rows, err = stmt.Query(email)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		j := &entity.Album{}
		err = rows.Scan(&j.ID, &j.Titulo, &j.AnoLancamento)
		u.Albums = append(u.Albums, *j)
	}
	return &u, nil
}

// Update an Ouvinte
func (r *OuvintePSQL) Update(e *entity.Ouvinte) error {
	_, err := r.db.Exec(`update deezefy.Ouvinte set email = ?, password = ?, data_nascimento = ?, primeiro_nome = ?, sobrenome = ?
	from deezefy.Ouvinte join deezefy.Usuario on(Ouvinte.fk_usuario = Usuario.email)
	where email = ?`, e.Usuario.Email, e.Usuario.Password, e.Usuario.Birthday, e.PrimeiroNome, e.Sobrenome, e.Usuario.Email)
	if err != nil {
		return err
	}
	// update related cria
	_, err = r.db.Exec(`delete from deezefy.Cria where fk_usuario = ?`, e.Usuario.Email)
	if err != nil {
		return err
	}
	for _, b := range e.Cria {
		_, err := r.db.Exec(`insert into deezefy.Cria
		(data_criacao, fk_playlist, fk_usuario) values(?,?,?)`, time.Now().Format("2006-01-02"), b.Nome, e.Usuario.Email)
		if err != nil {
			return err
		}
	}
	// update related artista
	_, err = r.db.Exec(`delete from deezefy.Artista where fk_ouvinte = ?`, e.Usuario.Email)
	if err != nil {
		return err
	}
	for _, b := range e.Seguindo {
		_, err := r.db.Exec(`insert into deezefy.Segue
		(fk_artista, fk_ouvinte) values(?,?)`, b.Usuario.Email, e.Usuario.Email)
		if err != nil {
			return err
		}
	}
	// update related musica
	_, err = r.db.Exec(`delete from deezefy.Curte where fk_ouvinte = ?`, e.Usuario.Email)
	if err != nil {
		return err
	}
	for _, b := range e.Curtidas {
		_, err := r.db.Exec(`insert into deezefy.Curte
		(fk_musica, fk_ouvinte) values(?,?)`, b.ID, e.Usuario.Email)
		if err != nil {
			return err
		}
	}
	// update related playlist
	_, err = r.db.Exec(`delete from deezefy.Ouvinte_Salva_Playlist where fk_ouvinte = ?`, e.Usuario.Email)
	if err != nil {
		return err
	}
	for _, b := range e.Playlists {
		_, err := r.db.Exec(`insert into deezefy.Ouvinte_Salva_Playlist
		(fk_playlist, fk_ouvinte) values(?,?)`, b.Nome, e.Usuario.Email)
		if err != nil {
			return err
		}
	}
	// update related album
	_, err = r.db.Exec(`delete from deezefy.Ouvinte_Salva_Album where fk_ouvinte = ?`, e.Usuario.Email)
	if err != nil {
		return err
	}
	for _, b := range e.Albums {
		_, err := r.db.Exec(`insert into deezefy.Ouvinte_Salva_Album
		(fk_album, fk_artista, fk_ouvinte) values(?,?,?)`, b.ID, b.Artista.Usuario.Email, e.Usuario.Email)
		if err != nil {
			return err
		}
	}
	return nil
}

// Search Ouvinte
func (r *OuvintePSQL) Search(query string) ([]*entity.Ouvinte, error) {
	stmt, err := r.db.Prepare(`select email from deezefy.Ouvinte
	join deezefy.Usuario on(Ouvinte.fk_usuario = Usuario.email) where email like ?`)
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
	var Ouvintes []*entity.Ouvinte
	for _, email := range emails {
		u, err := getOuvinte(email, r.db)
		if err != nil {
			return nil, err
		}
		Ouvintes = append(Ouvintes, u)
	}
	return Ouvintes, nil
}

// List Ouvintes
func (r *OuvintePSQL) List() ([]*entity.Ouvinte, error) {
	stmt, err := r.db.Prepare(`select email from deezefy.Ouvinte
	join deezefy.Usuario on(Ouvinte.fk_usuario = Usuario.email)`)
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
	var Ouvintes []*entity.Ouvinte
	for _, email := range emails {
		u, err := getOuvinte(email, r.db)
		if err != nil {
			return nil, err
		}
		Ouvintes = append(Ouvintes, u)
	}
	return Ouvintes, nil
}

// Delete an Ouvinte
func (r *OuvintePSQL) Delete(email string) error {
	_, err := r.db.Exec(`delete from deezefy.Ouvinte where fk_usuario = ?`, email)
	if err != nil {
		return err
	}
	return nil
}
