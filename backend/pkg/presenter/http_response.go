package presenter

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Errors  []string    `json:"errors"`
}

func SuccessReponse(w http.ResponseWriter, data interface{}, httpCode int) {
	resp := Response{
		Success: true,
		Data:    data,
		Errors:  []string{},
	}
	respJSON, _ := json.Marshal(resp)
	w.WriteHeader(httpCode)
	w.Write(respJSON)
}

func ErrResponse(w http.ResponseWriter, httpCode int, errs []error) {
	resp := Response{
		Success: false,
		Data:    nil,
		Errors:  []string{},
	}

	for _, err := range errs {
		resp.Errors = append(resp.Errors, err.Error())
	}
	respJSON, _ := json.Marshal(resp)
	w.WriteHeader(httpCode)
	w.Write(respJSON)
}

func UnknownErrResp(w http.ResponseWriter) {
	resp := Response{
		Success: false,
		Data:    nil,
		Errors:  []string{"internal server error"},
	}

	respJSON, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(respJSON)
}
