package qiita

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListTags(t *testing.T) {
	// 200
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_tags.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListTags(ctx, 1, 1, "")
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 400
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_tags.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListTags(ctx, 1, 1, "")
		if err == nil {
			t.Fail()
		}
	}()
}

func TestGetTag(t *testing.T) {
	// 200
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/get_tag.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.GetTag(ctx, "")
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 400
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/get_tag.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.GetTag(ctx, "")
		if err == nil {
			t.Fail()
		}
	}()
}

func TestListFollowingTags(t *testing.T) {
	// 200
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_following_tags.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListFollowingTags(ctx, "", 1, 1)
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 400
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_following_tags.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListFollowingTags(ctx, "", 1, 1)
		if err == nil {
			t.Fail()
		}
	}()
}

func TestUnfollowTag(t *testing.T) {
	// 204
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.UnfollowTag(ctx, "")
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
		err := c.UnfollowTag(ctx, "")
		if err == nil {
			t.Fail()
		}
	}()
}

func TestEnsureFollowingTag(t *testing.T) {
	// 204
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.EnsureFollowingTag(ctx, "")
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
		err := c.EnsureFollowingTag(ctx, "")
		if err == nil {
			t.Fail()
		}
	}()
}

func TestFollowTag(t *testing.T) {
	// 204
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.FollowTag(ctx, "")
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
		err := c.FollowTag(ctx, "")
		if err == nil {
			t.Fail()
		}
	}()
}
