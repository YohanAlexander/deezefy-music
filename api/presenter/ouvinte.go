package presenter

// Ouvinte entidade Ouvinte
type Ouvinte struct {
	Usuario      Usuario    `json:"usuario"`
	PrimeiroNome string     `json:"primeironome"`
	Sobrenome    string     `json:"sobrenome"`
	Telefones    []string   `json:"telefones"`
	Seguindo     []Artista  `json:"seguindo"`
	Curtidas     []Musica   `json:"curtidas"`
	Playlists    []Playlist `json:"playlists"`
	Albums       []Album    `json:"albums"`
}
