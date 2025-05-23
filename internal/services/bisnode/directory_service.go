package bisnode

import (
	"bisnode/internal/models"
	"context"
	"fmt"
)

// DirectoryService handles directory search operations
type DirectoryService struct {
	client *DirectoryClient
}

// NewDirectoryService creates a new DirectoryService
func NewDirectoryService(client *DirectoryClient) *DirectoryService {
	return &DirectoryService{
		client: client,
	}
}

// SearchByMobileNumber searches for a person by mobile number
func (s *DirectoryService) SearchByMobileNumber(ctx context.Context, mobileNumber string) (*models.DirectorySearchResponse, error) {
	if mobileNumber == "" {
		return nil, fmt.Errorf("mobile number cannot be empty")
	}

	// Remove any non-digit characters from the mobile number
	cleanNumber := cleanPhoneNumber(mobileNumber)

	// Search for the person in the directory
	result, err := s.client.SearchPerson(ctx, cleanNumber)
	if err != nil {
		return nil, fmt.Errorf("error searching directory: %w", err)
	}

	return result, nil
}

// SearchByOrganizationNumber searches for a company by organization number
func (s *DirectoryService) SearchByOrganizationNumber(ctx context.Context, orgNo string) (*models.DirectorySearchResponse, error) {
	if orgNo == "" {
		return nil, fmt.Errorf("organization number cannot be empty")
	}

	// Clean the organization number (remove spaces and dots)
	cleanOrgNo := cleanOrganizationNumber(orgNo)

	// Search for the company in the directory
	result, err := s.client.SearchByOrganizationNumber(ctx, cleanOrgNo)
	if err != nil {
		return nil, fmt.Errorf("error searching directory: %w", err)
	}

	return result, nil
}

// cleanPhoneNumber removes all non-digit characters from a phone number
func cleanPhoneNumber(phone string) string {
	var result []rune
	for _, r := range phone {
		if r >= '0' && r <= '9' {
			result = append(result, r)
		}
	}
	return string(result)
}

// cleanOrganizationNumber removes spaces and dots from an organization number
func cleanOrganizationNumber(orgNo string) string {
	var result []rune
	for _, r := range orgNo {
		if (r >= '0' && r <= '9') || (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
			result = append(result, r)
		}
	}
	return string(result)
}
