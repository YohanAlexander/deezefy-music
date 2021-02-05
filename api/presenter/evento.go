package presenter

// Evento entidade Evento
type Evento struct {
	Usuario Usuario `json:"usuario"`
	ID      int     `json:"id"`
	Nome    string  `json:"nome"`
	Data    string  `json:"data"`
}
