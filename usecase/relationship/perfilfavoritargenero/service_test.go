package perfilfavoritargenero

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	generomock "github.com/yohanalexander/deezefy-music/usecase/entity/genero/mock"
	perfilmock "github.com/yohanalexander/deezefy-music/usecase/entity/perfil/mock"
)

func Test_Favoritar(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	generoMock := generomock.NewMockUseCase(controller)
	perfilMock := perfilmock.NewMockUseCase(controller)
	service := NewService(generoMock, perfilMock)
	t.Run("Genero not found", func(t *testing.T) {
		g := &entity.Genero{
			Nome: "Genero",
		}
		p := &entity.Perfil{
			ID: 1,
		}
		generoMock.EXPECT().GetGenero(g.Nome).Return(nil, entity.ErrNotFound)
		err := service.Favoritar(g, p)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Perfil not found", func(t *testing.T) {
		g := &entity.Genero{
			Nome: "Genero",
		}
		p := &entity.Perfil{
			ID: 1,
		}
		generoMock.EXPECT().GetGenero(g.Nome).Return(g, nil)
		perfilMock.EXPECT().GetPerfil(p.ID).Return(nil, entity.ErrNotFound)
		err := service.Favoritar(g, p)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Perfil already followed", func(t *testing.T) {
		g := &entity.Genero{
			Nome: "Genero",
		}
		p := &entity.Perfil{
			ID: 1,
		}
		g.AddPerfil(*p)
		generoMock.EXPECT().GetGenero(g.Nome).Return(g, nil)
		perfilMock.EXPECT().GetPerfil(p.ID).Return(p, nil)
		err := service.Favoritar(g, p)
		assert.Equal(t, entity.ErrPerfilRegistered, err)
	})
	t.Run("Sucess", func(t *testing.T) {
		g := &entity.Genero{
			Nome: "Genero",
		}
		p := &entity.Perfil{
			ID: 1,
		}
		generoMock.EXPECT().GetGenero(g.Nome).Return(g, nil)
		perfilMock.EXPECT().GetPerfil(p.ID).Return(p, nil)
		generoMock.EXPECT().UpdateGenero(g).Return(nil)
		perfilMock.EXPECT().UpdatePerfil(p).Return(nil)
		err := service.Favoritar(g, p)
		assert.Nil(t, err)
	})
}

func Test_Desfavoritar(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	generoMock := generomock.NewMockUseCase(controller)
	perfilMock := perfilmock.NewMockUseCase(controller)
	service := NewService(generoMock, perfilMock)
	t.Run("Genero not found", func(t *testing.T) {
		g := &entity.Genero{
			Nome: "Genero",
		}
		p := &entity.Perfil{
			ID: 1,
		}
		generoMock.EXPECT().GetGenero(g.Nome).Return(nil, entity.ErrNotFound)
		err := service.Desfavoritar(g, p)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Perfil not found", func(t *testing.T) {
		g := &entity.Genero{
			Nome: "Genero",
		}
		p := &entity.Perfil{
			ID: 1,
		}
		generoMock.EXPECT().GetGenero(g.Nome).Return(g, nil)
		perfilMock.EXPECT().GetPerfil(p.ID).Return(nil, entity.ErrNotFound)
		err := service.Desfavoritar(g, p)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Perfil not followed", func(t *testing.T) {
		g := &entity.Genero{
			Nome: "Genero",
		}
		p := &entity.Perfil{
			ID: 1,
		}
		generoMock.EXPECT().GetGenero(g.Nome).Return(g, nil)
		perfilMock.EXPECT().GetPerfil(p.ID).Return(p, nil)
		err := service.Desfavoritar(g, p)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Success", func(t *testing.T) {
		g := &entity.Genero{
			Nome: "Genero",
		}
		p := &entity.Perfil{
			ID: 1,
		}
		g.AddPerfil(*p)
		p.AddGenero(*g)
		perfilMock.EXPECT().GetPerfil(p.ID).Return(p, nil)
		generoMock.EXPECT().GetGenero(g.Nome).Return(g, nil)
		generoMock.EXPECT().UpdateGenero(g).Return(nil)
		perfilMock.EXPECT().UpdatePerfil(p).Return(nil)
		err := service.Desfavoritar(g, p)
		assert.Nil(t, err)
	})
}
