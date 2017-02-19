package qiita

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// Represents an item posted from a user
type Item struct {
	Body         string    `json:"body"`
	Coediting    bool      `json:"coediting"`
	CreatedAt    string    `json:"created_at,omitempty"`
	Gist         bool      `json:"gist,omitempty"`
	Group        *Group    `json:"group,omitempty"`
	Id           string    `json:"id,omitempty"`
	Private      bool      `json:"private"`
	Tags         *Taggings `json:"tags"`
	Title        string    `json:"title"`
	Tweet        bool      `json:"tweet,omitempty"`
	RenderedBody string    `json:"rendered_body,omitempty"`
	UpdatedAt    string    `json:"updated_at,omitempty"`
	Url          string    `json:"url,omitempty"`
	User         *User     `json:"user,omitempty"`
}

type Items []Item

/*
	List the authenticated user's items in newest order

	GET /api/v2/authenticated_user/items
*/
func (c *Client) ListAuthenticatedUserItems(ctx context.Context, page, perPage uint) (*Items, error) {
	values := url.Values{}
	values.Add("page", fmt.Sprint(page))
	values.Add("per_page", fmt.Sprint(perPage))
	rawQuery := values.Encode()
	res, err := c.get(ctx, "/api/v2/authenticated_user/items", &rawQuery)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	var items Items
	if err := decodeBody(res, &items); err != nil {
		return nil, err
	}
	return &items, nil
}

/*
	List items.

	GET /api/v2/items
*/
func (c *Client) ListItems(ctx context.Context, page, perPage uint, query string) (*Items, error) {
	values := url.Values{}
	values.Add("page", fmt.Sprint(page))
	values.Add("per_page", fmt.Sprint(perPage))
	values.Add("query", query)
	rawQuery := values.Encode()
	res, err := c.get(ctx, "/api/v2/items", &rawQuery)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	var items Items
	if err := decodeBody(res, &items); err != nil {
		return nil, err
	}
	return &items, nil
}

/*
	Create an item.

	POST /api/v2/items
*/
func (c *Client) CreateItem(ctx context.Context, item Item) error {
	b, _ := json.Marshal(item)
	res, err := c.post(ctx, "/api/v2/items", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusCreated {
		return errors.New(res.Status)
	}
	return nil
}

/*
	Delete an item.

	DELETE /api/v2/items/:item_id
*/
func (c *Client) DeleteItem(ctx context.Context, itemId string) error {
	p := fmt.Sprintf("/api/v2/items/%s", itemId)
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
	Get an item.

	GET /api/v2/items/:item_id
*/
func (c *Client) GetItem(ctx context.Context, itemId string) (*Item, error) {
	p := fmt.Sprintf("/api/v2/items/%s", itemId)
	res, err := c.get(ctx, p, nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	var item Item
	if err := decodeBody(res, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

/*
	Update an item.

	PATCH /api/v2/items/:item_id
*/
func (c *Client) UpdateItem(ctx context.Context, item Item) error {
	b, _ := json.Marshal(item)
	p := fmt.Sprintf("/api/v2/items/%s", item.Id)
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
	Unlike an item (only available on Qiita:Team).

	DELETE /api/v2/items/:item_id/like
*/
func (c *Client) UnlikeItem(ctx context.Context, itemId string) error {
	p := fmt.Sprintf("/api/v2/items/%s/like", itemId)
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
	Like an item (only available on Qiita:Team).

	PUT /api/v2/items/:item_id/like
*/
func (c *Client) LikeItem(ctx context.Context, itemId string) error {
	p := fmt.Sprintf("/api/v2/items/%s/like", itemId)
	res, err := c.put(ctx, p, nil)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusNoContent {
		return errors.New(res.Status)
	}
	return nil
}

/*
	Stock an item.

	PUT /api/v2/items/:item_id/stock
*/
func (c *Client) StockItem(ctx context.Context, itemId string) error {
	p := fmt.Sprintf("/api/v2/items/%s/stock", itemId)
	res, err := c.put(ctx, p, nil)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusNoContent {
		return errors.New(res.Status)
	}
	return nil
}

/*
	Unstock an item.

	DELETE /api/v2/items/:item_id/stock
*/
func (c *Client) UnstockItem(ctx context.Context, itemId string) error {
	p := fmt.Sprintf("/api/v2/items/%s/stock", itemId)
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
	Check if you stocked an item.

	GET /api/v2/items/:item_id/stock
*/
func (c *Client) EnsureItemStock(ctx context.Context, itemId string) error {
	p := fmt.Sprintf("/api/v2/items/%s/stock", itemId)
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
	Check if you liked an item (only available on Qiita:Team).

	GET /api/v2/items/:item_id/like
*/
func (c *Client) EnsureItemLike(ctx context.Context, itemId string) error {
	p := fmt.Sprintf("/api/v2/items/%s/like", itemId)
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
	List tagged items in recently-tagged order.

	GET /api/v2/tags/:tag_id/items
*/
func (c *Client) ListTaggedItems(ctx context.Context, tagId string, page, perPage uint) (*Items, error) {
	p := fmt.Sprintf("/api/v2/tags/%s/items", tagId)
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
	var items Items
	if err := decodeBody(res, &items); err != nil {
		return nil, err
	}
	return &items, nil
}

/*
	List a user's items in newest order.

	GET /api/v2/users/:user_id/items
*/
func (c *Client) ListUserItems(ctx context.Context, userId string, page, perPage uint) (*Items, error) {
	p := fmt.Sprintf("/api/v2/users/%s/items", userId)
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
	var items Items
	if err := decodeBody(res, &items); err != nil {
		return nil, err
	}
	return &items, nil
}

/*
	List a user's stocked items in recently-stocked order.

	GET /api/v2/users/:user_id/stocks
*/
func (c *Client) ListUserStocks(ctx context.Context, userId string, page, perPage uint) (*Items, error) {
	p := fmt.Sprintf("/api/v2/users/%s/stocks", userId)
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
	var items Items
	if err := decodeBody(res, &items); err != nil {
		return nil, err
	}
	return &items, nil
}
