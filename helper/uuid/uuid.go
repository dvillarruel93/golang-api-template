package uuid

import "github.com/rs/xid"

// GenerateUUID generates a unique UUID based on the xid library
// Link to library docs https://github.com/rs/xid
func GenerateUUID() string {
	return xid.New().String()
}
