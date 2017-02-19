package qiita

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProjects(t *testing.T) {
	// 200
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/get_projects.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.ListProjects(ctx, 1, 1)
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 400
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/get_projects.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.ListProjects(ctx, 1, 1)
		if err == nil {
			t.Fail()
		}
	}()
}

func TestCreateProject(t *testing.T) {
	// 201
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.CreateProject(ctx, Project{})
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
		err := c.CreateProject(ctx, Project{})
		if err == nil {
			t.Fail()
		}
	}()
}

func TestDeleteProject(t *testing.T) {
	// 204
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.DeleteProject(ctx, 1)
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
		err := c.DeleteProject(ctx, 1)
		if err == nil {
			t.Fail()
		}
	}()
}

func TestGetProject(t *testing.T) {
	// 200
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/get_project.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.GetProject(ctx, "", 1, 1)
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 400
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/get_project.json")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		_, err := c.GetProject(ctx, "", 1, 1)
		if err == nil {
			t.Fail()
		}
	}()
}

func TestUpdateProject(t *testing.T) {
	// 200
	func() {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx, _ := context.WithCancel(context.Background())
		err := c.UpdateProject(ctx, Project{})
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
		err := c.UpdateProject(ctx, Project{})
		if err == nil {
			t.Fail()
		}
	}()
}
