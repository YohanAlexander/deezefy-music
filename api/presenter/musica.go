package presenter

// Musica entidade Musica
type Musica struct {
	ID        int        `json:"id"`
	Nome      string     `json:"nome"`
	Duracao   int        `json:"duracao"`
	Curtiu    []Ouvinte  `json:"curtiu"`
	Artistas  []Artista  `json:"artistas"`
	Playlists []Playlist `json:"playlists"`
	Albums    []Album    `json:"albums"`
	Generos   []Genero   `json:"generos"`
}
