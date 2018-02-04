package qiita

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_DeleteComment(t *testing.T) {
	t.Run("204", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.DeleteComment(ctx, "")
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
		err := c.DeleteComment(ctx, "")
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_GetComment(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/get_comment.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.GetComment(ctx, "")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/get_comment.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.GetComment(ctx, "")
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_UpdateComment(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.UpdateComment(ctx, Comment{})
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
		err := c.UpdateComment(ctx, Comment{})
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_ListComments(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_comments.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListComments(ctx, "")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_comments.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListComments(ctx, "")
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_PostComment(t *testing.T) {
	t.Run("201", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			http.ServeFile(w, r, "testdata/post_comment.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.PostComment(ctx, "", Comment{})
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/post_comment.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.PostComment(ctx, "", Comment{})
		if err == nil {
			t.Fail()
		}
	})
}
