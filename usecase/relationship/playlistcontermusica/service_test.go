package playlistcontermusica

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	musicamock "github.com/yohanalexander/deezefy-music/usecase/entity/musica/mock"
	playlistmock "github.com/yohanalexander/deezefy-music/usecase/entity/playlist/mock"
)

func Test_Conter(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	playlistMock := playlistmock.NewMockUseCase(controller)
	musicaMock := musicamock.NewMockUseCase(controller)
	service := NewService(playlistMock, musicaMock)
	t.Run("Playlist not found", func(t *testing.T) {
		p := &entity.Playlist{
			Nome: "Playlist",
		}
		m := &entity.Musica{
			ID: 1,
		}
		playlistMock.EXPECT().GetPlaylist(p.Nome).Return(nil, entity.ErrNotFound)
		err := service.Conter(p, m)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Musica not found", func(t *testing.T) {
		p := &entity.Playlist{
			Nome: "Playlist",
		}
		m := &entity.Musica{
			ID: 1,
		}
		playlistMock.EXPECT().GetPlaylist(p.Nome).Return(p, nil)
		musicaMock.EXPECT().GetMusica(m.ID).Return(nil, entity.ErrNotFound)
		err := service.Conter(p, m)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Musica already followed", func(t *testing.T) {
		p := &entity.Playlist{
			Nome: "Playlist",
		}
		m := &entity.Musica{
			ID: 1,
		}
		p.AddMusica(*m)
		playlistMock.EXPECT().GetPlaylist(p.Nome).Return(p, nil)
		musicaMock.EXPECT().GetMusica(m.ID).Return(m, nil)
		err := service.Conter(p, m)
		assert.Equal(t, entity.ErrMusicaRegistered, err)
	})
	t.Run("Sucess", func(t *testing.T) {
		p := &entity.Playlist{
			Nome: "Playlist",
		}
		m := &entity.Musica{
			ID: 1,
		}
		playlistMock.EXPECT().GetPlaylist(p.Nome).Return(p, nil)
		musicaMock.EXPECT().GetMusica(m.ID).Return(m, nil)
		playlistMock.EXPECT().UpdatePlaylist(p).Return(nil)
		musicaMock.EXPECT().UpdateMusica(m).Return(nil)
		err := service.Conter(p, m)
		assert.Nil(t, err)
	})
}

func Test_Desconter(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	playlistMock := playlistmock.NewMockUseCase(controller)
	musicaMock := musicamock.NewMockUseCase(controller)
	service := NewService(playlistMock, musicaMock)
	t.Run("Playlist not found", func(t *testing.T) {
		p := &entity.Playlist{
			Nome: "Playlist",
		}
		m := &entity.Musica{
			ID: 1,
		}
		playlistMock.EXPECT().GetPlaylist(p.Nome).Return(nil, entity.ErrNotFound)
		err := service.Desconter(p, m)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Musica not found", func(t *testing.T) {
		p := &entity.Playlist{
			Nome: "Playlist",
		}
		m := &entity.Musica{
			ID: 1,
		}
		playlistMock.EXPECT().GetPlaylist(p.Nome).Return(p, nil)
		musicaMock.EXPECT().GetMusica(m.ID).Return(nil, entity.ErrNotFound)
		err := service.Desconter(p, m)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Musica not followed", func(t *testing.T) {
		p := &entity.Playlist{
			Nome: "Playlist",
		}
		m := &entity.Musica{
			ID: 1,
		}
		playlistMock.EXPECT().GetPlaylist(p.Nome).Return(p, nil)
		musicaMock.EXPECT().GetMusica(m.ID).Return(m, nil)
		err := service.Desconter(p, m)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Success", func(t *testing.T) {
		p := &entity.Playlist{
			Nome: "Playlist",
		}
		m := &entity.Musica{
			ID: 1,
		}
		p.AddMusica(*m)
		m.AddPlaylist(*p)
		musicaMock.EXPECT().GetMusica(m.ID).Return(m, nil)
		playlistMock.EXPECT().GetPlaylist(p.Nome).Return(p, nil)
		playlistMock.EXPECT().UpdatePlaylist(p).Return(nil)
		musicaMock.EXPECT().UpdateMusica(m).Return(nil)
		err := service.Desconter(p, m)
		assert.Nil(t, err)
	})
}
