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
	ClientId string   `json:"client_id"`
	Scopes   []string `json:"scopes"`
	Token    string   `json:"token"`
}

type Auth struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
}

/*
	Create a new access token.

	POST /api/v2/access_tokens
*/
func (c *Client) CreateAccessToken(ctx context.Context, auth Auth) error {
	b, _ := json.Marshal(auth)
	res, err := c.post(ctx, "/api/v2/access_tokens", bytes.NewBuffer(b))
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
	p := fmt.Sprintf("/api/v2/access_tokens/%s", accessToken)
	res, err := c.delete(ctx, p)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusNoContent {
		return errors.New(res.Status)
	}
	return nil
}
