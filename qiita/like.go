package qiita

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// Represents a like to an item (only available on Qiita:Team).
type Like struct {
	CreatedAt string `json:"created_at"`
	User      User   `json:"user"`
}

type Likes []Like

/*
	List likes in newest order (only available on Qiita:Team).

	GET /api/v2/items/:item_id/likes
*/
func (c *Client) ListItemLikes(ctx context.Context, itemID string) (*Likes, error) {
	p := fmt.Sprintf("items/%s/likes", itemID)
	res, err := c.get(ctx, p, nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	var likes Likes
	if err := decodeBody(res, &likes); err != nil {
		return nil, err
	}
	return &likes, nil
}
