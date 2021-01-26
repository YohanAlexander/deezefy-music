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
