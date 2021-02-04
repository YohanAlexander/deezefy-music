package artistapossuirgenero

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	artistamock "github.com/yohanalexander/deezefy-music/usecase/entity/artista/mock"
	generomock "github.com/yohanalexander/deezefy-music/usecase/entity/genero/mock"
)

func Test_Possuir(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	artistaMock := artistamock.NewMockUseCase(controller)
	generoMock := generomock.NewMockUseCase(controller)
	service := NewService(artistaMock, generoMock)
	t.Run("Artista not found", func(t *testing.T) {
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		g := &entity.Genero{
			Nome: "Genero",
		}
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(nil, entity.ErrNotFound)
		err := service.Possuir(a, g)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Genero not found", func(t *testing.T) {
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		g := &entity.Genero{
			Nome: "Genero",
		}
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(a, nil)
		generoMock.EXPECT().GetGenero(g.Nome).Return(nil, entity.ErrNotFound)
		err := service.Possuir(a, g)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Genero already followed", func(t *testing.T) {
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		g := &entity.Genero{
			Nome: "Genero",
		}
		a.AddGenero(*g)
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(a, nil)
		generoMock.EXPECT().GetGenero(g.Nome).Return(g, nil)
		err := service.Possuir(a, g)
		assert.Equal(t, entity.ErrGeneroRegistered, err)
	})
	t.Run("Sucess", func(t *testing.T) {
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		g := &entity.Genero{
			Nome: "Genero",
		}
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(a, nil)
		generoMock.EXPECT().GetGenero(g.Nome).Return(g, nil)
		artistaMock.EXPECT().UpdateArtista(a).Return(nil)
		generoMock.EXPECT().UpdateGenero(g).Return(nil)
		err := service.Possuir(a, g)
		assert.Nil(t, err)
	})
}

func Test_Despossuir(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	artistaMock := artistamock.NewMockUseCase(controller)
	generoMock := generomock.NewMockUseCase(controller)
	service := NewService(artistaMock, generoMock)
	t.Run("Artista not found", func(t *testing.T) {
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		g := &entity.Genero{
			Nome: "Genero",
		}
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(nil, entity.ErrNotFound)
		err := service.Despossuir(a, g)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Genero not found", func(t *testing.T) {
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		g := &entity.Genero{
			Nome: "Genero",
		}
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(a, nil)
		generoMock.EXPECT().GetGenero(g.Nome).Return(nil, entity.ErrNotFound)
		err := service.Despossuir(a, g)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Genero not followed", func(t *testing.T) {
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		g := &entity.Genero{
			Nome: "Genero",
		}
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(a, nil)
		generoMock.EXPECT().GetGenero(g.Nome).Return(g, nil)
		err := service.Despossuir(a, g)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Success", func(t *testing.T) {
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		g := &entity.Genero{
			Nome: "Genero",
		}
		a.AddGenero(*g)
		g.AddArtista(*a)
		generoMock.EXPECT().GetGenero(g.Nome).Return(g, nil)
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(a, nil)
		artistaMock.EXPECT().UpdateArtista(a).Return(nil)
		generoMock.EXPECT().UpdateGenero(g).Return(nil)
		err := service.Despossuir(a, g)
		assert.Nil(t, err)
	})
}
