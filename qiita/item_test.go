package qiita

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListAuthenticatedUserItems(t *testing.T) {
	// 200
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_authenticated_user_items.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.ListAuthenticatedUserItems(ctx, 1, 1)
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 400
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_authenticated_user_items.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.ListAuthenticatedUserItems(ctx, 1, 1)
		if err == nil {
			t.Fail()
		}
	}()
}

func TestListItems(t *testing.T) {
	// 200
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_items.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.ListItems(ctx, 1, 1, "")
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 400
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_items.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.ListItems(ctx, 1, 1, "")
		if err == nil {
			t.Fail()
		}
	}()
}

func TestCreateItem(t *testing.T) {
	// 201
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.CreateItem(ctx, Item{})
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
		err := c.CreateItem(ctx, Item{})
		if err == nil {
			t.Fail()
		}
	}()
}

func TestDeleteItem(t *testing.T) {
	// 204
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.DeleteItem(ctx, "")
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
		err := c.DeleteItem(ctx, "")
		if err == nil {
			t.Fail()
		}
	}()
}

func TestGetItem(t *testing.T) {
	// 200
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/get_item.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.GetItem(ctx, "")
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 400
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/get_item.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.GetItem(ctx, "")
		if err == nil {
			t.Fail()
		}
	}()
}

func TestUpdateItem(t *testing.T) {
	// 200
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/update_item.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.UpdateItem(ctx, Item{})
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 400
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/update_item.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.UpdateItem(ctx, Item{})
		if err == nil {
			t.Fail()
		}
	}()
}

func TestUnlikeItem(t *testing.T) {
	// 204
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.UnlikeItem(ctx, "")
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
		err := c.UnlikeItem(ctx, "")
		if err == nil {
			t.Fail()
		}
	}()
}

func TestLikeItem(t *testing.T) {
	// 204
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.LikeItem(ctx, "")
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
		err := c.LikeItem(ctx, "")
		if err == nil {
			t.Fail()
		}
	}()
}

func TestStockItem(t *testing.T) {
	// 204
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.StockItem(ctx, "")
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
		err := c.StockItem(ctx, "")
		if err == nil {
			t.Fail()
		}
	}()
}

func TestUnstockItem(t *testing.T) {
	// 204
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.UnstockItem(ctx, "")
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
		err := c.UnstockItem(ctx, "")
		if err == nil {
			t.Fail()
		}
	}()
}

func TestEnsureItemStock(t *testing.T) {
	// 204
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.EnsureItemStock(ctx, "")
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
		err := c.EnsureItemStock(ctx, "")
		if err == nil {
			t.Fail()
		}
	}()
}

func TestEnsureItemLike(t *testing.T) {
	// 204
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.EnsureItemLike(ctx, "")
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
		err := c.EnsureItemLike(ctx, "")
		if err == nil {
			t.Fail()
		}
	}()
}

func TestListTaggedItems(t *testing.T) {
	// 200
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_tagged_items.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.ListTaggedItems(ctx, "", 1, 1)
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 400
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_tagged_items.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.ListTaggedItems(ctx, "", 1, 1)
		if err == nil {
			t.Fail()
		}
	}()
}

func TestListUserItems(t *testing.T) {
	// 200
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_user_items.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.ListUserItems(ctx, "", 1, 1)
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 400
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_user_items.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.ListUserItems(ctx, "", 1, 1)
		if err == nil {
			t.Fail()
		}
	}()
}

func TestListUserStocks(t *testing.T) {
	// 200
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_user_stocks.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.ListUserStocks(ctx, "", 1, 1)
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 400
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_user_stocks.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.ListUserStocks(ctx, "", 1, 1)
		if err == nil {
			t.Fail()
		}
	}()
}
