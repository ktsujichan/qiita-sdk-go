package qiita

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_GetProjects(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/get_projects.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListProjects(ctx, 1, 1)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/get_projects.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.ListProjects(ctx, 1, 1)
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_CreateProject(t *testing.T) {
	t.Run("201", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.CreateProject(ctx, Project{})
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
		err := c.CreateProject(ctx, Project{})
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_DeleteProject(t *testing.T) {
	t.Run("204", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.DeleteProject(ctx, 1)
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
		err := c.DeleteProject(ctx, 1)
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_GetProject(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "testdata/get_project.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.GetProject(ctx, "", 1, 1)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("400", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "testdata/get_project.json")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		_, err := c.GetProject(ctx, "", 1, 1)
		if err == nil {
			t.Fail()
		}
	})
}

func TestClient_UpdateProject(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			http.ServeFile(w, r, "")
		}))
		c, _ := mockClient(server)
		ctx := context.TODO()
		err := c.UpdateProject(ctx, Project{})
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
		err := c.UpdateProject(ctx, Project{})
		if err == nil {
			t.Fail()
		}
	})
}
