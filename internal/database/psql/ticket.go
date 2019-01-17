package psql

import (
	"database/sql"
	"hex-example/internal/ticket"
	"log"

	_ "github.com/lib/pq"
)

type ticketRepository struct {
	db *sql.DB
}

func NewPostgresTicketRepository(db *sql.DB) ticket.TicketRepository {
	return &ticketRepository{
		db,
	}
}

func (r *ticketRepository) Create(ticket *ticket.Ticket) error {
	r.db.QueryRow("INSERT INTO tickets(creator, assigned, title, description, status, points, created, updated) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		ticket.Creator, ticket.Assigned, ticket.Title, ticket.Description, ticket.Status, ticket.Points, ticket.Created, ticket.Updated).Scan(&ticket.ID)
	return nil
}

func (r *ticketRepository) FindById(id string) (*ticket.Ticket, error) {
	ticket := new(ticket.Ticket)
	err := r.db.QueryRow("SELECT id, creator, assigned, title, description, status, points, created, updated FROM tickets where id=$1", id).Scan(&ticket.ID, &ticket.Creator, &ticket.Assigned, &ticket.Title, &ticket.Description, &ticket.Status, &ticket.Points, &ticket.Created, &ticket.Updated)
	if err != nil {
		panic(err)
	}
	return ticket, nil
}

func (r *ticketRepository) FindAll() (tickets []*ticket.Ticket, err error) {
	rows, err := r.db.Query("SELECT id, creator, assigned, title, description, status, points, created, updated FROM tickets")
	defer rows.Close()

	for rows.Next() {
		ticket := new(ticket.Ticket)
		if err = rows.Scan(&ticket.ID, &ticket.Creator, &ticket.Assigned, &ticket.Title, &ticket.Description, &ticket.Status, &ticket.Points, &ticket.Created, &ticket.Updated); err != nil {
			log.Print(err)
			return nil, err
		}

		tickets = append(tickets, ticket)

	}
	return tickets, nil
}
