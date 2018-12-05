package ticket

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/gorilla/mux"
)

type TicketHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

type ticketHandler struct {
	ticketService TicketService
}

func NewTicketHandler(ticketService TicketService) TicketHandler {
	return &ticketHandler{
		ticketService,
	}
}

func (h *ticketHandler) Get(w http.ResponseWriter, r *http.Request) {
	tickets, err := h.ticketService.FindAllTickets()
	if err != nil {
		logrus.WithField("error", err).Error("Unable to find all tickets")
		http.Error(w, "Unable to find all tickets", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(tickets)
	if err != nil {
		logrus.WithField("error", err).Error("Error unmarshalling response")
		http.Error(w, "Unable to get ticket", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil{
		logrus.WithField("error", err).Error("Error writing response")
	}
}

func (h *ticketHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ticket, err := h.ticketService.FindTicketById(id)
	if err != nil {
		logrus.WithFields(logrus.Fields{"error":err, "id": id}).Error("Unable to find ticket")
		http.Error(w, "Unable to find ticket", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(ticket)
	if err != nil {
		logrus.WithField("error", err).Error("Error unmarshalling response")
		http.Error(w, "Unable to fetch tickets", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil{
		logrus.WithField("error", err).Error("Error writing response")
	}
}

func (h *ticketHandler) Create(w http.ResponseWriter, r *http.Request) {

	var ticket Ticket
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ticket); err != nil{
		logrus.Error("Unable to decode ticket")
		http.Error(w, "Bad format for ticket", http.StatusBadRequest)
		return
	}

	if err := h.ticketService.CreateTicket(&ticket); err != nil {
		logrus.WithField("error", err).Error("Unable to create ticket")
		http.Error(w, "Unable to create ticket", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(ticket)
	if err != nil {
		logrus.WithField("error", err).Error("Error unmarshalling response")
		http.Error(w, "Unable to create ticket", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(response); err != nil{
		logrus.WithField("error", err).Error("Error writing response")
	}

}
