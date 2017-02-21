package qiita

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func mockClient(server *httptest.Server) (*Client, error) {
	parsedURL, _ := url.ParseRequestURI(server.URL)
	return &Client{
		URL: parsedURL,
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
		Token: "",
	}, nil
}

func TestNewClient(t *testing.T) {
	config := NewConfig().WithEndpoint("http://qiita.com")
	_, err := NewClient("", *config)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewRequest(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	c, _ := mockClient(server)
	ctx := context.TODO()
	_, err := c.newRequest(ctx, "", "", nil)
	if err != nil {
		t.Fatal(err)
	}
}
