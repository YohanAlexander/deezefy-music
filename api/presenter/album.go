package presenter

// Album entidade Album
type Album struct {
	Artista       Artista   `json:"artista"`
	ID            int       `json:"id"`
	Titulo        string    `json:"titulo"`
	AnoLancamento int       `json:"anolancamento"`
	Ouvintes      []Ouvinte `json:"ouvintes"`
	Musicas       []Musica  `json:"musicas"`
}
