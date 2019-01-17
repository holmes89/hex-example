package ticket

type TicketRepository interface {
	Create(ticket *Ticket) error
	FindById(id string) (*Ticket, error)
	FindAll() ([]*Ticket, error)
}
