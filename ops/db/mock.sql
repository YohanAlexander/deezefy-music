-- -----------------------------------------------------
-- Schema deezefy
-- -----------------------------------------------------
-- \! echo 'Use Schema deezefy';
SET search_path TO deezefy;
-- -----------------------------------------------------
-- Table deezefy.Usuario
-- -----------------------------------------------------
-- \! echo 'Table deezefy.Usuario';
INSERT INTO deezefy.Usuario (email,senha,data_nascimento)
VALUES ('ouvinte1@email.com','senha1234','1990-05-12'),
    ('ouvinte2@email.com','senha1234','1998-05-27'),
    ('u2@email.com','senha1234','1976-01-01'),
    ('nirvana@email.com','senha1234','1987-01-01'),
    ('oasis@email.com','senha1234','1991-01-01');
-- -----------------------------------------------------
-- Table deezefy.Artista
-- -----------------------------------------------------
-- \! echo 'Table deezefy.Artista';
INSERT INTO deezefy.Artista (nome_artistico,biografia,ano_formacao,fk_usuario)
VALUES ('U2','Banda de rock irlandesa','1976','u2@email.com'),
    ('Nirvana','Banda de rock americana','1987','nirvana@email.com'),
    ('Oasis','Banda de rock inglesa','1991','oasis@email.com');
-- -----------------------------------------------------
-- Table deezefy.Ouvinte
-- -----------------------------------------------------
-- \! echo 'Table deezefy.Ouvinte';
INSERT INTO deezefy.Ouvinte (primeiro_nome,sobrenome,fk_usuario)
VALUES ('Primeiro','Ouvinte','ouvinte1@email.com'),
    ('Segundo','Ouvinte','ouvinte2@email.com');
-- -----------------------------------------------------
-- Table deezefy.Telefone
-- -----------------------------------------------------
-- \! echo 'Table deezefy.Telefone';
INSERT INTO deezefy.Telefone (telefone,fk_ouvinte)
VALUES ('+5579999999999','ouvinte1@email.com'),
    ('+5579999999999','ouvinte2@email.com');
-- -----------------------------------------------------
-- Table deezefy.Musica
-- -----------------------------------------------------
-- \! echo 'Table deezefy.Musica';
INSERT INTO deezefy.Musica (id,nome,duracao)
VALUES (1,'Hello',322),
    (2,'Roll With It',400),
    (3,'Wonderwall',419),
    (4,'Rock and Roll Star',523),
    (5,'Shakermaker',508),
    (6,'Live Forever',437),
    (7,'Where the streets have no name',538),
    (8,'With or without you',456),
    (9,'Sunday blood sunday',440),
    (10,'Seconds',311),
    (11,'New years day',536),
    (12,'Smells like teen spirit',502),
    (13,'Come as you are',339);
-- -----------------------------------------------------
-- Table deezefy.Album
-- -----------------------------------------------------
-- \! echo 'Table deezefy.Album';
INSERT INTO deezefy.Album (id,titulo,ano_lancamento,fk_artista)
VALUES (1,'Nevermind',1991,'nirvana@email.com'),
    (2,'The Joshua Tree',1987,'u2@email.com'),
    (3,'War',1983,'u2@email.com'),
    (4,'Definitely Maybe',1994,'oasis@email.com'),
    (5,'Morning Glory',1995,'oasis@email.com');
-- -----------------------------------------------------
-- Table deezefy.Playlist
-- -----------------------------------------------------
-- \! echo 'Table deezefy.Playlist';
INSERT INTO deezefy.Playlist (nome,status)
VALUES ('80s Rock','ativo'),
    ('Grunge','ativo');
-- -----------------------------------------------------
-- Table deezefy.Genero
-- -----------------------------------------------------
INSERT INTO deezefy.Genero (nome,estilo)
VALUES ('80s Rock','rock'),
    ('Grunge','rock');
-- -----------------------------------------------------
-- Table deezefy.Evento
-- -----------------------------------------------------
-- \! echo 'Table deezefy.Evento';
INSERT INTO deezefy.Evento (id,nome,fk_usuario)
VALUES (1,'Festival Woodstock','oasis@email.com'),
    (2,'Rock in Rio','nirvana@email.com');
-- -----------------------------------------------------
-- Table deezefy.Local
-- -----------------------------------------------------
-- \! echo 'Table deezefy.Local';
INSERT INTO deezefy.Local (id,pais,cidade)
VALUES (1,'EUA','New York'),
    (2,'Brazil','Rio');
-- -----------------------------------------------------
-- Table deezefy.Perfil
-- -----------------------------------------------------
-- \! echo 'Table deezefy.Perfil';
INSERT INTO deezefy.Perfil (id,informacoes_relevantes,fk_ouvinte)
VALUES (1,'Adora Rock','ouvinte1@email.com'),
    (2,'Fã do nirvana','ouvinte2@email.com');
-- -----------------------------------------------------
-- Table deezefy.Segue
-- -----------------------------------------------------
-- \! echo 'Table deezefy.Segue';
INSERT INTO deezefy.Segue (fk_ouvinte,fk_artista)
VALUES ('ouvinte2@email.com','nirvana@email.com'),
    ('ouvinte2@email.com','oasis@email.com'),
    ('ouvinte1@email.com','u2@email.com'),
    ('ouvinte1@email.com','oasis@email.com');
-- -----------------------------------------------------
-- Table deezefy.Curte
-- -----------------------------------------------------
-- \! echo 'Table deezefy.Curte';
INSERT INTO deezefy.Curte (fk_ouvinte,fk_musica)
VALUES ('ouvinte2@email.com',1),
    ('ouvinte2@email.com',4),
    ('ouvinte2@email.com',6),
    ('ouvinte2@email.com',7),
    ('ouvinte2@email.com',8),
    ('ouvinte2@email.com',9),
    ('ouvinte2@email.com',10),
    ('ouvinte1@email.com',3),
    ('ouvinte1@email.com',13),
    ('ouvinte1@email.com',2),
    ('ouvinte1@email.com',5),
    ('ouvinte1@email.com',11),
    ('ouvinte1@email.com',12);
-- -----------------------------------------------------
-- Table deezefy.Ouvinte_Salva_Playlist
-- -----------------------------------------------------
-- \! echo 'Table deezefy.Ouvinte_Salva_Playlist';
INSERT INTO deezefy.Ouvinte_Salva_Playlist (fk_ouvinte,fk_playlist)
VALUES ('ouvinte2@email.com','Grunge'),
    ('ouvinte1@email.com','Grunge'),
    ('ouvinte1@email.com','80s Rock');
-- -----------------------------------------------------
-- Table deezefy.Ouvinte_Salva_Album
-- -----------------------------------------------------
INSERT INTO deezefy.Ouvinte_Salva_Album (fk_ouvinte,fk_album,fk_artista)
VALUES ('ouvinte2@email.com',1,'nirvana@email.com'),
    ('ouvinte2@email.com',4,'oasis@email.com'),
    ('ouvinte1@email.com',2,'u2@email.com'),
    ('ouvinte1@email.com',3,'u2@email.com'),
    ('ouvinte1@email.com',5,'oasis@email.com');
-- -----------------------------------------------------
-- Table deezefy.Grava
-- -----------------------------------------------------
INSERT INTO deezefy.Grava (fk_musica,fk_artista)
VALUES (1,'oasis@email.com'),
    (2,'oasis@email.com'),
    (3,'oasis@email.com'),
    (4,'oasis@email.com'),
    (5,'oasis@email.com'),
    (6,'oasis@email.com'),
    (7,'u2@email.com'),
    (8,'u2@email.com'),
    (9,'u2@email.com'),
    (10,'u2@email.com'),
    (11,'u2@email.com'),
    (12,'nirvana@email.com'),
    (13,'nirvana@email.com');
-- -----------------------------------------------------
-- Table deezefy.Artistas_Favoritos
-- -----------------------------------------------------
-- \! echo 'Table deezefy.Artistas_Favoritos';
INSERT INTO deezefy.Artistas_Favoritos (fk_artista,fk_perfil,fk_ouvinte)
VALUES ('nirvana@email.com',2,'ouvinte2@email.com'),
    ('u2@email.com',1,'ouvinte1@email.com'),
    ('oasis@email.com',1,'ouvinte1@email.com');
-- -----------------------------------------------------
-- Table deezefy.Artista_Possui_Genero
-- -----------------------------------------------------
-- \! echo 'Table deezefy.Artista_Possui_Genero';
INSERT INTO deezefy.Artista_Possui_Genero (fk_genero,fk_artista)
VALUES ('80s Rock','oasis@email.com'),
    ('80s Rock','u2@email.com'),
    ('80s Rock','nirvana@email.com'),
    ('Grunge','nirvana@email.com');
-- -----------------------------------------------------
-- Table deezefy.Ocorre
-- -----------------------------------------------------
INSERT INTO deezefy.Ocorre (data,fk_local,fk_evento,fk_artista,fk_usuario)
VALUES ('1969-08-18',1,1,'oasis@email.com','oasis@email.com'),
    ('2010-08-21',2,2,'nirvana@email.com','nirvana@email.com');
-- -----------------------------------------------------
-- Table deezefy.Musica_em_Playlist
-- -----------------------------------------------------
INSERT INTO deezefy.Musica_em_Playlist (fk_musica,fk_playlist)
VALUES (12,'Grunge'),
    (13,'Grunge'),
    (1,'80s Rock'),
    (3,'80s Rock'),
    (4,'80s Rock'),
    (5,'80s Rock'),
    (6,'80s Rock'),
    (7,'80s Rock'),
    (10,'80s Rock'),
    (11,'80s Rock');
-- -----------------------------------------------------
-- Table deezefy.Album_Contem_Musica
-- -----------------------------------------------------
-- \! echo 'Table deezefy.Album_Contem_Musica';
INSERT INTO deezefy.Album_Contem_Musica (fk_musica,fk_album,fk_artista)
VALUES (1,5,'oasis@email.com'),
    (3,5,'oasis@email.com'),
    (2,5,'oasis@email.com'),
    (4,4,'oasis@email.com'),
    (5,4,'oasis@email.com'),
    (6,4,'oasis@email.com'),
    (7,2,'u2@email.com'),
    (8,2,'u2@email.com'),
    (9,3,'u2@email.com'),
    (10,3,'u2@email.com'),
    (11,3,'u2@email.com'),
    (12,1,'nirvana@email.com'),
    (13,1,'nirvana@email.com');
-- -----------------------------------------------------
-- Table deezefy.Musica_Possui_Generos
-- -----------------------------------------------------
-- \! echo 'Table deezefy.Musica_Possui_Generos';
INSERT INTO deezefy.Musica_Possui_Generos (fk_musica,fk_genero)
VALUES (1,'80s Rock'),
    (2,'80s Rock'),
    (3,'80s Rock'),
    (4,'80s Rock'),
    (5,'80s Rock'),
    (6,'80s Rock'),
    (7,'80s Rock'),
    (8,'80s Rock'),
    (9,'80s Rock'),
    (10,'80s Rock'),
    (11,'80s Rock'),
    (12,'Grunge'),
    (13,'Grunge');
-- -----------------------------------------------------
-- Table deezefy.Generos_Favoritos
-- -----------------------------------------------------
-- \! echo 'Table deezefy.Generos_Favoritos';
INSERT INTO deezefy.Generos_Favoritos (fk_genero,fk_perfil,fk_ouvinte)
VALUES ('Grunge',2,'ouvinte2@email.com'),
('80s Rock',1,'ouvinte1@email.com');
-- -----------------------------------------------------
-- Table deezefy.Cria
-- -----------------------------------------------------
INSERT INTO deezefy.Cria (fk_playlist,fk_usuario,data_criacao)
VALUES ('Grunge','ouvinte2@email.com','2019-10-21'),
('80s Rock','ouvinte1@email.com','2018-05-12');
-- -----------------------------------------------------
-- Schema deezefy
-- -----------------------------------------------------