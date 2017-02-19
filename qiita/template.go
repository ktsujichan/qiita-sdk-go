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

// Represents a template for generating an item boilerplate (only available on Qiita:Team).
type Template struct {
	*ExpandedTemplate
	Body  string    `json:"body"`
	Id    uint      `json:"id,omitempty"`
	Name  string    `json:"name"`
	Tags  *Taggings `json:"tags"`
	Title string    `json:"title"`
}

type Templates []Template

/*
	List templates in a team.

	GET /api/v2/templates
*/
func (c *Client) ListTemplates(ctx context.Context, page, perPage uint) (*Templates, error) {
	values := url.Values{}
	values.Add("page", fmt.Sprint(page))
	values.Add("per_page", fmt.Sprint(perPage))
	rawQuery := values.Encode()
	res, err := c.get(ctx, "/api/v2/templates", &rawQuery)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	var templates Templates
	if err := decodeBody(res, &templates); err != nil {
		return nil, err
	}
	return &templates, nil
}

/*
	Delete a template.

	DELETE /api/v2/templates/:template_id
*/
func (c *Client) DeleteTemplate(ctx context.Context, templateId uint) error {
	p := fmt.Sprintf("/api/v2/templates/%d", templateId)
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
	Get a template.

	GET /api/v2/templates/:template_id
*/
func (c *Client) GetTemplate(ctx context.Context, templateId uint) (*Template, error) {
	p := fmt.Sprintf("/api/v2/templates/%d", templateId)
	res, err := c.get(ctx, p, nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	var template Template
	if err := decodeBody(res, &template); err != nil {
		return nil, err
	}
	return &template, nil
}

/*
	Create a new template.

	POST /api/v2/templates
*/
func (c *Client) CreateTemplate(ctx context.Context, template Template) error {
	b, _ := json.Marshal(template)
	res, err := c.post(ctx, "/api/v2/templates", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusCreated {
		return errors.New(res.Status)
	}
	return nil
}

/*
	Update a template.

	PATCH /api/v2/templates/:template_id
*/
func (c *Client) UpdateTemplate(ctx context.Context, template Template) error {
	b, _ := json.Marshal(template)
	p := fmt.Sprintf("/api/v2/templates/%d", template.Id)
	res, err := c.patch(ctx, p, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}
	return nil
}
