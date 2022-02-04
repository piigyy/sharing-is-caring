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

func (s *httpServer) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var payload model.RegisterUserRequest

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Printf("error decoding register user payload: %v\n", err)
		if errors.Is(err, io.EOF) {
			presenter.ErrResponse(w, http.StatusBadRequest, errors.New("empty request body"))
			return
		}
		presenter.UnknownErrResp(w, err)
		return
	}

	log.Println("sending request to service")
	resp, err := s.authService.RegisterUser(r.Context(), payload)
	if err != nil {
		log.Printf("error from s.authService.RegisterUser: %v\n", err)
		if errors.Is(err, model.ErrUserDuplicated) {
			presenter.ErrResponse(w, http.StatusBadRequest, err)
			return
		}

		presenter.UnknownErrResp(w, err)
		return
	}

	log.Printf("success created new user")
	presenter.SuccessReponse(w, resp, http.StatusCreated)
}
