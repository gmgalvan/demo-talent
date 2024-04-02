package services

import (
	"fmt"
	"time"
)

// GenerateUniqueID generates a new unique ID for an expense.
func GenerateUniqueID(service string) string {
	return fmt.Sprintf("%s_%d", service, time.Now().UnixNano())
}
