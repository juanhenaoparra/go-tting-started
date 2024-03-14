package shared

import (
	"github.com/oklog/ulid/v2"
)

// NewID generates a new ID with the given product ID and resource ID prefix
func NewID(prefix string) string {
	return prefix + "_" + ulid.Make().String()
}
