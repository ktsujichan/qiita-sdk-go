package qiita

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Access token for Qiita API v2
type AccessToken struct {
	ClientID string   `json:"client_id"`
	Scopes   []string `json:"scopes"`
	Token    string   `json:"token"`
}

type Auth struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
}

/*
	Create a new access token.

	POST /api/v2/access_tokens
*/
func (c *Client) CreateAccessToken(ctx context.Context, auth Auth) error {
	b, err := json.Marshal(auth)
	if err != nil {
		return err
	}
	res, err := c.post(ctx, "access_tokens", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusCreated {
		return errors.New(res.Status)
	}
	return nil
}

/*
	DELETE /api/v2/access_tokens/:access_token

	Deactivate an access token.
*/
func (c *Client) DeleteAccessToken(ctx context.Context, accessToken string) error {
	p := fmt.Sprintf("access_tokens/%s", accessToken)
	res, err := c.delete(ctx, p)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusNoContent {
		return errors.New(res.Status)
	}
	return nil
}
