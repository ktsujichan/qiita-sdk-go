package qiita

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_CreateExpandedTemplate(t *testing.T) {
	t.Run("201", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			http.ServeFile(w, r, "testdata/create_expanded_template.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.CreateExpandedTemplate(ctx, Template{})
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/create_expanded_template.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.CreateExpandedTemplate(ctx, Template{})
		if err == nil {
			t.Fail()
		}
	})
}
