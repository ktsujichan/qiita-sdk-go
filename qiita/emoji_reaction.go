package qiita

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// An emoji reaction.
type Reaction struct {
	CreatedAt string `json:"created_at,omitempty"`
	ImageURL  string `json:"image_url,omitempty"`
	Name      string `json:"name"`
	User      User   `json:"user,omitempty"`
}

type Reactions []Reaction

/*
	Add an emoji reaction to a comment.

	POST /api/v2/comments/:comment_id/reactions
*/
func (c *Client) AddCommentReaction(ctx context.Context, commentID string, reaction Reaction) error {
	b, err := json.Marshal(reaction)
	if err != nil {
		return err
	}
	p := fmt.Sprintf("/api/v2/comments/%s/reactions", commentID)
	res, err := c.post(ctx, p, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusCreated {
		return errors.New(res.Status)
	}
	return nil
}

/*
	Add an emoji reaction to an item.

	POST /api/v2/items/:item_id/reactions
*/
func (c *Client) AddItemReaction(ctx context.Context, itemID string, reaction Reaction) error {
	b, err := json.Marshal(reaction)
	if err != nil {
		return err
	}
	p := fmt.Sprintf("/api/v2/items/%s/reactions", itemID)
	res, err := c.post(ctx, p, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusCreated {
		return errors.New(res.Status)
	}
	return nil
}

/*
	Add an emoji reaction to an project.

	POST /api/v2/projects/:project_id/reactions
*/
func (c *Client) AddProjectReaction(ctx context.Context, projectID uint, reaction Reaction) error {
	b, err := json.Marshal(reaction)
	if err != nil {
		return err
	}
	p := fmt.Sprintf("/api/v2/projects/%d/reactions", projectID)
	res, err := c.post(ctx, p, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusCreated {
		return errors.New(res.Status)
	}
	return nil
}

/*
	Delete an emoji reaction from a comment.

	DELETE /api/v2/comments/:comment_id/reactions/:reaction_name
*/
func (c *Client) DeleteCommentReaction(ctx context.Context, commentID, reactionName string) error {
	p := fmt.Sprintf("/api/v2/comments/%s/reactions/%s", commentID, reactionName)
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
	Delete an emoji reaction from an item.

	DELETE /api/v2/items/:item_id/reactions/:reaction_name
*/
func (c *Client) DeleteItemReaction(ctx context.Context, itemID, reactionName string) error {
	p := fmt.Sprintf("/api/v2/items/%s/reactions/%s", itemID, reactionName)
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
	Delete an emoji reaction from an project.

	DELETE /api/v2/projects/:project_id/reactions/:reaction_name
*/
func (c *Client) DeleteProjectReaction(ctx context.Context, projectID uint, reactionName string) error {
	p := fmt.Sprintf("/api/v2/projects/%d/reactions/%s", projectID, reactionName)
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
	List emoji reactions of an comment in recently-created order.

	GET /api/v2/comments/:comment_id/reactions
*/
func (c *Client) ListCommentReactions(ctx context.Context, commentID string) (*Reactions, error) {
	p := fmt.Sprintf("/api/v2/comments/%s/reactions", commentID)
	res, err := c.get(ctx, p, nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	var reactions Reactions
	if err := decodeBody(res, &reactions); err != nil {
		return nil, err
	}
	return &reactions, nil
}

/*
	List emoji reactions of an item in recently-created order.

	GET /api/v2/items/:item_id/reactions
*/
func (c *Client) ListItemReactions(ctx context.Context, itemID string) (*Reactions, error) {
	p := fmt.Sprintf("/api/v2/items/%s/reactions", itemID)
	res, err := c.get(ctx, p, nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	var reactions Reactions
	if err := decodeBody(res, &reactions); err != nil {
		return nil, err
	}
	return &reactions, nil
}

/*
	List emoji reactions of an project in recently-created order.

	GET /api/v2/projects/:project_id/reactions
*/
func (c *Client) ListProjectReactions(ctx context.Context, projectID string) (*Reactions, error) {
	p := fmt.Sprintf("/api/v2/projects/%s/reactions", projectID)
	res, err := c.get(ctx, p, nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	var reactions Reactions
	if err := decodeBody(res, &reactions); err != nil {
		return nil, err
	}
	return &reactions, nil
}
