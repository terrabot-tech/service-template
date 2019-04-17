package application

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// NotificationsHandler основной обработчик /notifications
func (app *Application) NotificationsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.getNotifications(w, r)
	case http.MethodPut:
		app.setNotificationWatched(w, r)
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (app *Application) getNotifications(w http.ResponseWriter, r *http.Request) {
	paramUserID := r.URL.Query().Get("userId")
	if len(paramUserID) <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userID, err := strconv.ParseUint(paramUserID, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := app.svc.Repository.GetNotificationsByUserID(userID)
	if err != nil {
		err = app.logger.Log("err", err)
	}
	_ = json.NewEncoder(w).Encode(res)
	return
}

func (app *Application) setNotificationWatched(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	notificationID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	paramUserID := r.URL.Query().Get("userId")
	if len(paramUserID) <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseUint(paramUserID, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	watched, err:= strconv.Atoi(string(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = app.svc.Repository.SetNotificationWatched(notificationID, userID, watched)
	if err != nil {
		_ = app.logger.Log("err", err)
	}
}