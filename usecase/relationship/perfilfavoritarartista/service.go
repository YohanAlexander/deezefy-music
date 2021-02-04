package perfilavoritarartista

import (
	"github.com/yohanalexander/deezefy-music/entity"
	"github.com/yohanalexander/deezefy-music/usecase/entity/artista"
	"github.com/yohanalexander/deezefy-music/usecase/entity/perfil"
)

// Service  interface
type Service struct {
	artistaService artista.UseCase
	perfilService  perfil.UseCase
}

// NewService create new use case
func NewService(a artista.UseCase, p perfil.UseCase) *Service {
	return &Service{
		artistaService: a,
		perfilService:  p,
	}
}

// Favoritar favorita um artista
func (s *Service) Favoritar(a *entity.Artista, p *entity.Perfil) error {
	a, err := s.artistaService.GetArtista(a.Usuario.Email)
	if err != nil {
		return err
	}
	p, err = s.perfilService.GetPerfil(p.ID)
	if err != nil {
		return err
	}

	err = p.AddArtista(*a)
	if err != nil {
		return err
	}
	err = a.AddPerfil(*p)
	if err != nil {
		return err
	}

	err = s.artistaService.UpdateArtista(a)
	if err != nil {
		return err
	}
	err = s.perfilService.UpdatePerfil(p)
	if err != nil {
		return err
	}
	return nil
}

// Desfavoritar favorita um artista
func (s *Service) Desfavoritar(a *entity.Artista, p *entity.Perfil) error {
	a, err := s.artistaService.GetArtista(a.Usuario.Email)
	if err != nil {
		return err
	}
	p, err = s.perfilService.GetPerfil(p.ID)
	if err != nil {
		return err
	}

	err = p.RemoveArtista(*a)
	if err != nil {
		return err
	}
	err = a.RemovePerfil(*p)
	if err != nil {
		return err
	}

	err = s.artistaService.UpdateArtista(a)
	if err != nil {
		return err
	}
	err = s.perfilService.UpdatePerfil(p)
	if err != nil {
		return err
	}
	return nil
}
