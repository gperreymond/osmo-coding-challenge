package utils

import uuid "github.com/satori/go.uuid"

// UUID retunrs an unique id basend on ulid algorithm
func UUID() string {
	id := uuid.NewV4()
	return id.String()
}
