package presenter

// Playlist entidade Playlist
type Playlist struct {
	Nome           string    `json:"nome"`
	Status         string    `json:"status"`
	DataCriacao    string    `json:"datacriacao"`
	Ouvintes       []Ouvinte `json:"ouvintes"`
	Musicas        []Musica  `json:"musicas"`
	NumeroOuvintes int       `json:"numeroouvintes"`
}
