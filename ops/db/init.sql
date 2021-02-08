-- -----------------------------------------------------
-- Schema deezefy
-- -----------------------------------------------------
-- \! echo "Schema deezefy";
CREATE SCHEMA IF NOT EXISTS deezefy;
-- \! echo "Use Schema deezefy";
SET search_path TO deezefy;
-- -----------------------------------------------------
-- TYPE deezefy.status
-- -----------------------------------------------------
-- \! echo "TYPE deezefy.status";
CREATE TYPE deezefy.status AS ENUM('ativo', 'inativo');
-- -----------------------------------------------------
-- TYPE deezefy.estilo
-- -----------------------------------------------------
-- \! echo "TYPE deezefy.estilo";
CREATE TYPE deezefy.estilo AS ENUM(
  'blues',
  'rock',
  'mpb',
  'samba',
  'sertanejo',
  'jazz',
  'classica'
);
-- -----------------------------------------------------
-- Table deezefy.Usuario
-- -----------------------------------------------------
-- \! echo "Table deezefy.Usuario";
CREATE TABLE IF NOT EXISTS deezefy.Usuario (
  email VARCHAR(45) NOT NULL,
  senha VARCHAR(100) NOT NULL,
  data_nascimento DATE NULL,
  PRIMARY KEY (email),
  UNIQUE (email)
);
-- -----------------------------------------------------
-- Table deezefy.Evento
-- -----------------------------------------------------
-- \! echo "Table deezefy.Evento";
CREATE TABLE IF NOT EXISTS deezefy.Evento (
  id SERIAL NOT NULL,
  nome VARCHAR(45) NOT NULL,
  fk_usuario VARCHAR(45) NOT NULL,
  PRIMARY KEY (id, fk_usuario),
  CONSTRAINT fk_evento FOREIGN KEY (fk_usuario) REFERENCES deezefy.Usuario (email) ON DELETE CASCADE ON UPDATE CASCADE
);
-- -----------------------------------------------------
-- Table deezefy.Ouvinte
-- -----------------------------------------------------
-- \! echo "Table deezefy.Ouvinte";
CREATE TABLE IF NOT EXISTS deezefy.Ouvinte (
  primeiro_nome VARCHAR(45) NOT NULL,
  sobrenome VARCHAR(45) NOT NULL,
  fk_usuario VARCHAR(45) NOT NULL,
  PRIMARY KEY (fk_usuario),
  CONSTRAINT fk_ouvinte FOREIGN KEY (fk_usuario) REFERENCES deezefy.Usuario (email) ON DELETE CASCADE ON UPDATE CASCADE
);
-- -----------------------------------------------------
-- Table deezefy.Artista
-- -----------------------------------------------------
-- \! echo "Table deezefy.Artista";
CREATE TABLE IF NOT EXISTS deezefy.Artista (
  nome_artistico VARCHAR(45) NOT NULL,
  biografia VARCHAR(45) NULL,
  ano_formacao INT NULL,
  fk_usuario VARCHAR(45) NOT NULL,
  PRIMARY KEY (fk_usuario),
  CONSTRAINT fk_usuario FOREIGN KEY (fk_usuario) REFERENCES deezefy.Usuario (email) ON DELETE CASCADE ON UPDATE CASCADE
);
-- -----------------------------------------------------
-- Table deezefy.Musica
-- -----------------------------------------------------
-- \! echo "Table deezefy.Musica";
CREATE TABLE IF NOT EXISTS deezefy.Musica (
  id SERIAL NOT NULL,
  nome VARCHAR(45) NOT NULL,
  duracao INT NOT NULL,
  PRIMARY KEY (id)
);
-- -----------------------------------------------------
-- Table deezefy.Perfil
-- -----------------------------------------------------
-- \! echo "Table deezefy.Perfil";
CREATE TABLE IF NOT EXISTS deezefy.Perfil (
  id SERIAL NOT NULL,
  informacoes_relevantes VARCHAR(45) NOT NULL,
  fk_ouvinte VARCHAR(45) NOT NULL,
  PRIMARY KEY (id, fk_ouvinte),
  CONSTRAINT fk_ouvinte FOREIGN KEY (fk_ouvinte) REFERENCES deezefy.Ouvinte (fk_usuario) ON DELETE CASCADE ON UPDATE CASCADE
);
-- -----------------------------------------------------
-- Table deezefy.Album
-- -----------------------------------------------------
-- \! echo "Table deezefy.Album";
CREATE TABLE IF NOT EXISTS deezefy.Album (
  id SERIAL NOT NULL,
  titulo VARCHAR(45) NOT NULL,
  ano_lancamento INT NULL,
  fk_artista VARCHAR(45) NULL,
  PRIMARY KEY (id, fk_artista),
  CONSTRAINT fk_artista FOREIGN KEY (fk_artista) REFERENCES deezefy.Artista (fk_usuario) ON DELETE CASCADE ON UPDATE CASCADE
);
-- -----------------------------------------------------
-- Table deezefy.Playlist
-- -----------------------------------------------------
-- \! echo "Table deezefy.Playlist";
CREATE TABLE IF NOT EXISTS deezefy.Playlist (
  nome VARCHAR(45) NOT NULL,
  status deezefy.status NOT NULL,
  PRIMARY KEY (nome)
);
-- -----------------------------------------------------
-- Table deezefy.Local
-- -----------------------------------------------------
-- \! echo "Table deezefy.Local";
CREATE TABLE IF NOT EXISTS deezefy.Local (
  id SERIAL NOT NULL,
  cidade VARCHAR(45) NOT NULL,
  pais VARCHAR(45) NOT NULL,
  PRIMARY KEY (id)
);
-- -----------------------------------------------------
-- Table deezefy.Cria
-- -----------------------------------------------------
-- \! echo "Table deezefy.Cria";
CREATE TABLE IF NOT EXISTS deezefy.Cria (
  data_criacao DATE NOT NULL,
  fk_playlist VARCHAR(45) NOT NULL,
  fk_usuario VARCHAR(45) NOT NULL,
  PRIMARY KEY (fk_playlist, fk_usuario),
  CONSTRAINT fk_playlist FOREIGN KEY (fk_playlist) REFERENCES deezefy.Playlist (nome) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT fk_usuario FOREIGN KEY (fk_usuario) REFERENCES deezefy.Usuario (email) ON DELETE CASCADE ON UPDATE CASCADE
);
-- -----------------------------------------------------
-- Table deezefy.Ocorre
-- -----------------------------------------------------
-- \! echo "Table deezefy.Ocorre";
CREATE TABLE IF NOT EXISTS deezefy.Ocorre (
  data DATE NOT NULL,
  fk_artista VARCHAR(45) NULL,
  fk_local INT NULL,
  fk_evento INT NULL,
  fk_usuario VARCHAR(45) NULL,
  PRIMARY KEY (fk_artista, fk_local, fk_evento, fk_usuario),
  CONSTRAINT fk_local FOREIGN KEY (fk_local) REFERENCES deezefy.Local (id) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT fk_artista FOREIGN KEY (fk_artista) REFERENCES deezefy.Artista (fk_usuario) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT fk_evento FOREIGN KEY (fk_evento, fk_usuario) REFERENCES deezefy.Evento (id, fk_usuario) ON DELETE CASCADE ON UPDATE CASCADE
);
-- -----------------------------------------------------
-- Table deezefy.Genero
-- -----------------------------------------------------
-- \! echo "Table deezefy.Genero";
CREATE TABLE IF NOT EXISTS deezefy.Genero (
  nome VARCHAR(45) NOT NULL,
  estilo deezefy.estilo NOT NULL,
  PRIMARY KEY (nome)
);
-- -----------------------------------------------------
-- Table deezefy.Generos_Favoritos
-- -----------------------------------------------------
-- \! echo "Table deezefy.Generos_Favoritos";
CREATE TABLE IF NOT EXISTS deezefy.Generos_Favoritos (
  fk_genero VARCHAR(45) NULL,
  fk_perfil INT NULL,
  fk_ouvinte VARCHAR(45) NULL,
  PRIMARY KEY (fk_genero, fk_perfil, fk_ouvinte),
  CONSTRAINT fk_genero FOREIGN KEY (fk_genero) REFERENCES deezefy.Genero (nome) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT fk_perfil FOREIGN KEY (fk_perfil, fk_ouvinte) REFERENCES deezefy.Perfil (id, fk_ouvinte) ON DELETE CASCADE ON UPDATE CASCADE
);
-- -----------------------------------------------------
-- Table deezefy.Musica_em_Playlist
-- -----------------------------------------------------
-- \! echo "Table deezefy.Musica_em_Playlist";
CREATE TABLE IF NOT EXISTS deezefy.Musica_em_Playlist (
  fk_musica INT NOT NULL,
  fk_playlist VARCHAR(45) NULL,
  PRIMARY KEY (fk_musica, fk_playlist),
  CONSTRAINT fk_musica FOREIGN KEY (fk_musica) REFERENCES deezefy.Musica (id) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT fk_playlist FOREIGN KEY (fk_playlist) REFERENCES deezefy.Playlist (nome) ON DELETE CASCADE ON UPDATE CASCADE
);
-- -----------------------------------------------------
-- Table deezefy.Musica_Possui_Generos
-- -----------------------------------------------------
-- \! echo "Table deezefy.Musica_Possui_Generos";
CREATE TABLE IF NOT EXISTS deezefy.Musica_Possui_Generos (
  fk_musica INT NULL,
  fk_genero VARCHAR(45) NULL,
  PRIMARY KEY (fk_musica, fk_genero),
  CONSTRAINT fk_musica FOREIGN KEY (fk_musica) REFERENCES deezefy.Musica (id) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT fk_genero FOREIGN KEY (fk_genero) REFERENCES deezefy.Genero (nome) ON DELETE CASCADE ON UPDATE CASCADE
);
-- -----------------------------------------------------
-- Table deezefy.Segue
-- -----------------------------------------------------
-- \! echo "Table deezefy.Segue";
CREATE TABLE IF NOT EXISTS deezefy.Segue (
  fk_artista VARCHAR(45) NULL,
  fk_ouvinte VARCHAR(45) NULL,
  PRIMARY KEY (fk_artista, fk_ouvinte),
  CONSTRAINT fk_artista FOREIGN KEY (fk_artista) REFERENCES deezefy.Artista (fk_usuario) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT fk_ouvinte FOREIGN KEY (fk_ouvinte) REFERENCES deezefy.Ouvinte (fk_usuario) ON DELETE CASCADE ON UPDATE CASCADE
);
-- -----------------------------------------------------
-- Table deezefy.Ouvinte_Salva_Playlist
-- -----------------------------------------------------
-- \! echo "Table deezefy.Ouvinte_Salva_Playlist";
CREATE TABLE IF NOT EXISTS deezefy.Ouvinte_Salva_Playlist (
  fk_ouvinte VARCHAR(45) NULL,
  fk_playlist VARCHAR(45) NULL,
  PRIMARY KEY (fk_ouvinte, fk_playlist),
  CONSTRAINT fk_ouvinte FOREIGN KEY (fk_ouvinte) REFERENCES deezefy.Ouvinte (fk_usuario) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT fk_playlist FOREIGN KEY (fk_playlist) REFERENCES deezefy.Playlist (nome) ON DELETE CASCADE ON UPDATE CASCADE
);
-- -----------------------------------------------------
-- Table deezefy.Artistas_Favoritos
-- -----------------------------------------------------
-- \! echo "Table deezefy.Artistas_Favoritos";
CREATE TABLE IF NOT EXISTS deezefy.Artistas_Favoritos (
  fk_artista VARCHAR(45) NULL,
  fk_perfil INT NULL,
  fk_ouvinte VARCHAR(45) NULL,
  PRIMARY KEY (fk_artista, fk_perfil, fk_ouvinte),
  CONSTRAINT fk_artista FOREIGN KEY (fk_artista) REFERENCES deezefy.Artista (fk_usuario) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT fk_perfil FOREIGN KEY (fk_perfil, fk_ouvinte) REFERENCES deezefy.Perfil (id, fk_ouvinte) ON DELETE CASCADE ON UPDATE CASCADE
);
-- -----------------------------------------------------
-- Table deezefy.Telefone
-- -----------------------------------------------------
-- \! echo "Table deezefy.Telefone";
CREATE TABLE IF NOT EXISTS deezefy.Telefone (
  telefone VARCHAR(45) NOT NULL,
  fk_ouvinte VARCHAR(45) NOT NULL,
  PRIMARY KEY (fk_ouvinte),
  CONSTRAINT fk_ouvinte FOREIGN KEY (fk_ouvinte) REFERENCES deezefy.Ouvinte (fk_usuario) ON DELETE CASCADE ON UPDATE CASCADE
);
-- -----------------------------------------------------
-- Table deezefy.Curte
-- -----------------------------------------------------
-- \! echo "Table deezefy.Curte";
CREATE TABLE IF NOT EXISTS deezefy.Curte (
  fk_musica INT NULL,
  fk_ouvinte VARCHAR(45) NULL,
  PRIMARY KEY (fk_musica, fk_ouvinte),
  CONSTRAINT fk_musica FOREIGN KEY (fk_musica) REFERENCES deezefy.Musica (id) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT fk_ouvinte FOREIGN KEY (fk_ouvinte) REFERENCES deezefy.Ouvinte (fk_usuario) ON DELETE CASCADE ON UPDATE CASCADE
);
-- -----------------------------------------------------
-- Table deezefy.Album_Contem_Musica
-- -----------------------------------------------------
-- \! echo "Table deezefy.Album_Contem_Musica";
CREATE TABLE IF NOT EXISTS deezefy.Album_Contem_Musica (
  fk_album INT NOT NULL,
  fk_artista VARCHAR(45) NOT NULL,
  fk_musica INT NOT NULL,
  PRIMARY KEY (fk_album, fk_artista, fk_musica),
  CONSTRAINT fk_album FOREIGN KEY (fk_album, fk_artista) REFERENCES deezefy.Album (id, fk_artista) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT fk_musica FOREIGN KEY (fk_musica) REFERENCES deezefy.Musica (id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- -----------------------------------------------------
-- Table deezefy.Ouvinte_Salva_Album
-- -----------------------------------------------------
-- \! echo "Table deezefy.Ouvinte_Salva_Album";
CREATE TABLE IF NOT EXISTS deezefy.Ouvinte_Salva_Album (
  fk_ouvinte VARCHAR(45) NOT NULL,
  fk_album INT NOT NULL,
  fk_artista VARCHAR(45) NOT NULL,
  PRIMARY KEY (fk_ouvinte, fk_album, fk_artista),
  CONSTRAINT fk_ouvinte FOREIGN KEY (fk_ouvinte) REFERENCES deezefy.Ouvinte (fk_usuario) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT fk_album FOREIGN KEY (fk_album, fk_artista) REFERENCES deezefy.Album (id, fk_artista) ON DELETE CASCADE ON UPDATE CASCADE
);
-- -----------------------------------------------------
-- Table deezefy.Artista_Possui_Genero
-- -----------------------------------------------------
-- \! echo "Table deezefy.Artista_Possui_Genero";
CREATE TABLE IF NOT EXISTS deezefy.Artista_Possui_Genero (
  fk_artista VARCHAR(45) NULL,
  fk_genero VARCHAR(45) NULL,
  PRIMARY KEY (fk_artista, fk_genero),
  CONSTRAINT fk_artista FOREIGN KEY (fk_artista) REFERENCES deezefy.Artista (fk_usuario) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT fk_genero FOREIGN KEY (fk_genero) REFERENCES deezefy.Genero (nome) ON DELETE CASCADE ON UPDATE CASCADE
);
-- -----------------------------------------------------
-- Table deezefy.Grava
-- -----------------------------------------------------
-- \! echo "Table deezefy.Grava";
CREATE TABLE IF NOT EXISTS deezefy.Grava (
  fk_musica INT NULL,
  fk_artista VARCHAR(45) NOT NULL,
  PRIMARY KEY (fk_musica, fk_artista),
  CONSTRAINT fk_musica FOREIGN KEY (fk_musica) REFERENCES deezefy.Musica (id) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT fk_artista FOREIGN KEY (fk_artista) REFERENCES deezefy.Artista (fk_usuario) ON DELETE CASCADE ON UPDATE CASCADE
);
-- -----------------------------------------------------
-- Schema deezefy
-- -----------------------------------------------------