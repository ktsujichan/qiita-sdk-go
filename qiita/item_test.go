package qiita

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_ListAuthenticatedUserItems(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_authenticated_user_items.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListAuthenticatedUserItems(ctx, 1, 1)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_authenticated_user_items.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListAuthenticatedUserItems(ctx, 1, 1)
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_ListItems(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_items.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListItems(ctx, 1, 1, "")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_items.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListItems(ctx, 1, 1, "")
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_CreateItem(t *testing.T) {
	t.Run("201", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.CreateItem(ctx, Item{})
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
		err := c.CreateItem(ctx, Item{})
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_DeleteItem(t *testing.T) {
	t.Run("204", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.DeleteItem(ctx, "")
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
		err := c.DeleteItem(ctx, "")
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_GetItem(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/get_item.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.GetItem(ctx, "")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/get_item.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.GetItem(ctx, "")
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_UpdateItem(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/update_item.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.UpdateItem(ctx, Item{})
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/update_item.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.UpdateItem(ctx, Item{})
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_UnlikeItem(t *testing.T) {
	t.Run("204", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.UnlikeItem(ctx, "")
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
		err := c.UnlikeItem(ctx, "")
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_LikeItem(t *testing.T) {
	t.Run("204", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.LikeItem(ctx, "")
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
		err := c.LikeItem(ctx, "")
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_StockItem(t *testing.T) {
	t.Run("204", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.StockItem(ctx, "")
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
		err := c.StockItem(ctx, "")
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_UnstockItem(t *testing.T) {
	t.Run("204", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.UnstockItem(ctx, "")
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
		err := c.UnstockItem(ctx, "")
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_EnsureItemStock(t *testing.T) {
	t.Run("204", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.EnsureItemStock(ctx, "")
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
		err := c.EnsureItemStock(ctx, "")
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_EnsureItemLike(t *testing.T) {
	t.Run("204", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.EnsureItemLike(ctx, "")
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
		err := c.EnsureItemLike(ctx, "")
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_ListTaggedItems(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_tagged_items.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListTaggedItems(ctx, "", 1, 1)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_tagged_items.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListTaggedItems(ctx, "", 1, 1)
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_ListUserItems(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_user_items.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListUserItems(ctx, "", 1, 1)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_user_items.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListUserItems(ctx, "", 1, 1)
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_ListUserStocks(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_user_stocks.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListUserStocks(ctx, "", 1, 1)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_user_stocks.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListUserStocks(ctx, "", 1, 1)
		if err == nil {
			t.Fail()
		}
	})
}
