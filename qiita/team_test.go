package qiita

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_ListTeams(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_teams.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListTeams(ctx)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_teams.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListTeams(ctx)
		if err == nil {
			t.Fail()
		}
	})
}
