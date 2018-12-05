package ticket

import (
	"github.com/sirupsen/logrus"
	"time"

	"github.com/google/uuid"
)

type TicketService interface {
	CreateTicket(ticket *Ticket) error
	FindTicketById(id string) (*Ticket, error)
	FindAllTickets() ([]*Ticket, error)
}

type ticketService struct {
	repo TicketRepository
}

func NewTicketService(repo TicketRepository) TicketService {
	return &ticketService{
		repo,
	}
}

func (s *ticketService) CreateTicket(ticket *Ticket) error {
	ticket.ID = uuid.New().String()
	ticket.Created = time.Now()
	ticket.Updated = time.Now()
	ticket.Status = "OPEN"

	if err := s.repo.Create(ticket); err != nil {
		logrus.WithField("error", err).Error("Error creating ticket")
		return err
	}

	logrus.WithField("id", ticket.ID).Info("Created new ticket")
	return nil
}

func (s *ticketService) FindTicketById(id string) (*Ticket, error) {
	ticket, err := s.repo.FindById(id)

	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err, "id": id}).Error("Error finding ticket")
		return nil, err
	}
	logrus.WithField("id", id).Info("Found ticket")
	return ticket, nil
}

func (s *ticketService) FindAllTickets() ([]*Ticket, error) {
	tickets, err := s.repo.FindAll()
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("Error finding all tickets")
		return nil, err
	}
	logrus.Info("Found all tickets")
	return tickets, nil
}
