package qiita

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_ListTags(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_tags.json")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListTags(ctx, 1, 1, "")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_tags.json")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListTags(ctx, 1, 1, "")
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_GetTag(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/get_tag.json")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.GetTag(ctx, "")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/get_tag.json")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.GetTag(ctx, "")
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_ListFollowingTags(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_following_tags.json")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListFollowingTags(ctx, "", 1, 1)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_following_tags.json")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListFollowingTags(ctx, "", 1, 1)
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_UnfollowTag(t *testing.T) {
	t.Run("204", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.UnfollowTag(ctx, "")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.UnfollowTag(ctx, "")
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_EnsureFollowingTag(t *testing.T) {
	t.Run("204", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.EnsureFollowingTag(ctx, "")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.EnsureFollowingTag(ctx, "")
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_FollowTag(t *testing.T) {
	t.Run("204", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.FollowTag(ctx, "")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.FollowTag(ctx, "")
		if err == nil {
			t.Fail()
		}
	})
}
