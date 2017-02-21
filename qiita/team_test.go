package qiita

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListTeams(t *testing.T) {
	// 200
	func() {
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
	}()

	// 400
	func() {
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
	}()
}
