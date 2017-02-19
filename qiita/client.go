package qiita

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"runtime"
)

type Client struct {
	URL        *url.URL
	HTTPClient *http.Client
	Token      string
}

var userAgent = fmt.Sprintf("QiitaGoClient/%s (%s)", version, runtime.Version())

func NewClient(token string, config Config) (*Client, error) {
	parsedURL, err := url.ParseRequestURI(config.Endpoint)
	if err != nil {
		return nil, err
	}
	return &Client{
		URL:        parsedURL,
		HTTPClient: &http.Client{},
		Token:      token,
	}, nil
}

func (c *Client) newRequest(ctx context.Context, method, p string, body io.Reader) (*http.Request, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, p)
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", userAgent)
	return req, nil
}

func decodeBody(res *http.Response, out interface{}) error {
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	return decoder.Decode(out)
}

func (c *Client) url(endpoint string) string {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, endpoint)
	return u.String()
}

func (c *Client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("User-Agent", userAgent)
	return c.HTTPClient.Do(req)
}

func (c *Client) get(ctx context.Context, endpoint string, rawQuery *string) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.url(endpoint), nil)
	if rawQuery != nil {
		req.URL.RawQuery = *rawQuery
	}
	if err != nil {
		return nil, err
	}
	return c.do(ctx, req)
}

func (client *Client) post(ctx context.Context, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", client.url(endpoint), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return client.do(ctx, req)
}

func (client *Client) patch(ctx context.Context, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("PATCH", client.url(endpoint), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return client.do(ctx, req)
}

func (client *Client) put(ctx context.Context, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("PUT", client.url(endpoint), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return client.do(ctx, req)
}

func (c *Client) delete(ctx context.Context, endpoint string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", c.url(endpoint), nil)
	if err != nil {
		return nil, err
	}
	return c.do(ctx, req)
}
