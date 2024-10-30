package mocks

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/DanielGabay/crm/contracts"
	"github.com/google/uuid"
)

type Mocks struct {
	contacts               []*contracts.Contact
	leads                  []*contracts.Lead
	customers              []*contracts.Customer
	communicationHistories []*contracts.CommunicationHistory
}

func New() *Mocks {
	return &Mocks{
		contacts:               []*contracts.Contact{},
		leads:                  []*contracts.Lead{},
		customers:              []*contracts.Customer{},
		communicationHistories: []*contracts.CommunicationHistory{},
	}
}

func (m *Mocks) CreateContact(createContactRequest *contracts.CreateContactRequest) (*contracts.Contact, error) {
	contact := &contracts.Contact{
		Id:        uuid.New(),
		Name:      createContactRequest.Name,
		Email:     createContactRequest.Email,
		Phone:     createContactRequest.Phone,
		Address:   createContactRequest.Address,
		CreatedAt: time.Now(),
	}
	m.contacts = append(m.contacts, contact)
	return contact, nil
}

func (m *Mocks) GetContacts(getContactRequest *contracts.GetContactsRequest) (*contracts.GetContactsResponse, error) {
	contacts := []*contracts.Contact{}
	for _, contact := range m.contacts {
		if (getContactRequest.Filter.ContactId == uuid.Nil || contact.Id == getContactRequest.Filter.ContactId) &&
			isSubstringIncluded(contact.Name, getContactRequest.Filter.Name) &&
			isSubstringIncluded(contact.Email, getContactRequest.Filter.Email) &&
			isSubstringIncluded(contact.Phone, getContactRequest.Filter.Phone) &&
			isSubstringIncluded(contact.Address, getContactRequest.Filter.Address) &&
			isDateTimeBetween(getContactRequest.FromDateTime, getContactRequest.ToDateTime, contact.CreatedAt) {
			contacts = append(contacts, contact)
		}
	}
	return &contracts.GetContactsResponse{
		Contacts:   contacts,
		TotalItems: len(contacts),
	}, nil
}

func (m *Mocks) UpdateContact(updateContactRequest *contracts.UpdateContactRequest) (*contracts.Contact, error) {
	contact, err := getElementByFieldValue(m.contacts, "Id", updateContactRequest.Id)
	if err != nil {
		return nil, err
	}
	contact.(*contracts.Contact).Name = updateContactRequest.Name
	contact.(*contracts.Contact).Email = updateContactRequest.Email
	contact.(*contracts.Contact).Phone = updateContactRequest.Phone
	contact.(*contracts.Contact).Address = updateContactRequest.Address
	return contact.(*contracts.Contact), nil
}

func (m *Mocks) DeleteContact(id string) error {
	for i := 0; i < len(m.contacts); i++ {
		if m.contacts[i].Id.String() == id {
			m.contacts = append(m.contacts[:i], m.contacts[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("contact with id %s not found", id)
}

func (m *Mocks) CreateLead(createLeadRequest *contracts.CreateLeadRequest) (*contracts.Lead, error) {
	lead := &contracts.Lead{
		Id:        uuid.New(),
		ContactId: createLeadRequest.ContactId,
		Source:    createLeadRequest.Source,
		Status:    contracts.New,
		CreatedAt: time.Now(),
	}
	m.leads = append(m.leads, lead)
	return lead, nil
}

func (m *Mocks) GetLeads(getLeadsRequest *contracts.GetLeadsRequest) (*contracts.GetLeadsResponse, error) {
	leads := []*contracts.Lead{}
	for _, lead := range m.leads {
		if (getLeadsRequest.Filter.LeadId == uuid.Nil || lead.Id == getLeadsRequest.Filter.LeadId) &&
			(getLeadsRequest.Filter.ContactId == uuid.Nil || lead.ContactId == getLeadsRequest.Filter.ContactId) &&
			(len(getLeadsRequest.Filter.Source) == 0 || reflect.DeepEqual(lead.Source, getLeadsRequest.Filter.Source)) &&
			isSubstringIncluded(lead.Status, getLeadsRequest.Filter.Status) &&
			isDateTimeBetween(getLeadsRequest.FromDateTime, getLeadsRequest.ToDateTime, lead.CreatedAt) {
			leads = append(leads, lead)
		}
	}
	return &contracts.GetLeadsResponse{
		Leads:      leads,
		TotalItems: len(leads),
	}, nil
}

func (m *Mocks) UpdateLead(updateLeadRequest *contracts.UpdateLeadRequest) (*contracts.Lead, error) {
	lead, err := getElementByFieldValue(m.leads, "Id", updateLeadRequest.Id)
	if err != nil {
		return nil, err
	}
	lead.(*contracts.Lead).Source = updateLeadRequest.Source
	lead.(*contracts.Lead).Status = updateLeadRequest.Status
	return lead.(*contracts.Lead), nil
}

func (m *Mocks) DeleteLead(id string) error {
	for i := 0; i < len(m.leads); i++ {
		if m.leads[i].Id.String() == id {
			m.leads = append(m.leads[:i], m.leads[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("lead with id %s not found", id)
}

func (m *Mocks) CreateCustomer(createCustomerRequest *contracts.CreateCustomerRequest) (*contracts.Customer, error) {
	customer := &contracts.Customer{
		Id:        uuid.New(),
		ContactId: createCustomerRequest.ContactId,
		Source:    createCustomerRequest.Source,
		CreatedAt: time.Now(),
	}
	m.customers = append(m.customers, customer)
	return customer, nil
}

func (m *Mocks) GetCustomers(getCustomersRequest *contracts.GetCustomersRequest) (*contracts.GetCustomersResponse, error) {
	customers := []*contracts.Customer{}
	for _, customer := range m.customers {
		if (getCustomersRequest.Filter.CustomerId == uuid.Nil || customer.Id == getCustomersRequest.Filter.CustomerId) &&
			(getCustomersRequest.Filter.ContactId == uuid.Nil || customer.ContactId == getCustomersRequest.Filter.ContactId) &&
			(len(getCustomersRequest.Filter.Source) == 0 || reflect.DeepEqual(customer.Source, getCustomersRequest.Filter.Source)) &&
			isDateTimeBetween(getCustomersRequest.FromDateTime, getCustomersRequest.ToDateTime, customer.CreatedAt) {
			customers = append(customers, customer)
		}
	}
	return &contracts.GetCustomersResponse{
		Customers:  customers,
		TotalItems: len(customers),
	}, nil
}

func (m *Mocks) UpdateCustomer(updateCustomerRequest *contracts.UpdateCustomerRequest) (*contracts.Customer, error) {
	customer, err := getElementByFieldValue(m.customers, "Id", updateCustomerRequest.Id)
	if err != nil {
		return nil, err
	}
	customer.(*contracts.Customer).Source = updateCustomerRequest.Source
	return customer.(*contracts.Customer), nil
}

func (m *Mocks) DeleteCustomer(id string) error {
	for i := 0; i < len(m.customers); i++ {
		if m.customers[i].Id.String() == id {
			m.customers = append(m.customers[:i], m.customers[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("customer with id %s not found", id)
}

func (m *Mocks) CreateCommunicationHistory(createCommunicationHistoryRequest *contracts.CreateCommunicationHistoryRequest) (*contracts.CommunicationHistory, error) {
	communicationHistory := &contracts.CommunicationHistory{
		Id:                  uuid.New(),
		ContactId:           createCommunicationHistoryRequest.ContactId,
		CreatedAt:           time.Now(),
		Title:               createCommunicationHistoryRequest.Title,
		Description:         createCommunicationHistoryRequest.Description,
		Note:                createCommunicationHistoryRequest.Note,
		TypeOfCommunication: createCommunicationHistoryRequest.TypeOfCommunication,
		Outcome:             createCommunicationHistoryRequest.Outcome,
	}
	m.communicationHistories = append(m.communicationHistories, communicationHistory)
	return communicationHistory, nil
}

func (m *Mocks) GetCommunicationHistory(getCommunicationHistoryRequest *contracts.GetCommunicationHistoryRequest) (*contracts.GetCommunicationHistoryResponse, error) {
	communicationHistories := []*contracts.CommunicationHistory{}
	for _, communicationHistory := range m.communicationHistories {
		if (getCommunicationHistoryRequest.Filter.CommunicationHistoryId == uuid.Nil || communicationHistory.Id == getCommunicationHistoryRequest.Filter.CommunicationHistoryId) &&
			(getCommunicationHistoryRequest.Filter.ContactId == uuid.Nil || communicationHistory.ContactId == getCommunicationHistoryRequest.Filter.ContactId) &&
			isSubstringIncluded(communicationHistory.Title, getCommunicationHistoryRequest.Filter.Title) &&
			isSubstringIncluded(communicationHistory.TypeOfCommunication, getCommunicationHistoryRequest.Filter.TypeOfCommunication) &&
			isSubstringIncluded(communicationHistory.Outcome, getCommunicationHistoryRequest.Filter.Outcome) &&
			isDateTimeBetween(getCommunicationHistoryRequest.FromDateTime, getCommunicationHistoryRequest.ToDateTime, communicationHistory.CreatedAt) {
			communicationHistories = append(communicationHistories, communicationHistory)
		}
	}
	return &contracts.GetCommunicationHistoryResponse{
		CommunicationHistories: communicationHistories,
		TotalItems:             len(communicationHistories),
	}, nil
}

func (m *Mocks) UpdateCommunicationHistory(updateCommunicationHistoryRequest *contracts.UpdateCommunicationHistoryRequest) (*contracts.CommunicationHistory, error) {
	communicationHistory, err := getElementByFieldValue(m.communicationHistories, "Id", updateCommunicationHistoryRequest.Id)
	if err != nil {
		return nil, err
	}
	communicationHistory.(*contracts.CommunicationHistory).Title = updateCommunicationHistoryRequest.Title
	communicationHistory.(*contracts.CommunicationHistory).Description = updateCommunicationHistoryRequest.Description
	communicationHistory.(*contracts.CommunicationHistory).Note = updateCommunicationHistoryRequest.Note
	communicationHistory.(*contracts.CommunicationHistory).TypeOfCommunication = updateCommunicationHistoryRequest.TypeOfCommunication
	communicationHistory.(*contracts.CommunicationHistory).Outcome = updateCommunicationHistoryRequest.Outcome
	return communicationHistory.(*contracts.CommunicationHistory), nil
}

func (m *Mocks) DeleteCommunicationHistory(id string) error {
	for i := 0; i < len(m.communicationHistories); i++ {
		if m.communicationHistories[i].Id.String() == id {
			m.communicationHistories = append(m.communicationHistories[:i], m.communicationHistories[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("communication history with id %s not found", id)
}

func getElementByFieldValue(arr interface{}, fieldName string, value interface{}) (interface{}, error) {
	for i := 0; i < reflect.ValueOf(arr).Len(); i++ {
		// Get the field value of the current element
		fieldValue := reflect.ValueOf(arr).Index(i).FieldByName(fieldName)

		// Check if the field value matches the given value
		if fieldValue.IsValid() && reflect.DeepEqual(fieldValue.Interface(), value) {
			// Return the element from the array
			return reflect.ValueOf(arr).Index(i).Interface(), nil
		}
	}

	// If no match is found, return nil
	return nil, fmt.Errorf("element with %s %v not found", fieldName, value)
}

func isSubstringIncluded(str, substr string) bool {
	if substr == "" {
		return true
	}
	return strings.Contains(str, substr)
}

func isDateTimeBetween(fromDateTime, toDateTime, dateTime time.Time) bool {
	// If both fromDateTime and toDateTime are empty, return true
	if fromDateTime.IsZero() && toDateTime.IsZero() {
		return true
	}

	// Check if dateTime is between fromDateTime and toDateTime
	if !fromDateTime.IsZero() && dateTime.Before(fromDateTime) {
		return false
	}
	if !toDateTime.IsZero() && dateTime.After(toDateTime) {
		return false
	}

	return true
}
