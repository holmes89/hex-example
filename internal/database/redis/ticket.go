package redis

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"hex-example/internal/ticket"
)

const ticketTable = "tickets"

type ticketRepository struct {
	connection *redis.Client
}

func NewRedisTicketRepository(connection *redis.Client) ticket.TicketRepository {
	return &ticketRepository{
		connection,
	}
}

func (r *ticketRepository) Create(ticket *ticket.Ticket) error {
	encoded, err := json.Marshal(ticket)

	if err != nil {
		logrus.Error("Unable to marshal ticket")
		return err
	}

	r.connection.HSet(ticketTable, ticket.ID, encoded) //Don't expire
	return nil
}

func (r *ticketRepository) FindById(id string) (*ticket.Ticket, error) {
	b, err := r.connection.HGet(ticketTable, id).Bytes()

	if err != nil {
		logrus.WithField("id", id).Error("Unable to fetch ticket")
		return nil, err
	}

	t := new(ticket.Ticket)
	err = json.Unmarshal(b, t)

	if err != nil {
		logrus.WithField("id", id).Error("Unable to unmarshal ticket")
		return nil, err
	}

	return t, nil
}

func (r *ticketRepository) FindAll() (tickets []*ticket.Ticket, err error) {
	ts := r.connection.HGetAll(ticketTable).Val()
	for key, value := range ts {
		t := new(ticket.Ticket)
		err = json.Unmarshal([]byte(value), t)

		if err != nil {
			logrus.WithField("id", key).Error("Unable to unmarshal ticket")
			return nil, err
		}

		t.ID = key
		tickets = append(tickets, t)
	}
	return tickets, nil
}
