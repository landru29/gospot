package api

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/landru29/gospot/internal/oauth"
)

// processLogin is the handler for GET /login.
func (s *Server) processLogin(writer http.ResponseWriter, req *http.Request) {
	http.Redirect(writer, req, s.oauth.SetStateOauthCookie(&writer), http.StatusTemporaryRedirect)
}

// processCallback is the handler for GET /callback.
func (s *Server) processCallback(writer http.ResponseWriter, req *http.Request) {
	if err := s.checkState(req); err != nil {
		s.renderError(writer, fmt.Sprintf("state cookie not found: %s", err))

		return
	}

	if req.FormValue("error") != "" {
		s.renderError(writer, req.FormValue("error"))

		return
	}

	token, err := s.oauth.Exchange(req.Context(), req.FormValue("code"))
	if err != nil {
		s.renderError(writer, fmt.Sprintf("could not exchange token: %s", err))

		return
	}

	http.Redirect(
		writer,
		req,
		fmt.Sprintf("%s?token=%s", s.conf.Redirect, url.QueryEscape(token.AccessToken)),
		http.StatusTemporaryRedirect,
	)
}

func (s *Server) checkState(req *http.Request) error {
	oauthState, err := req.Cookie(oauth.CookieStateName)
	if err != nil {
		s.log.WithError(err).Error("cookie not found")

		return err
	}

	if req.FormValue("state") != oauthState.Value {
		s.log.Error("invalid oauth state")

		return os.ErrInvalid
	}

	return nil
}

func (s *Server) renderError(writer http.ResponseWriter, msg string) {
	s.log.WithError(errors.New(msg)).Error("error occurred")

	_ = s.templates.Execute(writer, map[string]string{
		"description": msg,
	})
}
