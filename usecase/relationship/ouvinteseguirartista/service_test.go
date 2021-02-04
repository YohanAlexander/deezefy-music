package ouvinteseguirartista

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	artistamock "github.com/yohanalexander/deezefy-music/usecase/entity/artista/mock"
	ouvintemock "github.com/yohanalexander/deezefy-music/usecase/entity/ouvinte/mock"
)

func Test_Seguir(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	ouvinteMock := ouvintemock.NewMockUseCase(controller)
	artistaMock := artistamock.NewMockUseCase(controller)
	service := NewService(ouvinteMock, artistaMock)
	t.Run("Ouvinte not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(nil, entity.ErrNotFound)
		err := service.Seguir(o, a)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Artista not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(nil, entity.ErrNotFound)
		err := service.Seguir(o, a)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Artista already followed", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		o.AddArtista(*a)
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(a, nil)
		err := service.Seguir(o, a)
		assert.Equal(t, entity.ErrArtistaRegistered, err)
	})
	t.Run("Sucess", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(a, nil)
		ouvinteMock.EXPECT().UpdateOuvinte(o).Return(nil)
		artistaMock.EXPECT().UpdateArtista(a).Return(nil)
		err := service.Seguir(o, a)
		assert.Nil(t, err)
	})
}

func Test_Desseguir(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	ouvinteMock := ouvintemock.NewMockUseCase(controller)
	artistaMock := artistamock.NewMockUseCase(controller)
	service := NewService(ouvinteMock, artistaMock)
	t.Run("Ouvinte not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(nil, entity.ErrNotFound)
		err := service.Desseguir(o, a)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Artista not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(nil, entity.ErrNotFound)
		err := service.Desseguir(o, a)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Artista not followed", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(a, nil)
		err := service.Desseguir(o, a)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Success", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		o.AddArtista(*a)
		a.AddOuvinte(*o)
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(a, nil)
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		ouvinteMock.EXPECT().UpdateOuvinte(o).Return(nil)
		artistaMock.EXPECT().UpdateArtista(a).Return(nil)
		err := service.Desseguir(o, a)
		assert.Nil(t, err)
	})
}
