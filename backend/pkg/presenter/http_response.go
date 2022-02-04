package presenter

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

func SuccessReponse(w http.ResponseWriter, data interface{}, httpCode int) {
	resp := Response{
		Success: true,
		Message: "success",
		Data:    data,
		Error:   nil,
	}
	respJSON, _ := json.Marshal(resp)
	w.WriteHeader(httpCode)
	w.Write(respJSON)
}

func ErrResponse(w http.ResponseWriter, httpCode int, err error) {
	resp := Response{
		Success: false,
		Message: err.Error(),
		Data:    nil,
		Error:   err,
	}

	respJSON, _ := json.Marshal(resp)
	w.WriteHeader(httpCode)
	w.Write(respJSON)
}

func UnknownErrResp(w http.ResponseWriter, err error) {
	resp := Response{
		Success: false,
		Data:    nil,
		Message: http.StatusText(http.StatusInternalServerError),
		Error:   err,
	}

	respJSON, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(respJSON)
}
