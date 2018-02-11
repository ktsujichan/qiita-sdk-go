package qiita

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_AddCommentReaction(t *testing.T) {
	t.Run("201", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			http.ServeFile(w, r, "testdata/add_comment_reaction.json")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.AddCommentReaction(ctx, "", Reaction{})
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/add_comment_reaction.json")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.AddCommentReaction(ctx, "", Reaction{})
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_AddItemReaction(t *testing.T) {
	t.Run("201", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			http.ServeFile(w, r, "testdata/add_item_reaction.json")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.AddItemReaction(ctx, "", Reaction{})
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/add_item_reaction.json")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.AddItemReaction(ctx, "", Reaction{})
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_AddProjectReaction(t *testing.T) {
	t.Run("201", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			http.ServeFile(w, r, "testdata/add_project_reaction.json")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.AddProjectReaction(ctx, 1, Reaction{})
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/add_project_reaction.json")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.AddProjectReaction(ctx, 1, Reaction{})
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_DeleteCommentReaction(t *testing.T) {
	t.Run("204", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.DeleteCommentReaction(ctx, "", "")
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
		err := c.DeleteCommentReaction(ctx, "", "")
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_DeleteItemReaction(t *testing.T) {
	t.Run("204", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.DeleteItemReaction(ctx, "", "")
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
		err := c.DeleteItemReaction(ctx, "", "")
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_DeleteProjectReaction(t *testing.T) {
	t.Run("204", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.DeleteProjectReaction(ctx, 1, "")
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
		err := c.DeleteProjectReaction(ctx, 1, "")
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_ListCommentReactions(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_comment_reactions.json")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListCommentReactions(ctx, "")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_comment_reactions.json")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListCommentReactions(ctx, "")
		if err == nil {
			t.Fail()
		}
	})
}

func TestListItemReactions(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_item_reactions.json")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListItemReactions(ctx, "")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_item_reactions.json")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListItemReactions(ctx, "")
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_ListProjectReactions(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_project_reactions.json")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListProjectReactions(ctx, "")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_project_reactions.json")
		}))
		defer server.Close()
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListProjectReactions(ctx, "")
		if err == nil {
			t.Fail()
		}
	})
}
