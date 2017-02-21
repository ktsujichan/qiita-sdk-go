package qiita

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListTemplates(t *testing.T) {
	// 200
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/list_templates.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListTemplates(ctx, 1, 1)
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 400
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/list_templates.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListTemplates(ctx, 1, 1)
		if err == nil {
			t.Fail()
		}
	}()
}

func TestDeleteTemplate(t *testing.T) {
	// 204
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.DeleteTemplate(ctx, 1)
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
		err := c.DeleteTemplate(ctx, 1)
		if err == nil {
			t.Fail()
		}
	}()
}

func TestGetTemplate(t *testing.T) {
	// 200
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/get_template.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.GetTemplate(ctx, 1)
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 400
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/get_template.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.GetTemplate(ctx, 1)
		if err == nil {
			t.Fail()
		}
	}()
}

func TestCreateTemplate(t *testing.T) {
	// 201
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.CreateTemplate(ctx, Template{})
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
		err := c.CreateTemplate(ctx, Template{})
		if err == nil {
			t.Fail()
		}
	}()
}

func TestUpdateTemplate(t *testing.T) {
	// 200
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.UpdateTemplate(ctx, Template{})
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
		err := c.UpdateTemplate(ctx, Template{})
		if err == nil {
			t.Fail()
		}
	}()
}
