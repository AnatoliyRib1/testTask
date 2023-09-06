package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"golang.org/x/time/rate"

	"github.com/go-resty/resty/v2"

	"testTask/contracts"
)

type Client struct {
	client  *resty.Client
	baseURL string
	limiter *rate.Limiter
	mutex   sync.Mutex
}

func NewClient(url string) *Client {
	hc := &http.Client{}
	rc := resty.NewWithClient(hc)
	rc.OnAfterResponse(func(client *resty.Client, response *resty.Response) error {
		if response.IsError() {
			herr := contracts.HTTPError{}
			_ = json.Unmarshal(response.Body(), &herr)

			return &Error{Code: response.StatusCode(), Message: herr.Message}
		}
		return nil
	})

	return &Client{
		client:  rc,
		baseURL: url,
	}
}

func (c *Client) path(f string, args ...any) string {
	return fmt.Sprintf(c.baseURL+f, args...)
}
