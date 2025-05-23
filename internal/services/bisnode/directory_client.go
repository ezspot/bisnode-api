package bisnode

import (
	"bisnode/internal/config"
	"bisnode/internal/models"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// DirectoryClient handles communication with the Bisnode Directory Search API
type DirectoryClient struct {
	baseURL    string
	authHeader string
	httpClient *http.Client
}

// NewDirectoryClient creates a new DirectoryClient
func NewDirectoryClient(cfg *config.BisnodeConfig) *DirectoryClient {
	auth := base64.StdEncoding.EncodeToString(
		[]byte(fmt.Sprintf("%s:%s", cfg.ClientID, cfg.ClientSecret)),
	)

	return &DirectoryClient{
		baseURL:    "https://api.bisnode.no",
		authHeader: "Basic " + auth,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// SearchPerson searches for a person by mobile number
func (c *DirectoryClient) SearchPerson(ctx context.Context, mobileNumber string) (*models.DirectorySearchResponse, error) {
	url := fmt.Sprintf("%s/search/norway/directory", c.baseURL)
	
	// Prepare the request body
	reqBody := models.DirectorySearchRequest{}
	reqBody.Form.Type = "Freetext"
	reqBody.Form.SearchString = mobileNumber
	
	// Set search options
	reqBody.Options.SearchMode = 3      // Smart Exact + Phonetic
	reqBody.Options.OnlyFoundWords = true
	reqBody.Options.ListingType = 2     // Person only
	reqBody.Options.ResultLimit = 10    // Limit to 10 results

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", c.authHeader)
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status %d: %s", 
			resp.StatusCode, string(body))
	}

	var result models.DirectorySearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// SearchByOrganizationNumber searches for a company by organization number
func (c *DirectoryClient) SearchByOrganizationNumber(ctx context.Context, orgNo string) (*models.DirectorySearchResponse, error) {
	url := fmt.Sprintf("%s/search/norway/directory/%s", c.baseURL, orgNo)
	
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", c.authHeader)
	
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status %d: %s", 
			resp.StatusCode, string(body))
	}

	var result models.DirectorySearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}
