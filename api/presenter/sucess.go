package presenter

// Sucesso presenter
type Sucesso struct {
	Result     interface{} `json:"resultado"`
	StatusCode int         `json:"status"`
}

// SucessDelete delete sucess
var SucessDelete = "Delete Sucessful"

// Login presenter
type Login struct {
	Artista interface{} `json:"artista"`
	Ouvinte interface{} `json:"ouvinte"`
}
