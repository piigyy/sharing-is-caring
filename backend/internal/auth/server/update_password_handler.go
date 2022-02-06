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

func (s *httpServer) UpdatePasword(w http.ResponseWriter, r *http.Request) {
	var payload model.UpdatePasswordRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Printf("error decoding update user password payload: %v\n", err)
		if errors.Is(err, io.EOF) {
			presenter.ErrResponse(w, http.StatusBadRequest, errors.New("empty request body"))
			return
		}
		presenter.UnknownErrResp(w, err)
		return
	}

	credentials := r.Context().Value("credentials").(map[string]string)
	payload.Email = credentials["email"]

	err := s.authService.UpdatePassword(r.Context(), payload)
	if err != nil {
		if errors.Is(err, model.ErrUnAuthorized) {
			presenter.ErrResponse(w, http.StatusUnauthorized, err)
			return
		}

		log.Println("unknwon error handling update user password")
		presenter.UnknownErrResp(w, err)
		return
	}

	presenter.SuccessReponse(w, "update password success", http.StatusOK)
}
