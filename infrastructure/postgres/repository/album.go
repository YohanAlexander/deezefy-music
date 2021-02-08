package repository

import (
	"database/sql"
	"fmt"

	"github.com/yohanalexander/deezefy-music/entity"
)

// AlbumPSQL postgres repo
type AlbumPSQL struct {
	db *sql.DB
}

// NewAlbumPSQL create new repository
func NewAlbumPSQL(db *sql.DB) *AlbumPSQL {
	return &AlbumPSQL{
		db: db,
	}
}

// Create an Album
func (r *AlbumPSQL) Create(e *entity.Album) (int, error) {
	stmt, err := r.db.Prepare(`
		insert into deezefy.Album (id, titulo, ano_lancamento)
		values($1,$2,$3)`)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.Titulo,
		e.AnoLancamento,
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

// Get an Album
func (r *AlbumPSQL) Get(id int) (*entity.Album, error) {
	return getAlbum(id, r.db)
}

func getAlbum(id int, db *sql.DB) (*entity.Album, error) {
	stmt, err := db.Prepare(`
		select email, senha, data_nascimento, nome_artistico, biografia, ano_formacao, id, titulo, ano_lancamento from deezefy.Album
		join deezefy.Artista on(Artista.fk_usuario = Album_fk_artista)
		join deezefy.Usuario on(Usuario.email = Artista.fk_usuario)
		where id = $1`)
	if err != nil {
		return nil, err
	}
	var u entity.Album
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&u.Artista.Usuario.Email, &u.Artista.Usuario.Password, &u.Artista.Usuario.Birthday, &u.Artista.NomeArtistico, &u.Artista.Biografia, &u.Artista.AnoFormacao, &u.ID, &u.Titulo, &u.AnoLancamento)
	}
	// select related ouvinte
	stmt, err = db.Prepare(`
		select email, senha, data_nascimento, primeiro_nome, sobrenome from deezefy.Album
		join deezefy.Ouvinte_Salva_Album on(Ouvinte_Salva_Album.fk_album = Album.id)
		join deezefy.Ouvinte on(Ouvinte.fk_usuario = Ouvinte_Salva_Album.fk_ouvinte)
		join deezefy.Usuario on(Usuario.email = Ouvinte.fk_usuario)
		where Album.id = $1`)
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
		u.Salvou = append(u.Salvou, *j)
	}
	// select related musica
	stmt, err = db.Prepare(`
		select id, nome, duracao from deezefy.Musica
		join deezefy.Album_Contem_Musica on(Album_Contem_Musica.fk_musica = Musica.id)
		where Album.id = $1`)
	if err != nil {
		return nil, err
	}
	rows, err = stmt.Query(id)
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

// Update an Album
func (r *AlbumPSQL) Update(e *entity.Album) error {
	_, err := r.db.Exec(`
		update deezefy.Album set id = $1, titulo = $2, ano_lancamento = $3
		where id = $4`, e.ID, e.Titulo, e.AnoLancamento, e.ID)
	if err != nil {
		return err
	}
	// update related ouvinte
	_, err = r.db.Exec(`
		delete from deezefy.Ouvinte_Salva_Album
		where fk_album = $1`, e.ID)
	if err != nil {
		return err
	}
	for _, b := range e.Salvou {
		_, err := r.db.Exec(`
		insert into deezefy.Ouvinte_Salva_Album (fk_ouvinte, fk_artista, fk_album)
		values($1,$2,$3)`, b.Usuario.Email, e.Artista.Usuario.Email, e.ID)
		if err != nil {
			return err
		}
	}
	// update related musica
	_, err = r.db.Exec(`
		delete from deezefy.Album_Contem_Musica
		where fk_album = $1`, e.ID)
	if err != nil {
		return err
	}
	for _, b := range e.Musicas {
		_, err := r.db.Exec(`
		insert into deezefy.Album_Contem_Musica (fk_musica, fk_artista, fk_album)
		values($1,$2,$3)`, b.ID, e.Artista.Usuario.Email, e.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

// Search Album
func (r *AlbumPSQL) Search(query string) ([]*entity.Album, error) {
	stmt, err := r.db.Prepare(`
		select id from deezefy.Album
		where titulo like $1`)
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
	var Albums []*entity.Album
	for _, id := range ids {
		u, err := getAlbum(id, r.db)
		if err != nil {
			return nil, err
		}
		Albums = append(Albums, u)
	}
	return Albums, nil
}

// List Albums
func (r *AlbumPSQL) List() ([]*entity.Album, error) {
	stmt, err := r.db.Prepare(`
		select id from deezefy.Album`)
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
	var Albums []*entity.Album
	for _, id := range ids {
		u, err := getAlbum(id, r.db)
		if err != nil {
			return nil, err
		}
		Albums = append(Albums, u)
	}
	return Albums, nil
}

// Delete an Album
func (r *AlbumPSQL) Delete(id int) error {
	_, err := r.db.Exec(`
		delete from deezefy.Album
		where id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
