package qiita

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteComment(t *testing.T) {
	// 204
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.DeleteComment(ctx, "")
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
		err := c.DeleteComment(ctx, "")
		if err == nil {
			t.Fail()
		}
	}()
}

func TestGetComment(t *testing.T) {
	// 200
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/get_comment.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.GetComment(ctx, "")
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 400
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/get_comment.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.GetComment(ctx, "")
		if err == nil {
			t.Fail()
		}
	}()
}

func TestUpdateComment(t *testing.T) {
	// 200
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.UpdateComment(ctx, Comment{})
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
		err := c.UpdateComment(ctx, Comment{})
		if err == nil {
			t.Fail()
		}
	}()
}

func TestListComments(t *testing.T) {
	// 200
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_comments.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.ListComments(ctx, "")
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 400
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_comments.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.ListComments(ctx, "")
		if err == nil {
			t.Fail()
		}
	}()
}

func TestPostComment(t *testing.T) {
	// 201
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			http.ServeFile(w, r, "testdata/post_comment.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.PostComment(ctx, "", Comment{})
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 400
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/post_comment.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.PostComment(ctx, "", Comment{})
		if err == nil {
			t.Fail()
		}
	}()
}
