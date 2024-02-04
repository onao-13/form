package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

var logger = log.Default()

// Create обертка для ответа клиенту с кодом 201
func Create(w http.ResponseWriter, msg string) {
	if len(msg) == 0 {
		w.WriteHeader(http.StatusCreated)
		return
	}
	data := map[string]interface{}{
		"message": msg,
	}
	handle(w, http.StatusCreated, data)
}

// InternalServerError обертка для ответа клиенту с кодом 500
func InternalServerError(w http.ResponseWriter, msg string) {
	data := map[string]interface{}{
		"error": msg,
	}
	handle(w, http.StatusInternalServerError, data)
}

// BadRequest обертка для ответа клиенту с кодом 400
func BadRequest(w http.ResponseWriter, msg string) {
	data := map[string]interface{}{
		"error": msg,
	}
	handle(w, http.StatusBadRequest, data)
}

// handle ответ клиенту
func handle(w http.ResponseWriter, status int, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		logger.Println("Ошибка создания JSON ответа: ", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(b)
}
