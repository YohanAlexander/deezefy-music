package presenter

import (
	"github.com/yohanalexander/deezefy-music/entity"
)

// PostLocal entidade Local
type PostLocal struct {
	ID     int    `json:"id"`
	Cidade string `json:"cidade"`
	Pais   string `json:"pais"`
}

// MakeLocal seta os valores a partir da entidade
func (l *PostLocal) MakeLocal(local entity.Local) {
	l.ID = local.ID
	l.Cidade = local.Cidade
	l.Pais = local.Pais
}

// Local presenter Local
type Local struct {
	ID     int    `json:"id"`
	Cidade string `json:"cidade"`
	Pais   string `json:"pais"`
}

// AppendLocal adiciona presenter na lista
func AppendLocal(local entity.Local, locals []Local) []Local {
	l := &Local{}
	l.MakeLocal(local)
	locals = append(locals, *l)
	return locals
}

// MakeLocal seta os valores a partir da entidade
func (l *Local) MakeLocal(local entity.Local) {
	l.ID = local.ID
	l.Cidade = local.Cidade
	l.Pais = local.Pais
}
