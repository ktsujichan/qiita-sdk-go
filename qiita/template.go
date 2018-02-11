package qiita

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Represents a template for generating an item boilerplate (only available on Qiita:Team).
type Template struct {
	*ExpandedTemplate
	Body  string    `json:"body"`
	ID    uint      `json:"id,omitempty"`
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
	res, err := c.get(ctx, "templates", map[string]interface{}{
		"page":     page,
		"per_page": perPage,
	})
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
func (c *Client) DeleteTemplate(ctx context.Context, templateID uint) error {
	p := fmt.Sprintf("templates/%d", templateID)
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
func (c *Client) GetTemplate(ctx context.Context, templateID uint) (*Template, error) {
	p := fmt.Sprintf("templates/%d", templateID)
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
	b, err := json.Marshal(template)
	if err != nil {
		return err
	}
	res, err := c.post(ctx, "templates", bytes.NewBuffer(b))
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
	b, err := json.Marshal(template)
	if err != nil {
		return err
	}
	p := fmt.Sprintf("templates/%d", template.ID)
	res, err := c.patch(ctx, p, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}
	return nil
}
