package qiita

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListStockers(t *testing.T) {
	// 200
	func() {
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
	}()

	// 400
	func() {
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
	}()
}

func TestListUsers(t *testing.T) {
	// 200
	func() {
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
	}()

	// 400
	func() {
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
	}()
}

func TestGetUser(t *testing.T) {
	// 200
	func() {
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
	}()

	// 400
	func() {
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
	}()
}

func TestListFollowees(t *testing.T) {
	// 200
	func() {
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
	}()

	// 400
	func() {
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
	}()
}

func TestListFollowers(t *testing.T) {
	// 200
	func() {
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
	}()

	// 400
	func() {
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
	}()
}

func TestUnfollowUser(t *testing.T) {
	// 204
	func() {
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
	}()

	// 400
	func() {
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
	}()
}

func TestEnsureFollowingUser(t *testing.T) {
	// 204
	func() {
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
	}()

	// 400
	func() {
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
	}()
}

func TestFollowUser(t *testing.T) {
	// 204
	func() {
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
	}()

	// 400
	func() {
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
	}()
}
