package ticket

import "time"

type Ticket struct {
	ID          string    `json:"id" db:"id"`
	Creator     string    `json:"creator" db:"creator"`
	Assigned    string    `json:"assigned" db:"assigned"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Status      string    `json:"status" db:"status"`
	Points      int       `json:"points" db:"points"`
	Created     time.Time `json:"created" db:"created"`
	Updated     time.Time `json:"updated" db:"updated"`
	Deleted     time.Time `json:"deleted" db:"deleted"`
}
