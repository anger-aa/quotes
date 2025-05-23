package response

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Response struct {
	Message string
	Data    interface{}
}

func JSON(w http.ResponseWriter, message string, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := Response{
		Message: message,
		Data:    data,
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return errors.New("something went wrong")
	}

	return nil
}
