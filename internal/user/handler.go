package user

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UserHandler interface{
	CreateAccount(w http.ResponseWriter, r *http.Request)
	GetToken(w http.ResponseWriter, r *http.Request)
	//TODO Refresh Token
}

type userHandler struct {
	service UserService
}

func NewUserHandler(service UserService) UserHandler {
	return &userHandler{
		service,
	}
}

func (h *userHandler) CreateAccount(w http.ResponseWriter, r *http.Request){
	var account Account
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&account); err != nil{
		logrus.Error("Unable to decode account")
		http.Error(w, "Bad format for account", http.StatusBadRequest)
		return
	}
	if err := h.service.CreateAccount(&account); err != nil {
		logrus.WithField("error", err).Error("Unable to create account")
		http.Error(w, "Unable to create account", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(account)
	if err != nil {
		logrus.WithField("error", err).Error("Error unmarshalling response")
		http.Error(w, "Unable to create account", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(response); err != nil{
		logrus.WithField("error", err).Error("Error writing response")
	}
}

func (h *userHandler) GetToken(w http.ResponseWriter, r *http.Request){
	username, password, _ := r.BasicAuth()
	if username == "" || password == "" {
		http.Error(w, "No credentials provided", http.StatusUnauthorized)
		return
	}

	login, err := h.service.Login(username, password)
	if err != nil{
		logrus.WithFields(logrus.Fields{
			"error": err,
			"username": username,
		}).Error("Error generating token")
		http.Error(w, "Invalid login", http.StatusUnauthorized)
		return
	}

	response, err := json.Marshal(login)
	if err != nil {
		logrus.WithField("error", err).Error("Error unmarshalling response")
		http.Error(w, "Unable to login", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil{
		logrus.WithField("error", err).Error("Error writing response")
	}
}