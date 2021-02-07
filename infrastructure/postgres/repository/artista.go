package repository

import (
	"database/sql"
	"fmt"

	"github.com/yohanalexander/deezefy-music/entity"
)

// ArtistaPSQL postgres repo
type ArtistaPSQL struct {
	db *sql.DB
}

// NewArtistaPSQL create new repository
func NewArtistaPSQL(db *sql.DB) *ArtistaPSQL {
	return &ArtistaPSQL{
		db: db,
	}
}

// Create an artista
func (r *ArtistaPSQL) Create(e *entity.Artista) (string, error) {
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
		insert into deezefy.Artista (nome_artistico, biografia, ano_formacao, fk_usuario)
		values(?,?,?,?)`)
	if err != nil {
		return e.Usuario.Email, err
	}
	_, err = stmt.Exec(
		e.NomeArtistico,
		e.Biografia,
		e.AnoFormacao,
		e.Usuario.Email,
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

// Get an artista
func (r *ArtistaPSQL) Get(email string) (*entity.Artista, error) {
	return getArtista(email, r.db)
}

func getArtista(email string, db *sql.DB) (*entity.Artista, error) {
	stmt, err := db.Prepare(`select email, senha, data_nascimento, nome_artistico, biografia, ano_formacao from deezefy.Artista
		join deezefy.Usuario on(Artista.fk_usuario = Usuario.email)
		where email = ?`)
	if err != nil {
		return nil, err
	}
	var u entity.Artista
	rows, err := stmt.Query(email)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&u.Usuario.Email, &u.Usuario.Password, &u.Usuario.Birthday, &u.NomeArtistico, &u.Biografia, &u.AnoFormacao)
	}
	// select related ouvinte
	stmt, err = db.Prepare(`select email, senha, data_nascimento, primeiro_nome, sobrenome from deezefy.Ouvinte
	join deezefy.Usuario on(Usuario.email = Ouvinte.fk_usuario)
	join deezefy.Segue on(Ouvinte.fk_usuario = Segue.fk_ouvinte)
	where fk_artista = ?`)
	if err != nil {
		return nil, err
	}
	rows, err = stmt.Query(email)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		j := &entity.Ouvinte{}
		err = rows.Scan(&j.Usuario.Email, &j.Usuario.Password, &j.Usuario.Birthday, &j.PrimeiroNome, &j.Sobrenome)
		u.Seguidores = append(u.Seguidores, *j)
	}
	// select related musica
	stmt, err = db.Prepare(`select id, nome, duracao from deezefy.Musica
	join deezefy.Grava on(Grava.fk_musica = Musica.id)
	where fk_artista = ?`)
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
		u.Grava = append(u.Grava, *j)
	}
	// select related perfil
	stmt, err = db.Prepare(`select email, senha, data_nascimento, primeiro_nome, sobrenome, id, informacoes_relevantes from deezefy.Perfil
	join deezefy.Ouvinte on(Ouvinte.fk_usuario = Perfil.fk_ouvinte)
	join deezefy.Usuario on(Usuario.email = Ouvinte.fk_usuario)
	join deezefy.Artistas_Favoritos on(Artistas_Favoritos.fk_perfil = Perfil.id)
	where fk_artista = ?`)
	if err != nil {
		return nil, err
	}
	rows, err = stmt.Query(email)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		j := &entity.Perfil{}
		err = rows.Scan(&j.Ouvinte.Usuario.Email, &j.Ouvinte.Usuario.Password, &j.Ouvinte.Usuario.Birthday, &j.Ouvinte.PrimeiroNome, &j.Ouvinte.Sobrenome, &j.ID, &j.InformacoesRelevantes)
		u.Perfis = append(u.Perfis, *j)
	}
	// select related genero
	stmt, err = db.Prepare(`select nome, estilo from deezefy.Genero
	join deezefy.Artista_Possui_Genero on(Artista_Possui_Genero.fk_genero = Genero.nome)
	where fk_artista = ?`)
	if err != nil {
		return nil, err
	}
	rows, err = stmt.Query(email)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		j := &entity.Genero{}
		err = rows.Scan(&j.Nome, &j.Estilo)
		u.Generos = append(u.Generos, *j)
	}
	// select related album
	stmt, err = db.Prepare(`select id, titulo, ano_lancamento from deezefy.Album
	join deezefy.Artista_Possui_Genero on(Artista_Possui_Genero.fk_genero = Genero.nome)
	where fk_artista = ?`)
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

// Update an artista
func (r *ArtistaPSQL) Update(e *entity.Artista) error {
	_, err := r.db.Exec(`update deezefy.Artista set email = ?, password = ?, data_nascimento = ?, nome_artistico = ?, biografia = ?, ano_formacao = ?
	from deezefy.Artista join deezefy.Usuario on(Artista.fk_usuario = Usuario.email)
	where email = ?`, e.Usuario.Email, e.Usuario.Password, e.Usuario.Birthday, e.NomeArtistico, e.Biografia, e.AnoFormacao, e.Usuario.Email)
	if err != nil {
		return err
	}
	// update related ouvinte
	_, err = r.db.Exec(`delete from deezefy.Segue where fk_artista = ?`, e.Usuario.Email)
	if err != nil {
		return err
	}
	for _, b := range e.Seguidores {
		_, err := r.db.Exec(`insert into deezefy.Segue
		(fk_ouvinte, fk_artista) values(?,?)`, b.Usuario.Email, e.Usuario.Email)
		if err != nil {
			return err
		}
	}
	// update related musica
	_, err = r.db.Exec(`delete from deezefy.Grava where fk_artista = ?`, e.Usuario.Email)
	if err != nil {
		return err
	}
	for _, b := range e.Grava {
		_, err := r.db.Exec(`insert into deezefy.Grava
		(fk_musica, fk_artista) values(?,?)`, b.ID, e.Usuario.Email)
		if err != nil {
			return err
		}
	}
	// update related perfil
	_, err = r.db.Exec(`delete from deezefy.Artistas_Favoritos where fk_artista = ?`, e.Usuario.Email)
	if err != nil {
		return err
	}
	for _, b := range e.Perfis {
		_, err := r.db.Exec(`insert into deezefy.Artistas_Favoritos
		(fk_perfil, fk_ouvinte, fk_artista) values(?,?,?)`, b.ID, b.Ouvinte.Usuario.Email, e.Usuario.Email)
		if err != nil {
			return err
		}
	}
	// update related genero
	_, err = r.db.Exec(`delete from deezefy.Artista_Possui_Genero where fk_artista = ?`, e.Usuario.Email)
	if err != nil {
		return err
	}
	for _, b := range e.Generos {
		_, err := r.db.Exec(`insert into deezefy.Artista_Possui_Genero
		(fk_genero, fk_artista) values(?,?)`, b.Nome, e.Usuario.Email)
		if err != nil {
			return err
		}
	}
	return nil
}

// Search artista
func (r *ArtistaPSQL) Search(query string) ([]*entity.Artista, error) {
	stmt, err := r.db.Prepare(`select email from deezefy.Artista
	join deezefy.Usuario on(Artista.fk_usuario = Usuario.email) where email like ?`)
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
	var artistas []*entity.Artista
	for _, email := range emails {
		u, err := getArtista(email, r.db)
		if err != nil {
			return nil, err
		}
		artistas = append(artistas, u)
	}
	return artistas, nil
}

// List artistas
func (r *ArtistaPSQL) List() ([]*entity.Artista, error) {
	stmt, err := r.db.Prepare(`select email from deezefy.Artista
	join deezefy.Usuario on(Artista.fk_usuario = Usuario.email)`)
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
	var artistas []*entity.Artista
	for _, email := range emails {
		u, err := getArtista(email, r.db)
		if err != nil {
			return nil, err
		}
		artistas = append(artistas, u)
	}
	return artistas, nil
}

// Delete an artista
func (r *ArtistaPSQL) Delete(email string) error {
	_, err := r.db.Exec(`delete from deezefy.Artista where fk_usuario = ?`, email)
	if err != nil {
		return err
	}
	return nil
}
