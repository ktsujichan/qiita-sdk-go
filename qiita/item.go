package qiita

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Represents an item posted from a user
type Item struct {
	Body         string    `json:"body"`
	Coediting    bool      `json:"coediting"`
	CreatedAt    string    `json:"created_at,omitempty"`
	Gist         bool      `json:"gist,omitempty"`
	Group        *Group    `json:"group,omitempty"`
	ID           string    `json:"id,omitempty"`
	Private      bool      `json:"private"`
	Tags         *Taggings `json:"tags"`
	Title        string    `json:"title"`
	Tweet        bool      `json:"tweet,omitempty"`
	RenderedBody string    `json:"rendered_body,omitempty"`
	UpdatedAt    string    `json:"updated_at,omitempty"`
	URL          string    `json:"url,omitempty"`
	User         *User     `json:"user,omitempty"`
}

type Items []Item

/*
	List the authenticated user's items in newest order

	GET /api/v2/authenticated_user/items
*/
func (c *Client) ListAuthenticatedUserItems(ctx context.Context, page, perPage uint) (*Items, error) {
	res, err := c.get(ctx, "authenticated_user/items", map[string]interface{}{
		"page":     page,
		"per_page": perPage,
	})
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
	res, err := c.get(ctx, "items", map[string]interface{}{
		"page":     page,
		"per_page": perPage,
		"query":    query,
	})
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
	b, err := json.Marshal(item)
	if err != nil {
		return err
	}
	res, err := c.post(ctx, "items", bytes.NewBuffer(b))
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
func (c *Client) DeleteItem(ctx context.Context, itemID string) error {
	p := fmt.Sprintf("items/%s", itemID)
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
func (c *Client) GetItem(ctx context.Context, itemID string) (*Item, error) {
	p := fmt.Sprintf("items/%s", itemID)
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
	b, err := json.Marshal(item)
	if err != nil {
		return err
	}
	p := fmt.Sprintf("items/%s", item.ID)
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
func (c *Client) UnlikeItem(ctx context.Context, itemID string) error {
	p := fmt.Sprintf("items/%s/like", itemID)
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
func (c *Client) LikeItem(ctx context.Context, itemID string) error {
	p := fmt.Sprintf("items/%s/like", itemID)
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
func (c *Client) StockItem(ctx context.Context, itemID string) error {
	p := fmt.Sprintf("items/%s/stock", itemID)
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
func (c *Client) UnstockItem(ctx context.Context, itemID string) error {
	p := fmt.Sprintf("items/%s/stock", itemID)
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
func (c *Client) EnsureItemStock(ctx context.Context, itemID string) error {
	p := fmt.Sprintf("items/%s/stock", itemID)
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
func (c *Client) EnsureItemLike(ctx context.Context, itemID string) error {
	p := fmt.Sprintf("items/%s/like", itemID)
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
func (c *Client) ListTaggedItems(ctx context.Context, tagID string, page, perPage uint) (*Items, error) {
	p := fmt.Sprintf("tags/%s/items", tagID)
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
func (c *Client) ListUserItems(ctx context.Context, userID string, page, perPage uint) (*Items, error) {
	p := fmt.Sprintf("users/%s/items", userID)
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
func (c *Client) ListUserStocks(ctx context.Context, userID string, page, perPage uint) (*Items, error) {
	p := fmt.Sprintf("users/%s/stocks", userID)
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
	var items Items
	if err := decodeBody(res, &items); err != nil {
		return nil, err
	}
	return &items, nil
}
