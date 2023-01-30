// Package spotify manages the spotify catalog.
package spotify

import (
	"net/url"
	"path"

	"github.com/landru29/gospot/internal/music"
)

// Client is a spotify client.
type Client struct {
	baseURL url.URL
	user    *music.User
}

// New creates a new client.
func New(baseURL string) (*Client, error) {
	spotifyURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		baseURL: *spotifyURL,
	}, nil
}

func (c *Client) urlWith(dir string) *url.URL {
	return &url.URL{
		Scheme:      c.baseURL.Scheme,
		Opaque:      c.baseURL.Opaque,
		User:        c.baseURL.User,
		Host:        c.baseURL.Host,
		Path:        path.Join(c.baseURL.Path, dir),
		RawPath:     c.baseURL.RawPath,
		OmitHost:    c.baseURL.OmitHost,
		ForceQuery:  c.baseURL.ForceQuery,
		RawQuery:    c.baseURL.RawQuery,
		Fragment:    c.baseURL.Fragment,
		RawFragment: c.baseURL.RawFragment,
	}
}
