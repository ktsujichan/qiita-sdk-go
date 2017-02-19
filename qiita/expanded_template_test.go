package qiita

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateExpandedTemplate(t *testing.T) {
	// 201
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			http.ServeFile(w, r, "testdata/create_expanded_template.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.CreateExpandedTemplate(ctx, Template{})
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 400
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/create_expanded_template.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.CreateExpandedTemplate(ctx, Template{})
		if err == nil {
			t.Fail()
		}
	}()
}
