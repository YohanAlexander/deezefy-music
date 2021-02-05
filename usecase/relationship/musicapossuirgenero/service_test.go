package musicapossuirgenero

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	generomock "github.com/yohanalexander/deezefy-music/usecase/entity/genero/mock"
	musicamock "github.com/yohanalexander/deezefy-music/usecase/entity/musica/mock"
)

func Test_Possuir(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	musicaMock := musicamock.NewMockUseCase(controller)
	generoMock := generomock.NewMockUseCase(controller)
	service := NewService(musicaMock, generoMock)
	t.Run("Musica not found", func(t *testing.T) {
		m := &entity.Musica{
			ID: 1,
		}
		g := &entity.Genero{
			Nome: "Genero",
		}
		musicaMock.EXPECT().GetMusica(m.ID).Return(nil, entity.ErrNotFound)
		err := service.Possuir(m, g)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Genero not found", func(t *testing.T) {
		m := &entity.Musica{
			ID: 1,
		}
		g := &entity.Genero{
			Nome: "Genero",
		}
		musicaMock.EXPECT().GetMusica(m.ID).Return(m, nil)
		generoMock.EXPECT().GetGenero(g.Nome).Return(nil, entity.ErrNotFound)
		err := service.Possuir(m, g)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Genero already followed", func(t *testing.T) {
		m := &entity.Musica{
			ID: 1,
		}
		g := &entity.Genero{
			Nome: "Genero",
		}
		m.AddGenero(*g)
		musicaMock.EXPECT().GetMusica(m.ID).Return(m, nil)
		generoMock.EXPECT().GetGenero(g.Nome).Return(g, nil)
		err := service.Possuir(m, g)
		assert.Equal(t, entity.ErrGeneroRegistered, err)
	})
	t.Run("Sucess", func(t *testing.T) {
		m := &entity.Musica{
			ID: 1,
		}
		g := &entity.Genero{
			Nome: "Genero",
		}
		musicaMock.EXPECT().GetMusica(m.ID).Return(m, nil)
		generoMock.EXPECT().GetGenero(g.Nome).Return(g, nil)
		musicaMock.EXPECT().UpdateMusica(m).Return(nil)
		generoMock.EXPECT().UpdateGenero(g).Return(nil)
		err := service.Possuir(m, g)
		assert.Nil(t, err)
	})
}

func Test_Despossuir(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	musicaMock := musicamock.NewMockUseCase(controller)
	generoMock := generomock.NewMockUseCase(controller)
	service := NewService(musicaMock, generoMock)
	t.Run("Musica not found", func(t *testing.T) {
		m := &entity.Musica{
			ID: 1,
		}
		g := &entity.Genero{
			Nome: "Genero",
		}
		musicaMock.EXPECT().GetMusica(m.ID).Return(nil, entity.ErrNotFound)
		err := service.Despossuir(m, g)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Genero not found", func(t *testing.T) {
		m := &entity.Musica{
			ID: 1,
		}
		g := &entity.Genero{
			Nome: "Genero",
		}
		musicaMock.EXPECT().GetMusica(m.ID).Return(m, nil)
		generoMock.EXPECT().GetGenero(g.Nome).Return(nil, entity.ErrNotFound)
		err := service.Despossuir(m, g)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Genero not followed", func(t *testing.T) {
		m := &entity.Musica{
			ID: 1,
		}
		g := &entity.Genero{
			Nome: "Genero",
		}
		musicaMock.EXPECT().GetMusica(m.ID).Return(m, nil)
		generoMock.EXPECT().GetGenero(g.Nome).Return(g, nil)
		err := service.Despossuir(m, g)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Success", func(t *testing.T) {
		m := &entity.Musica{
			ID: 1,
		}
		g := &entity.Genero{
			Nome: "Genero",
		}
		m.AddGenero(*g)
		g.AddMusica(*m)
		generoMock.EXPECT().GetGenero(g.Nome).Return(g, nil)
		musicaMock.EXPECT().GetMusica(m.ID).Return(m, nil)
		musicaMock.EXPECT().UpdateMusica(m).Return(nil)
		generoMock.EXPECT().UpdateGenero(g).Return(nil)
		err := service.Despossuir(m, g)
		assert.Nil(t, err)
	})
}
