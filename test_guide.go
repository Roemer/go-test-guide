package gotestguide

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Error for 404 not found responses.
var ErrNotFound = errors.New("404 Not Found")

var TruePtr = Ptr(true)
var FalsePtr = Ptr(false)

// Ptr is a helper that returns a pointer to v.
func Ptr[T any](v T) *T {
	return &v
}

// A client to interact with the test.guide API.
type Client struct {
	baseUrl *url.URL
	authKey string
	debug   bool

	// API for up- and download of artifacts to/from test.guide.
	Artifacts ArtifactsServiceInterface
	// API for handling project and system settings in test.guide.
	Platform PlatformServiceInterface
	// API for handling test reports (test case executions) in test.guide.
	ReportManagement ReportManagementServiceInterface
	// API for accessing user management functionality of test.guide.
	UserManagement UserManagementServiceInterface
}

// Create a new client for the test.guide API.
func NewClient(baseUrl, authKey string) (*Client, error) {
	parsedUrl, err := url.Parse(baseUrl)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}
	client := &Client{
		baseUrl: parsedUrl,
		authKey: authKey,
	}
	client.Artifacts = &ArtifactsService{client: client}
	client.Platform = &PlatformService{client: client}
	client.ReportManagement = &ReportManagementService{client: client}
	client.UserManagement = &UserManagementService{client: client}
	return client, nil
}

// Create a new HTTP request with authentication.
func (c *Client) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", c.baseUrl, path), body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("TestGuide-AuthKey", c.authKey)
	return req, nil
}

// Execute an HTTP request and decode the response into the provided variable.
func (c *Client) Do(req *http.Request, v any) (*http.Response, error) {
	if c.debug {
		fmt.Printf("Sending request: %s %s\n", req.Method, req.URL)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer func() {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}()

	// For debugging
	if c.debug {
		bodyBytes, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Println(string(bodyBytes))
		resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	// Verify the response
	err = c.checkResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	// Decode the response body if a variable is provided
	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return nil, err
		}
	}
	return resp, nil
}

// Enable debugging (printing) of all received values.
func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c *Client) checkResponse(resp *http.Response) error {
	if resp.StatusCode == 404 {
		return ErrNotFound
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, body)
	}
	return nil
}
