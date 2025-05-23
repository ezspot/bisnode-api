package models

import "time"

// Person represents a person in the Bisnode API
type Person struct {
	ID                string    `json:"id,omitempty"`
	FirstName         string    `json:"firstName,omitempty"`
	MiddleName        string    `json:"middleName,omitempty"`
	LastName          string    `json:"lastName,omitempty"`
	DateOfBirth       string    `json:"dateOfBirth,omitempty"`
	Gender            string    `json:"gender,omitempty"`
	NationalID        string    `json:"nationalId,omitempty"`
	MobileNumber      string    `json:"mobileNumber,omitempty"`
	Email             string    `json:"email,omitempty"`
	RegistrationDate  time.Time `json:"registrationDate,omitempty"`
	LastUpdated       time.Time `json:"lastUpdated,omitempty"`
	Addresses         []Address `json:"addresses,omitempty"`
	PhoneNumbers      []Phone   `json:"phoneNumbers,omitempty"`
	EmailAddresses    []Email   `json:"emailAddresses,omitempty"`
	EmploymentDetails []string  `json:"employmentDetails,omitempty"`
}

// Address represents a person's address
type Address struct {
	StreetAddress string `json:"streetAddress,omitempty"`
	PostalCode   string `json:"postalCode,omitempty"`
	City         string `json:"city,omitempty"`
	Country      string `json:"country,omitempty"`
	AddressType  string `json:"addressType,omitempty"` // e.g., HOME, WORK
	IsCurrent    bool   `json:"isCurrent,omitempty"`
}

// Phone represents a phone number
type Phone struct {
	Number     string `json:"number,omitempty"`
	Type       string `json:"type,omitempty"` // e.g., MOBILE, HOME, WORK
	IsVerified bool   `json:"isVerified,omitempty"`
}

// Email represents an email address
type Email struct {
	Address    string `json:"address,omitempty"`
	Type       string `json:"type,omitempty"` // e.g., PERSONAL, WORK
	IsVerified bool   `json:"isVerified,omitempty"`
}

// SearchRequest represents a person search request
type SearchRequest struct {
	MobileNumber string `json:"mobileNumber" validate:"required"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}
