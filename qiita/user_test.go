package qiita

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_ListStockers(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_stockers.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListStockers(ctx, "", 1, 1)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_stockers.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListStockers(ctx, "", 1, 1)
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_ListUsers(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_users.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListUsers(ctx, 1, 1)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_users.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListUsers(ctx, 1, 1)
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_GetUser(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/get_user.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.GetUser(ctx, "")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/get_user.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.GetUser(ctx, "")
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_ListFollowees(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_followees.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListFollowees(ctx, "", 1, 1)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_followees.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListFollowees(ctx, "", 1, 1)
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_ListFollowers(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_followers.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListFollowers(ctx, "", 1, 1)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_followers.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListFollowers(ctx, "", 1, 1)
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_UnfollowUser(t *testing.T) {
	t.Run("204", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.UnfollowUser(ctx, "")
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
		err := c.UnfollowUser(ctx, "")
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_EnsureFollowingUser(t *testing.T) {
	t.Run("204", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.EnsureFollowingUser(ctx, "")
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
		err := c.EnsureFollowingUser(ctx, "")
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_FollowUser(t *testing.T) {
	t.Run("204", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.FollowUser(ctx, "")
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
		err := c.FollowUser(ctx, "")
		if err == nil {
			t.Fail()
		}
	})
}
