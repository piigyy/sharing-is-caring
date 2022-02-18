package server

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/piigyy/sharing-is-caring/internal/gateway/model"
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
		if errors.Is(err, model.ErrUserNotFound) {
			presenter.ErrResponse(w, http.StatusUnauthorized, model.ErrUnAuthorized)
			return
		}

		presenter.UnknownErrResp(w, err)
		return
	}

	presenter.SuccessReponse(w, resp, http.StatusOK)
}

func (s *httpServer) GetUserDetail(w http.ResponseWriter, r *http.Request) {
	credentials := r.Context().Value("credentials").(map[string]string)
	user, err := s.authService.GetUserDetailByEmail(r.Context(), credentials["email"])
	if err != nil {
		if errors.Is(err, model.ErrUserNotFound) {
			presenter.ErrResponse(w, http.StatusNotFound, model.ErrUserNotFound)
			return
		}

		presenter.UnknownErrResp(w, err)
		return
	}

	presenter.SuccessReponse(w, user, http.StatusOK)
}

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
