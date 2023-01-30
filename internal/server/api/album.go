package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) listAlbums(writer http.ResponseWriter, req *http.Request) {
	bearer, err := bearerFromRequest(req)
	if err != nil {
		writer.WriteHeader(http.StatusForbidden)

		return
	}

	log := s.log.WithField("bearer", bearer)

	albums, err := s.catalog.Albums(req.Context(), bearer)
	if err != nil {
		log.WithError(err).Error("cannot get albums")

		writer.WriteHeader(http.StatusNotFound)

		return
	}

	if err := json.NewEncoder(writer).Encode(albums); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	writer.WriteHeader(http.StatusOK)
}
