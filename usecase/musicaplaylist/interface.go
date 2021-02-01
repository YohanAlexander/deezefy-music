package musicaplaylist

import der "github.com/yohanalexander/deezefy-music/entity/musicaplaylist"

// MusicaPlaylist interface
type MusicaPlaylist interface {
	Get(musica int, playlist string) (*der.MusicaPlaylist, error)
	GetByMusica(musica int) (*der.MusicaPlaylist, error)
	GetByPlaylist(playlist string) (*der.MusicaPlaylist, error)
	Search(query string) ([]*der.MusicaPlaylist, error)
	List() ([]*der.MusicaPlaylist, error)
	Create(e *der.MusicaPlaylist) (int, string, error)
	Update(e *der.MusicaPlaylist) error
	Delete(musica int, playlist string) error
}

// Repository interface
type Repository interface {
	MusicaPlaylist
}

// UseCase interface
type UseCase interface {
	GetMusicaPlaylist(musica int, playlist string) (*der.MusicaPlaylist, error)
	GetMusicaPlaylistByMusica(Musica int) (*der.MusicaPlaylist, error)
	GetMusicaPlaylistByPlaylist(playlist string) (*der.MusicaPlaylist, error)
	SearchMusicaPlaylists(query string) ([]*der.MusicaPlaylist, error)
	ListMusicaPlaylists() ([]*der.MusicaPlaylist, error)
	CreateMusicaPlaylist(musica int, playlist string) (int, string, error)
	UpdateMusicaPlaylist(e *der.MusicaPlaylist) error
	DeleteMusicaPlaylist(musica int, playlist string) error
}
