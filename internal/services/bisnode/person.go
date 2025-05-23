package bisnode

import (
	"bisnode/internal/models"
	"context"
	"fmt"
)

// PersonService handles person-related business logic
type PersonService struct {
	client *Client
}

// NewPersonService creates a new PersonService
func NewPersonService(client *Client) *PersonService {
	return &PersonService{
		client: client,
	}
}

// SearchByMobileNumber searches for a person by mobile number
func (s *PersonService) SearchByMobileNumber(ctx context.Context, mobileNumber string) (*models.Person, error) {
	if mobileNumber == "" {
		return nil, fmt.Errorf("mobile number cannot be empty")
	}

	// You might want to add validation for the mobile number format here

	person, err := s.client.SearchPerson(ctx, mobileNumber)
	if err != nil {
		return nil, fmt.Errorf("error searching person: %w", err)
	}

	return person, nil
}
