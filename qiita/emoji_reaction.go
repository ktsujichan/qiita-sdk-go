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
	CreatedAt string `json:"created_at,omitemoty"`
	ImageUrl  string `json:"image_url,omitemoty"`
	Name      string `json:"name"`
	User      User   `json:"user,omitemoty"`
}

type Reactions []Reaction

/*
	Add an emoji reaction to a comment.

	POST /api/v2/comments/:comment_id/reactions
*/
func (c *Client) AddCommentReaction(ctx context.Context, commentId string, reaction Reaction) error {
	b, _ := json.Marshal(reaction)
	p := fmt.Sprintf("/api/v2/comments/%s/reactions", commentId)
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
func (c *Client) AddItemReaction(ctx context.Context, itemId string, reaction Reaction) error {
	b, _ := json.Marshal(reaction)
	p := fmt.Sprintf("/api/v2/items/%s/reactions", itemId)
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
func (c *Client) AddProjectReaction(ctx context.Context, projectId uint, reaction Reaction) error {
	b, _ := json.Marshal(reaction)
	p := fmt.Sprintf("/api/v2/projects/%d/reactions", projectId)
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
func (c *Client) DeleteCommentReaction(ctx context.Context, commentId, reactionName string) error {
	p := fmt.Sprintf("/api/v2/comments/%s/reactions/%s", commentId, reactionName)
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
func (c *Client) DeleteItemReaction(ctx context.Context, itemId, reactionName string) error {
	p := fmt.Sprintf("/api/v2/items/%s/reactions/%s", itemId, reactionName)
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
func (c *Client) DeleteProjectReaction(ctx context.Context, projectId uint, reactionName string) error {
	p := fmt.Sprintf("/api/v2/projects/%d/reactions/%s", projectId, reactionName)
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
func (c *Client) ListCommentReactions(ctx context.Context, commentId string) (*Reactions, error) {
	p := fmt.Sprintf("/api/v2/comments/%s/reactions", commentId)
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
func (c *Client) ListItemReactions(ctx context.Context, itemId string) (*Reactions, error) {
	p := fmt.Sprintf("/api/v2/items/%s/reactions", itemId)
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
func (c *Client) ListProjectReactions(ctx context.Context, projectId string) (*Reactions, error) {
	p := fmt.Sprintf("/api/v2/projects/%s/reactions", projectId)
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
