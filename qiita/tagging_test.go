package qiita

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddItemTagging(t *testing.T) {
	// 201
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.AddItemTagging(ctx, "", Tagging{})
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 400
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.AddItemTagging(ctx, "", Tagging{})
		if err == nil {
			t.Fail()
		}
	}()
}

func TestDeleteItemTagging(t *testing.T) {
	// 204
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.DeleteItemTagging(ctx, "", "")
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 400
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.DeleteItemTagging(ctx, "", "")
		if err == nil {
			t.Fail()
		}
	}()
}
