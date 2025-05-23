package bisnode

import (
	"bisnode/internal/config"
	"bisnode/internal/models"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

// MotorVehicleClient handles communication with the Bisnode Motor Vehicle API
type MotorVehicleClient struct {
	baseURL    string
	authHeader string
	httpClient *http.Client
}

// NewMotorVehicleClient creates a new MotorVehicleClient
func NewMotorVehicleClient(cfg *config.BisnodeConfig) *MotorVehicleClient {
	auth := fmt.Sprintf("%s:%s", cfg.ClientID, cfg.ClientSecret)
	auth = "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))

	return &MotorVehicleClient{
		baseURL:    "https://api.bisnode.no",
		authHeader: auth,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// SearchByLicenseNumber searches for a vehicle by license number or VIN
func (c *MotorVehicleClient) SearchByLicenseNumber(ctx context.Context, searchTerm string) (*models.MotorVehicleSearchResponse, error) {
	log.Printf("Searching for motor vehicle with search term: %s", searchTerm)

	if searchTerm == "" {
		return nil, fmt.Errorf("search term cannot be empty")
	}

	// URL encode the search term to handle special characters
	encodedTerm := url.QueryEscape(searchTerm)
	url := fmt.Sprintf("%s/search/norway/motorvehicle/v2/%s", c.baseURL, encodedTerm)
	log.Printf("Sending request to: %s", url)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		err = fmt.Errorf("failed to create request: %w", err)
		log.Printf("Error creating request: %v", err)
		return nil, err
	}

	// Log the first 10 characters of the auth header for debugging
	if len(c.authHeader) > 10 {
		log.Printf("Using auth header: %s...", c.authHeader[:10])
	} else {
		log.Printf("Using auth header: %s", c.authHeader)
	}
	req.Header.Set("Authorization", c.authHeader)
	req.Header.Set("Accept", "application/json")

	log.Printf("Sending request with headers: %+v", req.Header)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		err = fmt.Errorf("request failed: %w", err)
		log.Printf("Request failed: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	log.Printf("Received response with status: %d %s", resp.StatusCode, resp.Status)

	// Read the response body
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		log.Printf("Error reading response body: %v", readErr)
		return nil, fmt.Errorf("failed to read response body: %w", readErr)
	}

	// Log the response body for debugging
	log.Printf("Response body: %s", string(body))

	// Check for error status codes
	if resp.StatusCode >= 400 {
		err = fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
		log.Printf("API error: %v", err)
		return nil, err
	}

	// Try to unmarshal the response
	var result models.MotorVehicleSearchResponse
	if err := json.Unmarshal(body, &result); err != nil {
		err = fmt.Errorf("failed to decode response: %w. Response body: %s", err, string(body))
		log.Printf("Error decoding response: %v", err)
		return nil, err
	}

	log.Printf("Successfully retrieved motor vehicle data")
	return &result, nil
}

// SearchByVIN searches for a vehicle by VIN (Vehicle Identification Number)
func (c *MotorVehicleClient) SearchByVIN(ctx context.Context, vin string) (*models.MotorVehicleSearchResponse, error) {
	// VIN search uses the same endpoint as license number search
	return c.SearchByLicenseNumber(ctx, vin)
}
