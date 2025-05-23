package bisnode

import (
	"bisnode/internal/config"
	"bisnode/internal/models"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
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

// SearchByLicenseNumber searches for a vehicle by license number
func (c *MotorVehicleClient) SearchByLicenseNumber(ctx context.Context, licenseNumber string) (*models.MotorVehicleSearchResponse, error) {
	// URL encode the license number to handle special characters
	encodedLicense := url.QueryEscape(licenseNumber)
	url := fmt.Sprintf("%s/search/norway/motorvehicle/v2/%s", c.baseURL, encodedLicense)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", c.authHeader)
	req.Header.Set("Accept", "application/json")

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

	var result models.MotorVehicleSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// SearchByVIN searches for a vehicle by VIN (Vehicle Identification Number)
func (c *MotorVehicleClient) SearchByVIN(ctx context.Context, vin string) (*models.MotorVehicleSearchResponse, error) {
	// VIN search uses the same endpoint as license number search
	return c.SearchByLicenseNumber(ctx, vin)
}
