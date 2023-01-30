package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/landru29/gospot/internal/music"
)

type item struct {
	Album *music.Album `json:"album"`
}

type albumListResult struct {
	Items []item `json:"items"`
}

// Albums implements the Cataloger interface.
func (c *Client) Albums(ctx context.Context, token string) ([]music.Album, error) {
	urlRequest := c.urlWith("me/albums")

	values := url.Values{}
	values.Add("market", "ES")

	urlRequest.RawQuery = values.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlRequest.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	var albumResp albumListResult

	err = json.NewDecoder(resp.Body).Decode(&albumResp)
	if err != nil {
		return nil, err
	}

	albums := []music.Album{}

	for _, item := range albumResp.Items {
		if item.Album != nil {
			albums = append(albums, *item.Album)
		}
	}

	return albums, nil
}
