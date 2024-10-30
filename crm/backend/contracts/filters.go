package contracts

import "github.com/google/uuid"

type ContactFilter struct {
	ContactId uuid.UUID `json:"contact_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
}

type LeadFilter struct {
	LeadId    uuid.UUID `json:"lead_id"`
	ContactId uuid.UUID `json:"contact_id"`
	Source    []string  `json:"source"`
	Status    string    `json:"status"`
}

type CustomerFilter struct {
	CustomerId uuid.UUID `json:"customer_id"`
	ContactId  uuid.UUID `json:"contact_id"`
	Source     []string  `json:"source"`
}

type CommunicationHistoryFilter struct {
	CommunicationHistoryId uuid.UUID `json:"communication_history_id"`
	ContactId              uuid.UUID `json:"contact_id"`
	Title                  string    `json:"title"`
	TypeOfCommunication    string    `json:"type_of_communication"`
	Outcome                string    `json:"outcome"`
}
