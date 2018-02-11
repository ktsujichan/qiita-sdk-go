package qiita

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// You can preview the expanded result of a given template. This is available only on Qiita:Team.
type ExpandedTemplate struct {
	ExpandedBody  string   `json:"expanded_body"`
	ExpandedTags  Taggings `json:"expanded_tags"`
	ExpandedTitle string   `json:"expanded_title"`
}

/*
	Get a template where its variables are expanded.

	POST /api/v2/expanded_templates
*/
func (c *Client) CreateExpandedTemplate(ctx context.Context, template Template) (*ExpandedTemplate, error) {
	b, err := json.Marshal(template)
	if err != nil {
		return nil, err
	}
	res, err := c.post(ctx, "expanded_templates", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusCreated {
		return nil, errors.New(res.Status)
	}
	var expandedTemplate ExpandedTemplate
	if err := decodeBody(res, &expandedTemplate); err != nil {
		return nil, err
	}
	return &expandedTemplate, nil
}
