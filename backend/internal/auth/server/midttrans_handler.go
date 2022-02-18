package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/piigyy/sharing-is-caring/internal/auth/model"
	"github.com/piigyy/sharing-is-caring/pkg/presenter"
)

func (s *httpServer) MidtransWebHookHandler(w http.ResponseWriter, r *http.Request) {
	var data model.MidtransWebHook

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		log.Printf("error decoding webhook data: %v", err)
	}

	fmt.Printf("data: %+v\n", data)
	presenter.SuccessReponse(w, "success", http.StatusOK)
}
