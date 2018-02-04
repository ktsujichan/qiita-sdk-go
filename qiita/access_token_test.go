package qiita

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_CreateAccessToken(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.CreateAccessToken(ctx, Auth{})
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.CreateAccessToken(ctx, Auth{})
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_DeleteAccessToken(t *testing.T) {
	t.Run("204", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.DeleteAccessToken(ctx, "")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.DeleteAccessToken(ctx, "")
		if err == nil {
			t.Fail()
		}
	})
}
