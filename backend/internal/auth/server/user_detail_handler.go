package server

import (
	"errors"
	"net/http"

	"github.com/piigyy/sharing-is-caring/internal/auth/model"
	"github.com/piigyy/sharing-is-caring/pkg/presenter"
)

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
