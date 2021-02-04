package ouvintecurtirmusica

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	musicamock "github.com/yohanalexander/deezefy-music/usecase/entity/musica/mock"
	ouvintemock "github.com/yohanalexander/deezefy-music/usecase/entity/ouvinte/mock"
)

func Test_Curtir(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	ouvinteMock := ouvintemock.NewMockUseCase(controller)
	musicaMock := musicamock.NewMockUseCase(controller)
	service := NewService(ouvinteMock, musicaMock)
	t.Run("Ouvinte not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(nil, entity.ErrNotFound)
		err := service.Curtir(o, m)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Musica not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		musicaMock.EXPECT().GetMusica(m.ID).Return(nil, entity.ErrNotFound)
		err := service.Curtir(o, m)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Musica already followed", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		o.AddMusica(*m)
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		musicaMock.EXPECT().GetMusica(m.ID).Return(m, nil)
		err := service.Curtir(o, m)
		assert.Equal(t, entity.ErrMusicaRegistered, err)
	})
	t.Run("Sucess", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		musicaMock.EXPECT().GetMusica(m.ID).Return(m, nil)
		ouvinteMock.EXPECT().UpdateOuvinte(o).Return(nil)
		musicaMock.EXPECT().UpdateMusica(m).Return(nil)
		err := service.Curtir(o, m)
		assert.Nil(t, err)
	})
}

func Test_Descurtir(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	ouvinteMock := ouvintemock.NewMockUseCase(controller)
	musicaMock := musicamock.NewMockUseCase(controller)
	service := NewService(ouvinteMock, musicaMock)
	t.Run("Ouvinte not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(nil, entity.ErrNotFound)
		err := service.Descurtir(o, m)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Musica not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		musicaMock.EXPECT().GetMusica(m.ID).Return(nil, entity.ErrNotFound)
		err := service.Descurtir(o, m)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Musica not followed", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		musicaMock.EXPECT().GetMusica(m.ID).Return(m, nil)
		err := service.Descurtir(o, m)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Success", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		o.AddMusica(*m)
		m.AddOuvinte(*o)
		musicaMock.EXPECT().GetMusica(m.ID).Return(m, nil)
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		ouvinteMock.EXPECT().UpdateOuvinte(o).Return(nil)
		musicaMock.EXPECT().UpdateMusica(m).Return(nil)
		err := service.Descurtir(o, m)
		assert.Nil(t, err)
	})
}
