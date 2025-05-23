package response

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type Response struct {
	Message string
	Data    interface{}
}

func JSON(w http.ResponseWriter, message string, data interface{}) error {
	var buf bytes.Buffer

	response := Response{
		Message: message,
		Data:    data,
	}

	err := json.NewEncoder(&buf).Encode(response)
	if err != nil {
		return errors.New("something went wrong")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(buf.Bytes())
	if err != nil {
		return errors.New("something went wrong")
	}

	return nil
}
