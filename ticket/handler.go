package ticket

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
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
	tickets, _ := h.ticketService.FindAllTickets()

	response, _ := json.Marshal(tickets)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

func (h *ticketHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ticket, _ := h.ticketService.FindTicketById(id)

	response, _ := json.Marshal(ticket)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

func (h *ticketHandler) Create(w http.ResponseWriter, r *http.Request) {

	var ticket Ticket
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&ticket)
	_ = h.ticketService.CreateTicket(&ticket)

	response, _ := json.Marshal(ticket)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)

}