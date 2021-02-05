package presenter

// Usuario entidade usuario
type Usuario struct {
	Email       string     `json:"email"`
	Password    string     `json:"-"`
	Birthday    string     `json:"datanascimento"`
	Organizador []Evento   `json:"eventos"`
	Cria        []Playlist `json:"playlists"`
	Idade       int        `json:"idade"`
}
