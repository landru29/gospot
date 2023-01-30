package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/landru29/gospot/internal/music"
)

func (c *Client) ensureUser(ctx context.Context, token string) error {
	if c.user != nil {
		return nil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.urlWith("me").String(), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	var user music.User

	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return err
	}

	c.user = &user

	return nil
}
