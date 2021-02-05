package presenter

import "github.com/yohanalexander/deezefy-music/entity"

// Local entidade Local
type Local struct {
	ID     int    `json:"id"`
	Cidade string `json:"cidade"`
	Pais   string `json:"pais"`
}

// AppendLocal adiciona presenter na lista
func AppendLocal(local entity.Local, locals []*Local) []*Local {
	l := &Local{}
	l.GetLocal(local)
	locals = append(locals, l)
	return locals
}

// GetLocal seta os valores a partir da entidade
func (l *Local) GetLocal(local entity.Local) {
	l.ID = local.ID
	l.Cidade = local.Cidade
	l.Pais = local.Pais
}
