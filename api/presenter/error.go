package presenter

import "errors"

// Erro presenter
type Erro struct {
	Message    string `json:"erro"`
	StatusCode int    `json:"status"`
}

// ErrJSON json error
var ErrJSON = errors.New("JSON Error")

// ErrNotFound not found
var ErrNotFound = errors.New("Not Found")

// ErrUnexpected database error
var ErrUnexpected = errors.New("Unexpected Error")

// ErrInvalidEntity domain error
var ErrInvalidEntity = errors.New("InvalidEntity Error")
