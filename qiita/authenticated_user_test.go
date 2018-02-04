package qiita

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_GetAuthenticatedUser(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/get_authenticated_user.json")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.GetAuthenticatedUser(ctx)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/get_authenticated_user.json")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.GetAuthenticatedUser(ctx)
		if err == nil {
			t.Fail()
		}
	})
}
