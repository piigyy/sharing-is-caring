package server

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/piigyy/sharing-is-caring/internal/auth/model"
	"github.com/piigyy/sharing-is-caring/pkg/presenter"
)

func (s *httpServer) Login(w http.ResponseWriter, r *http.Request) {
	var payload model.LoginRequest

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Printf("error decoding login payload: %v\n", err)
		if errors.Is(err, io.EOF) {
			presenter.ErrResponse(w, http.StatusBadRequest, errors.New("empty request body"))
			return
		}
		presenter.UnknownErrResp(w, err)
		return
	}

	resp, err := s.authService.Login(r.Context(), payload)
	if err != nil {
		if errors.Is(err, model.ErrUnAuthorized) {
			presenter.ErrResponse(w, http.StatusUnauthorized, model.ErrUnAuthorized)
			return
		}

		presenter.UnknownErrResp(w, err)
		return
	}

	presenter.SuccessReponse(w, resp, http.StatusOK)
}
