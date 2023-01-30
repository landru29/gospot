// Package server describes a generic server.
package server

import "context"

const (
	// ReadHeaderTimeoutSeconds is a HTTP timeout.
	ReadHeaderTimeoutSeconds = 2
)

// Router is an http router.
type Router interface {
	Start(ctx context.Context)
}
