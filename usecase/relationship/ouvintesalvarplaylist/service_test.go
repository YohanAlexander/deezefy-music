package ouvintesalvarplaylist

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	ouvintemock "github.com/yohanalexander/deezefy-music/usecase/entity/ouvinte/mock"
	playlistmock "github.com/yohanalexander/deezefy-music/usecase/entity/playlist/mock"
)

func Test_Salvar(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	ouvinteMock := ouvintemock.NewMockUseCase(controller)
	playlistMock := playlistmock.NewMockUseCase(controller)
	service := NewService(ouvinteMock, playlistMock)
	t.Run("Ouvinte not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		p := &entity.Playlist{
			Nome: "Playlist",
		}
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(nil, entity.ErrNotFound)
		err := service.Salvar(o, p)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Playlist not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		p := &entity.Playlist{
			Nome: "Playlist",
		}
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		playlistMock.EXPECT().GetPlaylist(p.Nome).Return(nil, entity.ErrNotFound)
		err := service.Salvar(o, p)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Playlist already followed", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		p := &entity.Playlist{
			Nome: "Playlist",
		}
		o.AddPlaylist(*p)
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		playlistMock.EXPECT().GetPlaylist(p.Nome).Return(p, nil)
		err := service.Salvar(o, p)
		assert.Equal(t, entity.ErrPlaylistRegistered, err)
	})
	t.Run("Sucess", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		p := &entity.Playlist{
			Nome: "Playlist",
		}
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		playlistMock.EXPECT().GetPlaylist(p.Nome).Return(p, nil)
		ouvinteMock.EXPECT().UpdateOuvinte(o).Return(nil)
		playlistMock.EXPECT().UpdatePlaylist(p).Return(nil)
		err := service.Salvar(o, p)
		assert.Nil(t, err)
	})
}

func Test_Dessalvar(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	ouvinteMock := ouvintemock.NewMockUseCase(controller)
	playlistMock := playlistmock.NewMockUseCase(controller)
	service := NewService(ouvinteMock, playlistMock)
	t.Run("Ouvinte not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		p := &entity.Playlist{
			Nome: "Playlist",
		}
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(nil, entity.ErrNotFound)
		err := service.Dessalvar(o, p)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Playlist not found", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		p := &entity.Playlist{
			Nome: "Playlist",
		}
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		playlistMock.EXPECT().GetPlaylist(p.Nome).Return(nil, entity.ErrNotFound)
		err := service.Dessalvar(o, p)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Playlist not followed", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		p := &entity.Playlist{
			Nome: "Playlist",
		}
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		playlistMock.EXPECT().GetPlaylist(p.Nome).Return(p, nil)
		err := service.Dessalvar(o, p)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Success", func(t *testing.T) {
		o := &entity.Ouvinte{
			Usuario: entity.Usuario{
				Email: "ouvinte@email.com",
			},
		}
		p := &entity.Playlist{
			Nome: "Playlist",
		}
		o.AddPlaylist(*p)
		p.AddOuvinte(*o)
		playlistMock.EXPECT().GetPlaylist(p.Nome).Return(p, nil)
		ouvinteMock.EXPECT().GetOuvinte(o.Usuario.Email).Return(o, nil)
		ouvinteMock.EXPECT().UpdateOuvinte(o).Return(nil)
		playlistMock.EXPECT().UpdatePlaylist(p).Return(nil)
		err := service.Dessalvar(o, p)
		assert.Nil(t, err)
	})
}
