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
	Id           string `json:"id,omitempty"`
	RenderedBody string `json:"rendered_body,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	User         User   `json:"user,omitempty"`
}

type Comments []Comment

/*
	Delete a comment.

	DELETE /api/v2/comments/:comment_id
*/
func (c *Client) DeleteComment(ctx context.Context, commentId string) error {
	p := fmt.Sprintf("/api/v2/comments/%s", commentId)
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
func (c *Client) GetComment(ctx context.Context, commentId string) (*Comment, error) {
	p := fmt.Sprintf("/api/v2/comments/%s", commentId)
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
	b, _ := json.Marshal(comment)
	p := fmt.Sprintf("/api/v2/comments/%s", comment.Id)
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
func (c *Client) ListComments(ctx context.Context, itemId string) (*Comments, error) {
	p := fmt.Sprintf("/api/v2/items/%s/comments", itemId)
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
func (c *Client) PostComment(ctx context.Context, itemId string, comment Comment) error {
	b, _ := json.Marshal(comment)
	p := fmt.Sprintf("/api/v2/items/%s/comments", itemId)
	res, err := c.post(ctx, p, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusCreated {
		return errors.New(res.Status)
	}
	return nil
}
