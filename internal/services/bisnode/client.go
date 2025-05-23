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
	"strings"
	"time"
)

// Client represents a Bisnode API client
type Client struct {
	httpClient *http.Client
	baseURL    string
	authHeader string
}

// NewClient creates a new Bisnode API client with Basic Authentication
func NewClient(cfg *config.BisnodeConfig) *Client {
	auth := base64.StdEncoding.EncodeToString(
		[]byte(fmt.Sprintf("%s:%s", cfg.ClientID, cfg.ClientSecret)),
	)

	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL:    strings.TrimSuffix(cfg.BaseURL, "/"),
		authHeader: "Basic " + auth,
	}
}

// doRequest performs an HTTP request with the given context and returns the response body
func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", c.authHeader)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API request failed with status %d: %s",
			resp.StatusCode, string(body))
	}

	return body, nil
}

// SearchPerson searches for a person by mobile number
func (c *Client) SearchPerson(ctx context.Context, mobileNumber string) (*models.Person, error) {
	url := fmt.Sprintf("%s/persons/search", c.baseURL)

	reqBody := struct {
		MobileNumber string `json:"mobileNumber"`
	}{
		MobileNumber: mobileNumber,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		strings.NewReader(string(jsonBody)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var person models.Person
	if err := json.Unmarshal(body, &person); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &person, nil
}
