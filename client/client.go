package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gitub.com/wellingtonchida/micro/types"
)

type Client struct {
	endpoint string
}

func New(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
	}
}

func (c *Client) FetchPrice(ctx context.Context, ticker string) (*types.PriceResponse, error) {
	endpoint := fmt.Sprintf("%s?ticker=%s", c.endpoint, ticker)
	req, err := http.NewRequest("get", endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		httpErr := map[string]any{}
		if err := json.NewDecoder(resp.Body).Decode(&httpErr); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("unexpected status code: %s", httpErr["error"])
	}

	priceresp := new(types.PriceResponse)
	if err := json.NewDecoder(resp.Body).Decode(priceresp); err != nil {
		return nil, err
	}

	return priceresp, nil
}
