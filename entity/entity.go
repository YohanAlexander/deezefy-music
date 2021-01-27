package entity

import "github.com/google/uuid"

// ID entity UUID
type ID = uuid.UUID

// NewID create a new entity UUID
func NewID() ID {
	return ID(uuid.New())
}

// StringToID convert a string to an entity UUID
func StringToID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}
