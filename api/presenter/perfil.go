package presenter

import "github.com/yohanalexander/deezefy-music/entity"

// Perfil entidade Perfil
type Perfil struct {
	Ouvinte               entity.Ouvinte   `json:"ouvinte"`
	ID                    int              `json:"id"`
	InformacoesRelevantes string           `json:"informacoesrelevantes"`
	ArtistasFavoritos     []entity.Artista `json:"artistasfavoritos"`
	GenerosFavoritos      []entity.Genero  `json:"generosfavoritos"`
}

// GetPerfil seta os valores a partir da entidade
func (p *Perfil) GetPerfil(perfil entity.Perfil) {
	p.Ouvinte = perfil.Ouvinte
	p.ID = perfil.ID
	p.InformacoesRelevantes = perfil.InformacoesRelevantes
	p.ArtistasFavoritos = perfil.ArtistasFavoritos
	p.GenerosFavoritos = perfil.GenerosFavoritos
}
