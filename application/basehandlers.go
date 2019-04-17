package application

import (
	"bytes"
	"net/http"

	"github.com/terrabot-tech/service-template/models"
)

// AddHeaders Добавляет заголовки ответу
func (app *Application) AddHeaders(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Token, Authorization")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		next(w, r)
	}
}

// HealthHandler check is alive
func (app *Application) HealthHandler(w http.ResponseWriter, r *http.Request) {
	sum := models.GetHashSum()
	if !bytes.Equal(sum, app.hashSum) {
		_ = app.logger.Log("msg", "New Configuration")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
