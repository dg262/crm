package contracts

import (
	"time"

	"github.com/google/uuid"
)

type CreateContactRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type GetContactsRequest struct {
	Filter       *ContactFilter `json:"filter"`
	FromDateTime time.Time      `json:"from_date_time"`
	ToDateTime   time.Time      `json:"to_date_time"`
}

type UpdateContactRequest struct {
	Id      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Phone   string    `json:"phone"`
	Address string    `json:"address"`
}

type GetContactsResponse struct {
	Contacts   []*Contact `json:"contacts"`
	TotalItems int        `json:"total_items"`
}

type CreateLeadRequest struct {
	ContactId uuid.UUID `json:"contact_id"`
	Source    []string  `json:"source"`
}

type GetLeadsRequest struct {
	Filter       *LeadFilter `json:"filter"`
	FromDateTime time.Time   `json:"from_date_time"`
	ToDateTime   time.Time   `json:"to_date_time"`
}

type UpdateLeadRequest struct {
	Id     uuid.UUID `json:"id"`
	Source []string  `json:"source"`
	Status string    `json:"status"`
}

type GetLeadsResponse struct {
	Leads      []*Lead `json:"leads"`
	TotalItems int     `json:"total_items"`
}

type CreateCustomerRequest struct {
	ContactId uuid.UUID `json:"contact_id"`
	Source    []string  `json:"source"`
}

type GetCustomersRequest struct {
	Filter       *CustomerFilter `json:"filter"`
	FromDateTime time.Time       `json:"from_date_time"`
	ToDateTime   time.Time       `json:"to_date_time"`
}

type UpdateCustomerRequest struct {
	Id     uuid.UUID `json:"id"`
	Source []string  `json:"source"`
}

type GetCustomersResponse struct {
	Customers  []*Customer `json:"customers"`
	TotalItems int         `json:"total_items"`
}

type CreateCommunicationHistoryRequest struct {
	ContactId           uuid.UUID `json:"contact_id"`
	Title               string    `json:"title"`
	Description         string    `json:"description"`
	Note                string    `json:"note"`
	TypeOfCommunication string    `json:"type_of_communication"`
	Outcome             string    `json:"outcome"`
}

type GetCommunicationHistoryRequest struct {
	Filter       *CommunicationHistoryFilter `json:"filter"`
	FromDateTime time.Time                   `json:"from_date_time"`
	ToDateTime   time.Time                   `json:"to_date_time"`
}

type UpdateCommunicationHistoryRequest struct {
	Id                  uuid.UUID `json:"id"`
	Title               string    `json:"title"`
	Description         string    `json:"description"`
	Note                string    `json:"note"`
	TypeOfCommunication string    `json:"type_of_communication"`
	Outcome             string    `json:"outcome"`
}

type GetCommunicationHistoryResponse struct {
	CommunicationHistories []*CommunicationHistory `json:"communication_histories"`
	TotalItems             int                     `json:"total_items"`
}
