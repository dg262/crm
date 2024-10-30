package contracts

import (
	"time"

	"github.com/google/uuid"
)

type Contact struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

type Lead struct {
	Id        uuid.UUID `json:"id"`
	ContactId uuid.UUID `json:"contact_id"`
	Source    []string  `json:"source"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type Customer struct {
	Id        uuid.UUID `json:"id"`
	ContactId uuid.UUID `json:"contact_id"`
	Source    []string  `json:"source"`
	CreatedAt time.Time `json:"created_at"`
}

type CommunicationHistory struct {
	Id                  uuid.UUID `json:"id"`
	ContactId           uuid.UUID `json:"contact_id"`
	CreatedAt           time.Time `json:"created_at"`
	Title               string    `json:"title"`
	Description         string    `json:"description"`
	Note                string    `json:"note"`
	TypeOfCommunication string    `json:"type_of_communication"`
	Outcome             string    `json:"outcome"`
}
