package qiita

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddCommentReaction(t *testing.T) {
	// 201
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			http.ServeFile(w, r, "testdata/add_comment_reaction.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.AddCommentReaction(ctx, "", Reaction{})
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 400
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/add_comment_reaction.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.AddCommentReaction(ctx, "", Reaction{})
		if err == nil {
			t.Fail()
		}
	}()
}

func TestAddItemReaction(t *testing.T) {
	// 201
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			http.ServeFile(w, r, "testdata/add_item_reaction.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.AddItemReaction(ctx, "", Reaction{})
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 400
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/add_item_reaction.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.AddItemReaction(ctx, "", Reaction{})
		if err == nil {
			t.Fail()
		}
	}()
}

func TestAddProjectReaction(t *testing.T) {
	// 201
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			http.ServeFile(w, r, "testdata/add_project_reaction.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.AddProjectReaction(ctx, 1, Reaction{})
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 400
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/add_project_reaction.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.AddProjectReaction(ctx, 1, Reaction{})
		if err == nil {
			t.Fail()
		}
	}()
}

func TestDeleteCommentReaction(t *testing.T) {
	// 204
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.DeleteCommentReaction(ctx, "", "")
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
		err := c.DeleteCommentReaction(ctx, "", "")
		if err == nil {
			t.Fail()
		}
	}()
}

func TestDeleteItemReaction(t *testing.T) {
	// 204
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.DeleteItemReaction(ctx, "", "")
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
		err := c.DeleteItemReaction(ctx, "", "")
		if err == nil {
			t.Fail()
		}
	}()
}

func TestDeleteProjectReaction(t *testing.T) {
	// 204
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.DeleteProjectReaction(ctx, 1, "")
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
		err := c.DeleteProjectReaction(ctx, 1, "")
		if err == nil {
			t.Fail()
		}
	}()
}

func TestListCommentReactions(t *testing.T) {
	// 200
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_comment_reactions.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.ListCommentReactions(ctx, "")
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 400
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_comment_reactions.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.ListCommentReactions(ctx, "")
		if err == nil {
			t.Fail()
		}
	}()
}

func TestListItemReactions(t *testing.T) {
	// 200
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_item_reactions.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.ListItemReactions(ctx, "")
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 400
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_item_reactions.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.ListItemReactions(ctx, "")
		if err == nil {
			t.Fail()
		}
	}()
}

func TestListProjectReactions(t *testing.T) {
	// 200
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_project_reactions.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.ListProjectReactions(ctx, "")
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 400
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_project_reactions.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.ListProjectReactions(ctx, "")
		if err == nil {
			t.Fail()
		}
	}()
}
