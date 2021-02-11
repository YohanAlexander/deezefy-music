package presenter

import "github.com/yohanalexander/deezefy-music/entity"

// PostPerfil entidade Perfil
type PostPerfil struct {
	Ouvinte               entity.Ouvinte   `json:"ouvinte"`
	ID                    int              `json:"id"`
	InformacoesRelevantes string           `json:"informacoes_relevantes"`
	ArtistasFavoritos     []entity.Artista `json:"artistas_favoritos"`
	GenerosFavoritos      []entity.Genero  `json:"generos_favoritos"`
}

// MakePerfil seta os valores a partir da entidade
func (p *PostPerfil) MakePerfil(perfil entity.Perfil) {
	p.Ouvinte = perfil.Ouvinte
	p.ID = perfil.ID
	p.InformacoesRelevantes = perfil.InformacoesRelevantes
	p.ArtistasFavoritos = perfil.ArtistasFavoritos
	p.GenerosFavoritos = perfil.GenerosFavoritos
}

// Perfil presenter Perfil
type Perfil struct {
	Ouvinte               Ouvinte   `json:"ouvinte"`
	ID                    int       `json:"id"`
	InformacoesRelevantes string    `json:"informacoes_relevantes"`
	ArtistasFavoritos     []Artista `json:"artistas_favoritos"`
	GenerosFavoritos      []Genero  `json:"generos_favoritos"`
}

// AppendPerfil adiciona presenter na lista
func AppendPerfil(perfil entity.Perfil, perfils []Perfil) []Perfil {
	p := &Perfil{}
	p.MakePerfil(perfil)
	perfils = append(perfils, *p)
	return perfils
}

// MakePerfil seta os valores a partir da entidade
func (p *Perfil) MakePerfil(perfil entity.Perfil) {
	p.Ouvinte.MakeOuvinte(perfil.Ouvinte)
	p.ID = perfil.ID
	p.InformacoesRelevantes = perfil.InformacoesRelevantes
	for _, artista := range perfil.ArtistasFavoritos {
		p.ArtistasFavoritos = AppendArtista(artista, p.ArtistasFavoritos)
	}
	for _, genero := range perfil.GenerosFavoritos {
		p.GenerosFavoritos = AppendGenero(genero, p.GenerosFavoritos)
	}
}
