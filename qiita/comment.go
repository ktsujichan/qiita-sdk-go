package qiita

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// A comment posted on an item
type Comment struct {
	Body         string `json:"body"`
	CreatedAt    string `json:"created_at,omitempty"`
	ID           string `json:"id,omitempty"`
	RenderedBody string `json:"rendered_body,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	User         User   `json:"user,omitempty"`
}

type Comments []Comment

/*
	Delete a comment.

	DELETE /api/v2/comments/:comment_id
*/
func (c *Client) DeleteComment(ctx context.Context, commentID string) error {
	p := fmt.Sprintf("comments/%s", commentID)
	res, err := c.delete(ctx, p)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusNoContent {
		return errors.New(res.Status)
	}
	return nil
}

/*
	Get a comment.

	GET /api/v2/comments/:comment_id
*/
func (c *Client) GetComment(ctx context.Context, commentID string) (*Comment, error) {
	p := fmt.Sprintf("comments/%s", commentID)
	res, err := c.get(ctx, p, nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	var comment Comment
	if err := decodeBody(res, &comment); err != nil {
		return nil, err
	}
	return &comment, nil
}

/*
	Update a comment.

	PATCH /api/v2/comments/:comment_id
*/
func (c *Client) UpdateComment(ctx context.Context, comment Comment) error {
	b, err := json.Marshal(comment)
	if err != nil {
		return err
	}
	p := fmt.Sprintf("comments/%s", comment.ID)
	res, err := c.patch(ctx, p, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}
	return nil
}

/*
	GET /api/v2/items/:item_id/comments

	List comments on an item in newest order.
*/
func (c *Client) ListComments(ctx context.Context, itemID string) (*Comments, error) {
	p := fmt.Sprintf("items/%s/comments", itemID)
	res, err := c.get(ctx, p, nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	var comments Comments
	if err := decodeBody(res, &comments); err != nil {
		return nil, err
	}
	return &comments, nil
}

/*
	Post a comment on an item.

	POST /api/v2/items/:item_id/comments
*/
func (c *Client) PostComment(ctx context.Context, itemID string, comment Comment) error {
	b, err := json.Marshal(comment)
	if err != nil {
		return err
	}
	p := fmt.Sprintf("items/%s/comments", itemID)
	res, err := c.post(ctx, p, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusCreated {
		return errors.New(res.Status)
	}
	return nil
}
