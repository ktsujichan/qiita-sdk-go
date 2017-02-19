package qiita

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// A tag attached to an item
type Tag struct {
	FollowersCount uint   `json:"followers_count"`
	IconUrl        string `json:"icon_url"`
	Id             string `json:"id"`
	ItemsCount     uint   `json:"items_count"`
}

type Tags []Tag

/*
	List tags in newest order.

	GET /api/v2/tags
*/
func (c *Client) ListTags(ctx context.Context, page, perPage uint, sort string) (*Tags, error) {
	values := url.Values{}
	values.Add("page", fmt.Sprint(page))
	values.Add("per_page", fmt.Sprint(perPage))
	values.Add("sort", sort)
	rawQuery := values.Encode()
	res, err := c.get(ctx, "/api/v2/tags", &rawQuery)
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
func (c *Client) GetTag(ctx context.Context, tagId string) (*Tag, error) {
	p := fmt.Sprintf("/api/v2/tags/%s", tagId)
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
func (c *Client) ListFollowingTags(ctx context.Context, userId string, page, perPage uint) (*Tags, error) {
	p := fmt.Sprintf("/api/v2/users/%s/following_tags", userId)
	values := url.Values{}
	values.Add("page", fmt.Sprint(page))
	values.Add("per_page", fmt.Sprint(perPage))
	rawQuery := values.Encode()
	res, err := c.get(ctx, p, &rawQuery)
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
func (c *Client) UnfollowTag(ctx context.Context, tagId string) error {
	p := fmt.Sprintf("/api/v2/tags/%s/following", tagId)
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
func (c *Client) EnsureFollowingTag(ctx context.Context, tagId string) error {
	p := fmt.Sprintf("/api/v2/tags/%s/following", tagId)
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
func (c *Client) FollowTag(ctx context.Context, tagId string) error {
	p := fmt.Sprintf("/api/v2/tags/%s/following", tagId)
	res, err := c.put(ctx, p, nil)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusNoContent {
		return errors.New(res.Status)
	}
	return nil
}
