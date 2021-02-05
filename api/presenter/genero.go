package presenter

// Genero entidade Genero
type Genero struct {
	Nome     string    `json:"nome"`
	Estilo   string    `json:"estilo"`
	Artistas []Artista `json:"artistas"`
	Musicas  []Musica  `json:"musicas"`
	Perfis   []Perfil  `json:"perfis"`
}
