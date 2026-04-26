package utils

import (
	"fmt"

	"github.com/google/uuid"
)

// ParseUUIDs converts a slice of string UUIDs to uuid.UUID slice.
func ParseUUIDs(ids []string) ([]uuid.UUID, error) {
	result := make([]uuid.UUID, len(ids))
	for i, id := range ids {
		parsed, err := uuid.Parse(id)
		if err != nil {
			return nil, fmt.Errorf("invalid UUID '%s': %w", id, err)
		}
		result[i] = parsed
	}
	return result, nil
}
