package presenter

import "errors"

// Erro presenter
type Erro struct {
	Message    string `json:"erro"`
	StatusCode int    `json:"status"`
}

// ErrJSON not found
var ErrJSON = errors.New("JSON error")

// ErrNotFound not found
var ErrNotFound = errors.New("Not Found")

// ErrUnexpected database error
var ErrUnexpected = errors.New("Unexpected error")
