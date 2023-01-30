// Package oauth manages the OAuth2 protocole with spotify.
package oauth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"

	"github.com/landru29/gospot/internal/app"
	"golang.org/x/oauth2"
)

// Client is the Auth client.
type Client struct {
	conf  *oauth2.Config
	debug bool
}

const (
	maxAgeCookieState = 60 * 10
	stateBitSize      = 16

	// CookieStateName is the cookie name for state.
	CookieStateName = "check-state"

	authURL = "https://accounts.spotify.com/authorize"

	tokenURL = "https://accounts.spotify.com/api/token" //nolint: gosec
)

// New creates a new Google provider.
func New(conf *app.Config) (*Client, error) {
	redirectURL, err := url.Parse(conf.APIBaseURL)
	if err != nil {
		return nil, fmt.Errorf("cannot parse apiBaseUrl in configuration file: %w", err)
	}

	redirectURL.Path = path.Join(redirectURL.Path, "callback")

	return &Client{
		conf: &oauth2.Config{
			RedirectURL:  redirectURL.String(),
			ClientID:     conf.ClientID,
			ClientSecret: conf.ClientSecret,
			Scopes:       []string{"user-read-private", "user-library-read"},
			Endpoint: oauth2.Endpoint{
				AuthURL:  authURL,
				TokenURL: tokenURL,
			},
		},
		debug: conf.Debug,
	}, nil
}

// SetStateOauthCookie generate a cookie with state to avoid CSRF attacking.
func (c *Client) SetStateOauthCookie(writer *http.ResponseWriter) string {
	b := make([]byte, stateBitSize)
	_, _ = rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	http.SetCookie(
		*writer,
		&http.Cookie{
			Name:   CookieStateName,
			Value:  state,
			MaxAge: maxAgeCookieState,
			Domain: "",
			Path:   "/callback",
		},
	)

	opts := []oauth2.AuthCodeOption{}
	if c.debug {
		// show_dialog forces the user to reconnect.
		opts = append(opts, oauth2.SetAuthURLParam("show_dialog", "true"))
	}

	return c.conf.AuthCodeURL(state, opts...)
}

// Exchange retrieves the code.
func (c *Client) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	return c.conf.Exchange(ctx, code)
}
