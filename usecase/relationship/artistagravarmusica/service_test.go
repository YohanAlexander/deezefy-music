package artistagravarmusica

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	artistamock "github.com/yohanalexander/deezefy-music/usecase/entity/artista/mock"
	musicamock "github.com/yohanalexander/deezefy-music/usecase/entity/musica/mock"
)

func Test_Gravar(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	artistaMock := artistamock.NewMockUseCase(controller)
	musicaMock := musicamock.NewMockUseCase(controller)
	service := NewService(artistaMock, musicaMock)
	t.Run("Artista not found", func(t *testing.T) {
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(nil, entity.ErrNotFound)
		err := service.Gravar(a, m)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Musica not found", func(t *testing.T) {
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(a, nil)
		musicaMock.EXPECT().GetMusica(m.ID).Return(nil, entity.ErrNotFound)
		err := service.Gravar(a, m)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Musica already followed", func(t *testing.T) {
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		a.AddMusica(*m)
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(a, nil)
		musicaMock.EXPECT().GetMusica(m.ID).Return(m, nil)
		err := service.Gravar(a, m)
		assert.Equal(t, entity.ErrMusicaRegistered, err)
	})
	t.Run("Sucess", func(t *testing.T) {
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(a, nil)
		musicaMock.EXPECT().GetMusica(m.ID).Return(m, nil)
		artistaMock.EXPECT().UpdateArtista(a).Return(nil)
		musicaMock.EXPECT().UpdateMusica(m).Return(nil)
		err := service.Gravar(a, m)
		assert.Nil(t, err)
	})
}

func Test_Desgravar(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	artistaMock := artistamock.NewMockUseCase(controller)
	musicaMock := musicamock.NewMockUseCase(controller)
	service := NewService(artistaMock, musicaMock)
	t.Run("Artista not found", func(t *testing.T) {
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(nil, entity.ErrNotFound)
		err := service.Desgravar(a, m)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Musica not found", func(t *testing.T) {
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(a, nil)
		musicaMock.EXPECT().GetMusica(m.ID).Return(nil, entity.ErrNotFound)
		err := service.Desgravar(a, m)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Musica not followed", func(t *testing.T) {
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(a, nil)
		musicaMock.EXPECT().GetMusica(m.ID).Return(m, nil)
		err := service.Desgravar(a, m)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Success", func(t *testing.T) {
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		m := &entity.Musica{
			ID: 1,
		}
		a.AddMusica(*m)
		m.AddArtista(*a)
		musicaMock.EXPECT().GetMusica(m.ID).Return(m, nil)
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(a, nil)
		artistaMock.EXPECT().UpdateArtista(a).Return(nil)
		musicaMock.EXPECT().UpdateMusica(m).Return(nil)
		err := service.Desgravar(a, m)
		assert.Nil(t, err)
	})
}
