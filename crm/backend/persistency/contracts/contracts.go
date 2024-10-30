package contracts

import "github.com/DanielGabay/crm/contracts"

type DalInterface interface {
	CreateContact(createContactRequest *contracts.CreateContactRequest) (*contracts.Contact, error)
	GetContacts(getContactRequest *contracts.GetContactsRequest) (*contracts.GetContactsResponse, error)
	UpdateContact(updateContactRequest *contracts.UpdateContactRequest) (*contracts.Contact, error)
	DeleteContact(id string) error
	CreateLead(createLeadRequest *contracts.CreateLeadRequest) (*contracts.Lead, error)
	GetLeads(getLeadsRequest *contracts.GetLeadsRequest) (*contracts.GetLeadsResponse, error)
	UpdateLead(updateLeadRequest *contracts.UpdateLeadRequest) (*contracts.Lead, error)
	DeleteLead(id string) error
	CreateCustomer(createCustomerRequest *contracts.CreateCustomerRequest) (*contracts.Customer, error)
	GetCustomers(getCustomersRequest *contracts.GetCustomersRequest) (*contracts.GetCustomersResponse, error)
	UpdateCustomer(updateCustomerRequest *contracts.UpdateCustomerRequest) (*contracts.Customer, error)
	DeleteCustomer(id string) error
	CreateCommunicationHistory(createCommunicationHistoryRequest *contracts.CreateCommunicationHistoryRequest) (*contracts.CommunicationHistory, error)
	GetCommunicationHistory(getCommunicationHistoryRequest *contracts.GetCommunicationHistoryRequest) (*contracts.GetCommunicationHistoryResponse, error)
	UpdateCommunicationHistory(updateCommunicationHistoryRequest *contracts.UpdateCommunicationHistoryRequest) (*contracts.CommunicationHistory, error)
	DeleteCommunicationHistory(id string) error
}
