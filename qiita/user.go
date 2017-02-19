package qiita

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// A Qiita user (a.k.a. account)
type User struct {
	Description       string `json:"description"`
	FacebookId        string `json:"facebook_id"`
	FolloweesCount    uint   `json:"followees_count"`
	FollowersCount    uint   `json:"followers_count"`
	GitHubLoginName   string `json:"github_login_name"`
	Id                string `json:"id"`
	ItemCount         uint   `json:"item_count"`
	LinkedinId        string `json:"linkedin_id"`
	Location          string `json:"location"`
	Name              string `json:"name"`
	Organization      string `json:"organization"`
	PermanentId       uint   `json:"permanent_id"`
	ProfileImageUrl   string `json:"profile_image_url"`
	TwitterScreenName string `json:"twitter_screen_name"`
	WebsiteUrl        string `json:"website_url"`
}

type Users []User

/*
	List users who stocked an item in recent-stocked order.

	GET /api/v2/items/:item_id/stockers
*/
func (c *Client) ListStockers(ctx context.Context, itemId string, page, perPage uint) (*Users, error) {
	p := fmt.Sprintf("/api/v2/items/%s/stockers", itemId)
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
	var users Users
	if err := decodeBody(res, &users); err != nil {
		return nil, err
	}
	return &users, nil
}

/*
	List users in newest order.

	GET /api/v2/users
*/
func (c *Client) ListUsers(ctx context.Context, page, perPage uint) (*Users, error) {
	values := url.Values{}
	values.Add("page", fmt.Sprint(page))
	values.Add("per_page", fmt.Sprint(perPage))
	rawQuery := values.Encode()
	res, err := c.get(ctx, "/api/v2/users", &rawQuery)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	var users Users
	if err := decodeBody(res, &users); err != nil {
		return nil, err
	}
	return &users, nil
}

/*
	Get a user.

	GET /api/v2/users/:user_id
*/
func (c *Client) GetUser(ctx context.Context, userId string) (*User, error) {
	p := fmt.Sprintf("/api/v2/users/%s", userId)
	res, err := c.get(ctx, p, nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	var user User
	if err := decodeBody(res, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

/*
	List users a user is following.

	GET /api/v2/users/:user_id/followees
*/
func (c *Client) ListFollowees(ctx context.Context, userId string, page, perPage uint) (*Users, error) {
	p := fmt.Sprintf("/api/v2/users/%s/followees", userId)
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
	var users Users
	if err := decodeBody(res, &users); err != nil {
		return nil, err
	}
	return &users, nil
}

/*
	List users who are following a user.

	GET /api/v2/users/:user_id/followers
*/
func (c *Client) ListFollowers(ctx context.Context, userId string, page, perPage uint) (*Users, error) {
	p := fmt.Sprintf("/api/v2/users/%s/followers", userId)
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
	var users Users
	if err := decodeBody(res, &users); err != nil {
		return nil, err
	}
	return &users, nil
}

/*
	Unfollow a user.

	DELETE /api/v2/users/:user_id/following
*/
func (c *Client) UnfollowUser(ctx context.Context, userId string) error {
	p := fmt.Sprintf("/api/v2/users/%s/following", userId)
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
	Check if the current user is following a user.

	GET /api/v2/users/:user_id/following
*/
func (c *Client) EnsureFollowingUser(ctx context.Context, userId string) error {
	p := fmt.Sprintf("/api/v2/users/%s/following", userId)
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
	Follow a user.

	PUT /api/v2/users/:user_id/following
*/
func (c *Client) FollowUser(ctx context.Context, userId string) error {
	p := fmt.Sprintf("/api/v2/users/%s/following", userId)
	res, err := c.put(ctx, p, nil)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusNoContent {
		return errors.New(res.Status)
	}
	return nil
}
