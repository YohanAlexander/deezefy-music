package albumcontermusica

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	musicamock "github.com/yohanalexander/deezefy-music/usecase/entity/musica/mock"
	albummock "github.com/yohanalexander/deezefy-music/usecase/entity/album/mock"
)

func Test_Conter(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	albumMock := albummock.NewMockUseCase(controller)
	musicaMock := musicamock.NewMockUseCase(controller)
	service := NewService(albumMock, musicaMock)
	t.Run("Album not found", func(t *testing.T) {
		a := &entity.Album{
			ID: 1,
		}
		m := &entity.Musica{
			ID: 1,
		}
		albumMock.EXPECT().GetAlbum(a.ID).Return(nil, entity.ErrNotFound)
		err := service.Conter(a, m)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Musica not found", func(t *testing.T) {
		a := &entity.Album{
			ID: 1,
		}
		m := &entity.Musica{
			ID: 1,
		}
		albumMock.EXPECT().GetAlbum(a.ID).Return(a, nil)
		musicaMock.EXPECT().GetMusica(m.ID).Return(nil, entity.ErrNotFound)
		err := service.Conter(a, m)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Musica already followed", func(t *testing.T) {
		a := &entity.Album{
			ID: 1,
		}
		m := &entity.Musica{
			ID: 1,
		}
		a.AddMusica(*m)
		albumMock.EXPECT().GetAlbum(a.ID).Return(a, nil)
		musicaMock.EXPECT().GetMusica(m.ID).Return(m, nil)
		err := service.Conter(a, m)
		assert.Equal(t, entity.ErrMusicaRegistered, err)
	})
	t.Run("Sucess", func(t *testing.T) {
		a := &entity.Album{
			ID: 1,
		}
		m := &entity.Musica{
			ID: 1,
		}
		albumMock.EXPECT().GetAlbum(a.ID).Return(a, nil)
		musicaMock.EXPECT().GetMusica(m.ID).Return(m, nil)
		albumMock.EXPECT().UpdateAlbum(a).Return(nil)
		musicaMock.EXPECT().UpdateMusica(m).Return(nil)
		err := service.Conter(a, m)
		assert.Nil(t, err)
	})
}

func Test_Desconter(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	albumMock := albummock.NewMockUseCase(controller)
	musicaMock := musicamock.NewMockUseCase(controller)
	service := NewService(albumMock, musicaMock)
	t.Run("Album not found", func(t *testing.T) {
		a := &entity.Album{
			ID: 1,
		}
		m := &entity.Musica{
			ID: 1,
		}
		albumMock.EXPECT().GetAlbum(a.ID).Return(nil, entity.ErrNotFound)
		err := service.Desconter(a, m)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Musica not found", func(t *testing.T) {
		a := &entity.Album{
			ID: 1,
		}
		m := &entity.Musica{
			ID: 1,
		}
		albumMock.EXPECT().GetAlbum(a.ID).Return(a, nil)
		musicaMock.EXPECT().GetMusica(m.ID).Return(nil, entity.ErrNotFound)
		err := service.Desconter(a, m)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Musica not followed", func(t *testing.T) {
		a := &entity.Album{
			ID: 1,
		}
		m := &entity.Musica{
			ID: 1,
		}
		albumMock.EXPECT().GetAlbum(a.ID).Return(a, nil)
		musicaMock.EXPECT().GetMusica(m.ID).Return(m, nil)
		err := service.Desconter(a, m)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Success", func(t *testing.T) {
		a := &entity.Album{
			ID: 1,
		}
		m := &entity.Musica{
			ID: 1,
		}
		a.AddMusica(*m)
		m.AddAlbum(*a)
		musicaMock.EXPECT().GetMusica(m.ID).Return(m, nil)
		albumMock.EXPECT().GetAlbum(a.ID).Return(a, nil)
		albumMock.EXPECT().UpdateAlbum(a).Return(nil)
		musicaMock.EXPECT().UpdateMusica(m).Return(nil)
		err := service.Desconter(a, m)
		assert.Nil(t, err)
	})
}
