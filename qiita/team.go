package qiita

import (
	"context"
	"errors"
	"net/http"
)

// Represents a team on Qiita:Team (only available on Qiita:Team).
type Team struct {
	Archive bool   `json:"archive"`
	ID      string `json:"id"`
	Name    string `json:"name"`
}

type Teams []Team

/*
	List teams the user belongs to in newest order.

	GET /api/v2/teams
*/
func (c *Client) ListTeams(ctx context.Context) (*Teams, error) {
	res, err := c.get(ctx, "/api/v2/teams", nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	var teams Teams
	if err := decodeBody(res, &teams); err != nil {
		return nil, err
	}
	return &teams, nil
}
