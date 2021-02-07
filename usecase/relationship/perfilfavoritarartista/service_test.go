package perfilfavoritarartista

import (
	"testing"

	"github.com/yohanalexander/deezefy-music/entity"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	artistamock "github.com/yohanalexander/deezefy-music/usecase/entity/artista/mock"
	perfilmock "github.com/yohanalexander/deezefy-music/usecase/entity/perfil/mock"
)

func Test_Favoritar(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	artistaMock := artistamock.NewMockUseCase(controller)
	perfilMock := perfilmock.NewMockUseCase(controller)
	service := NewService(artistaMock, perfilMock)
	t.Run("Artista not found", func(t *testing.T) {
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		p := &entity.Perfil{
			Ouvinte: entity.Ouvinte{
				Usuario: entity.Usuario{
					Email: "ouvinte@email.com",
				},
			},
		}
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(nil, entity.ErrNotFound)
		err := service.Favoritar(a, p)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Perfil not found", func(t *testing.T) {
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		p := &entity.Perfil{
			Ouvinte: entity.Ouvinte{
				Usuario: entity.Usuario{
					Email: "ouvinte@email.com",
				},
			},
		}
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(a, nil)
		perfilMock.EXPECT().GetPerfil(p.Ouvinte.Usuario.Email).Return(nil, entity.ErrNotFound)
		err := service.Favoritar(a, p)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Perfil already followed", func(t *testing.T) {
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		p := &entity.Perfil{
			Ouvinte: entity.Ouvinte{
				Usuario: entity.Usuario{
					Email: "ouvinte@email.com",
				},
			},
		}
		a.AddPerfil(*p)
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(a, nil)
		perfilMock.EXPECT().GetPerfil(p.Ouvinte.Usuario.Email).Return(p, nil)
		err := service.Favoritar(a, p)
		assert.Equal(t, entity.ErrPerfilRegistered, err)
	})
	t.Run("Sucess", func(t *testing.T) {
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		p := &entity.Perfil{
			Ouvinte: entity.Ouvinte{
				Usuario: entity.Usuario{
					Email: "ouvinte@email.com",
				},
			},
		}
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(a, nil)
		perfilMock.EXPECT().GetPerfil(p.Ouvinte.Usuario.Email).Return(p, nil)
		artistaMock.EXPECT().UpdateArtista(a).Return(nil)
		perfilMock.EXPECT().UpdatePerfil(p).Return(nil)
		err := service.Favoritar(a, p)
		assert.Nil(t, err)
	})
}

func Test_Desfavoritar(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	artistaMock := artistamock.NewMockUseCase(controller)
	perfilMock := perfilmock.NewMockUseCase(controller)
	service := NewService(artistaMock, perfilMock)
	t.Run("Artista not found", func(t *testing.T) {
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		p := &entity.Perfil{
			Ouvinte: entity.Ouvinte{
				Usuario: entity.Usuario{
					Email: "ouvinte@email.com",
				},
			},
		}
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(nil, entity.ErrNotFound)
		err := service.Desfavoritar(a, p)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Perfil not found", func(t *testing.T) {
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		p := &entity.Perfil{
			Ouvinte: entity.Ouvinte{
				Usuario: entity.Usuario{
					Email: "ouvinte@email.com",
				},
			},
		}
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(a, nil)
		perfilMock.EXPECT().GetPerfil(p.Ouvinte.Usuario.Email).Return(nil, entity.ErrNotFound)
		err := service.Desfavoritar(a, p)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Perfil not followed", func(t *testing.T) {
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		p := &entity.Perfil{
			Ouvinte: entity.Ouvinte{
				Usuario: entity.Usuario{
					Email: "ouvinte@email.com",
				},
			},
		}
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(a, nil)
		perfilMock.EXPECT().GetPerfil(p.Ouvinte.Usuario.Email).Return(p, nil)
		err := service.Desfavoritar(a, p)
		assert.Equal(t, entity.ErrNotFound, err)
	})
	t.Run("Success", func(t *testing.T) {
		a := &entity.Artista{
			Usuario: entity.Usuario{
				Email: "artista@email.com",
			},
		}
		p := &entity.Perfil{
			Ouvinte: entity.Ouvinte{
				Usuario: entity.Usuario{
					Email: "ouvinte@email.com",
				},
			},
		}
		a.AddPerfil(*p)
		p.AddArtista(*a)
		perfilMock.EXPECT().GetPerfil(p.Ouvinte.Usuario.Email).Return(p, nil)
		artistaMock.EXPECT().GetArtista(a.Usuario.Email).Return(a, nil)
		artistaMock.EXPECT().UpdateArtista(a).Return(nil)
		perfilMock.EXPECT().UpdatePerfil(p).Return(nil)
		err := service.Desfavoritar(a, p)
		assert.Nil(t, err)
	})
}
