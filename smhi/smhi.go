package smhi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	baseURL = "https://opendata-download-metobs.smhi.se/"
)

// Client is a client
type Client struct {
	client  *http.Client
	BaseURL *url.URL

	common service

	Temperatures *TemperatureService
}

type service struct {
	client *Client
}

// Service is the main service
type Service service

// NewClient creates a new client
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	parsedURL, _ := url.Parse(baseURL)
	c := &Client{client: httpClient, BaseURL: parsedURL}

	c.common.client = c

	c.Temperatures = (*TemperatureService)(&c.common)

	return c
}

// NewRequest creates a new request for a resource
func (c *Client) NewRequest(method, urlStr string) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// Do executes the request
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (resp *http.Response, err error) {
	req = req.WithContext(ctx)

	resp, err = c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}

	defer func() {
		err = resp.Body.Close()
	}()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, ioErr := io.Copy(w, resp.Body)
			if ioErr == io.EOF {
				err = ioErr
			}
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return resp, err
}

// CheckResponse checks the response for error codes in the response
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	return &ErrorResponse{Response: r}

}

// ErrorResponse wraps a response as an error
type ErrorResponse struct {
	Response *http.Response
}

// Error implements the errorer interface
func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("Error code %d", e.Response.StatusCode)
}
