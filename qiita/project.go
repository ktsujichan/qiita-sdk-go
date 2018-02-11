package qiita

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Represents a project on Qiita:Team (only available on Qiita:Team).
type Project struct {
	Archived     bool      `json:"archived"`
	Body         string    `json:"body"`
	CreatedAt    string    `json:"created_at,omitempty"`
	ID           uint      `json:"id,omitempty"`
	Name         string    `json:"name"`
	RenderedBody string    `json:"rendered_body,omitempty"`
	Tags         *Taggings `json:"tags,omitempty"`
	UpdatedAt    string    `json:"updated_at,omitempty"`
}

type Projects []Project

/*
	List projects in newest order.

	GET /api/v2/projects
*/
func (c *Client) ListProjects(ctx context.Context, page, perPage uint) (*Projects, error) {
	res, err := c.get(ctx, "projects", map[string]interface{}{
		"page":     page,
		"per_page": perPage,
	})
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	var projects Projects
	if err := decodeBody(res, &projects); err != nil {
		return nil, err
	}
	return &projects, nil
}

/*
	Create a new project.

	POST /api/v2/projects
*/
func (c *Client) CreateProject(ctx context.Context, project Project) error {
	b, err := json.Marshal(project)
	if err != nil {
		return err
	}
	res, err := c.post(ctx, "projects", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusCreated {
		return errors.New(res.Status)
	}
	return nil
}

/*
	Delete a project.

	DELETE /api/v2/projects/:project_id
*/
func (c *Client) DeleteProject(ctx context.Context, projectID uint) error {
	p := fmt.Sprintf("projects/%d", projectID)
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
	Get a project.

	GET /api/v2/projects/:project_id
*/
func (c *Client) GetProject(ctx context.Context, projectID string, page, perPage uint) (*Project, error) {
	p := fmt.Sprintf("projects/%s", projectID)
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
	var project Project
	if err := decodeBody(res, &project); err != nil {
		return nil, err
	}
	return &project, nil
}

/*
	Update a project

	PATCH /api/v2/projects/:project_id
*/
func (c *Client) UpdateProject(ctx context.Context, project Project) error {
	b, err := json.Marshal(project)
	if err != nil {
		return err
	}
	p := fmt.Sprintf("projects/%d", project.ID)
	res, err := c.patch(ctx, p, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}
	return nil
}
