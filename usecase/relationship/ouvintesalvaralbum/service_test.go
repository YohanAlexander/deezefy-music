package ouvintesalvaralbum

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	albummock "github.com/yohanalexander/deezefy-music/usecase/entity/album/mock"
	ouvintemock "github.com/yohanalexander/deezefy-music/usecase/entity/ouvinte/mock"
)

func Test_Salvar(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	ouvinteMock := ouvintemock.NewMockUseCase(controller)
	albumMock := albummock.NewMockUseCase(controller)
	service := NewService(ouvinteMock, albumMock)
	t.Run("Ouvinte not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Album{
			ID: 1,
		}
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(nil, entity.ErrNotFound)
		err := service.Salvar(o, a)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Album not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Album{
			ID: 1,
		}
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		albumMock.EXPECT().GetAlbum(a.ID).Return(nil, entity.ErrNotFound)
		err := service.Salvar(o, a)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Album already followed", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Album{
			ID: 1,
		}
		o.AddAlbum(*a)
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		albumMock.EXPECT().GetAlbum(a.ID).Return(a, nil)
		err := service.Salvar(o, a)
		assert.Equal(t, entity.ErrAlbumRegistered, err)
	})
	t.Run("Sucess", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Album{
			ID: 1,
		}
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		albumMock.EXPECT().GetAlbum(a.ID).Return(a, nil)
		ouvinteMock.EXPECT().UpdateOuvinte(o).Return(nil)
		albumMock.EXPECT().UpdateAlbum(a).Return(nil)
		err := service.Salvar(o, a)
		assert.Nil(t, err)
	})
}

func Test_DessalvarAlbum(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	ouvinteMock := ouvintemock.NewMockUseCase(controller)
	albumMock := albummock.NewMockUseCase(controller)
	service := NewService(ouvinteMock, albumMock)
	t.Run("Ouvinte not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Album{
			ID: 1,
		}
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(nil, entity.ErrNotFound)
		err := service.Dessalvar(o, a)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Album not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Album{
			ID: 1,
		}
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		albumMock.EXPECT().GetAlbum(a.ID).Return(nil, entity.ErrNotFound)
		err := service.Dessalvar(o, a)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Album not followed", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Album{
			ID: 1,
		}
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		albumMock.EXPECT().GetAlbum(a.ID).Return(a, nil)
		err := service.Dessalvar(o, a)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Success", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Album{
			ID: 1,
		}
		o.AddAlbum(*a)
		a.AddOuvinte(*o)
		albumMock.EXPECT().GetAlbum(a.ID).Return(a, nil)
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		ouvinteMock.EXPECT().UpdateOuvinte(o).Return(nil)
		albumMock.EXPECT().UpdateAlbum(a).Return(nil)
		err := service.Dessalvar(o, a)
		assert.Nil(t, err)
	})
}
