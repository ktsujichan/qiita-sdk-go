package qiita

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// A tag attached to an item
type Tag struct {
	FollowersCount uint   `json:"followers_count"`
	IconURL        string `json:"icon_url"`
	ID             string `json:"id"`
	ItemsCount     uint   `json:"items_count"`
}

type Tags []Tag

/*
	List tags in newest order.

	GET /api/v2/tags
*/
func (c *Client) ListTags(ctx context.Context, page, perPage uint, sort string) (*Tags, error) {
	res, err := c.get(ctx, "tags", map[string]interface{}{
		"page":     page,
		"per_page": perPage,
		"sort":     sort,
	})
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	var tags Tags
	if err := decodeBody(res, &tags); err != nil {
		return nil, err
	}
	return &tags, nil
}

/*
	Get a tag.

	GET /api/v2/tags/:tag_id
*/
func (c *Client) GetTag(ctx context.Context, tagID string) (*Tag, error) {
	p := fmt.Sprintf("tags/%s", tagID)
	res, err := c.get(ctx, p, nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	var tag Tag
	if err := decodeBody(res, &tag); err != nil {
		return nil, err
	}
	return &tag, nil
}

/*
	List tags a user is following to in recently-tagged order.

	GET /api/v2/users/:user_id/following_tags
*/
func (c *Client) ListFollowingTags(ctx context.Context, userID string, page, perPage uint) (*Tags, error) {
	p := fmt.Sprintf("users/%s/following_tags", userID)
	res, err := c.get(ctx, p, map[string]interface{}{
		"page":     page,
		"per_page": perPage,
	})
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	var tags Tags
	if err := decodeBody(res, &tags); err != nil {
		return nil, err
	}
	return &tags, nil
}

/*
 	Unfollow a tag.

	DELETE /api/v2/tags/:tag_id/following
*/
func (c *Client) UnfollowTag(ctx context.Context, tagID string) error {
	p := fmt.Sprintf("tags/%s/following", tagID)
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
	Check if you are following a tag or not.

	GET /api/v2/tags/:tag_id/following
*/
func (c *Client) EnsureFollowingTag(ctx context.Context, tagID string) error {
	p := fmt.Sprintf("tags/%s/following", tagID)
	res, err := c.get(ctx, p, nil)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusNoContent {
		return errors.New(res.Status)
	}
	return nil
}

/*
	Follow a tag.

	PUT /api/v2/tags/:tag_id/following
*/
func (c *Client) FollowTag(ctx context.Context, tagID string) error {
	p := fmt.Sprintf("tags/%s/following", tagID)
	res, err := c.put(ctx, p, nil)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusNoContent {
		return errors.New(res.Status)
	}
	return nil
}
