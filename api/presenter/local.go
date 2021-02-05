package presenter

// Local entidade Local
type Local struct {
	ID     int    `json:"id"`
	Cidade string `json:"cidade"`
	Pais   string `json:"pais"`
}
