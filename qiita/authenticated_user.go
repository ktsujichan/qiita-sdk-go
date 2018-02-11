package qiita

import (
	"context"
	"errors"
	"net/http"
)

// An user currently authenticated by a given access token. This resources has more detailed information than normal User resource.
type AuthenticatedUser struct {
	*User
	ImageMonthlyUploadLimit     uint `json:"image_monthly_upload_limit"`
	ImageMonthlyUploadRemaining uint `json:"image_monthly_upload_remaining"`
	TeamOnly                    bool `json:"team_only"`
}

/*
	Get a user associated to the current access token.

	GET /api/v2/authenticated_user
*/
func (c *Client) GetAuthenticatedUser(ctx context.Context) (*AuthenticatedUser, error) {
	res, err := c.get(ctx, "authenticated_user", nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	var authenticatedUser AuthenticatedUser
	if err := decodeBody(res, &authenticatedUser); err != nil {
		return nil, err
	}
	return &authenticatedUser, nil
}
