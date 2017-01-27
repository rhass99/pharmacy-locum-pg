package api

import (
	"github.com/satori/go.uuid"
)

func UUIDG() (uuidg string) {
	return uuid.NewV1().String()
}
