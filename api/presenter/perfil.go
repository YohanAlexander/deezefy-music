package presenter

// Perfil entidade Perfil
type Perfil struct {
	Ouvinte               Ouvinte   `json:"ouvinte"`
	ID                    int       `json:"id"`
	InformacoesRelevantes string    `json:"informacoesrelevantes"`
	ArtistasFavoritos     []Artista `json:"artistasfavoritos"`
	GenerosFavoritos      []Genero  `json:"generosfavoritos"`
}
