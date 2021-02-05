package presenter

// Artista entidade Artista
type Artista struct {
	Usuario       Usuario   `json:"usuario"`
	NomeArtistico string    `json:"nomeartistico"`
	Biografia     string    `json:"biografia"`
	AnoFormacao   int       `json:"anoformacao"`
	Seguidores    []Ouvinte `json:"seguidores"`
	Musicas       []Musica  `json:"musicas"`
	Perfis        []Perfil  `json:"perfis"`
	Generos       []Genero  `json:"generos"`
	Albums        []Album   `json:"albums"`
}
