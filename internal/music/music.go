// Package music is the music manager.
package music

import "context"

// Album is a music album.
type Album struct {
	Label  string  `json:"label"`
	Name   string  `json:"name"`
	Images []Image `json:"images"`
}

type Image struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	URL    string `json:"url"`
}

// User is a user.
type User struct {
	ID string `json:"id"`
}

// Cataloger is the catalog manager.
type Cataloger interface {
	Albums(ctx context.Context, token string) ([]Album, error)
}
