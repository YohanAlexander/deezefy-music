package entity

import "errors"

// ErrNotFound not found
var ErrNotFound = errors.New("Not Found")

// ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("Invalid Entity")

// ErrCannotBeDeleted cannot be deleted
var ErrCannotBeDeleted = errors.New("Cannot Be Deleted")

// ErrPhoneRegistered cannot register
var ErrPhoneRegistered = errors.New("Phone already registered")

// ErrArtistaRegistered cannot register
var ErrArtistaRegistered = errors.New("Artista already registered")

// ErrOuvinteRegistered cannot register
var ErrOuvinteRegistered = errors.New("Ouvinte already registered")
