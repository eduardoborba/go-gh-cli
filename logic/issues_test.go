package logic

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchIssues(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/repos/rails/rails/issues" {
			t.Errorf("Expected to request '/repos/rails/rails/issues', got: %s", r.URL.Path)
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept: application/json header, got: %s", r.Header.Get("Accept"))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id":1847156211,"url":"https://api.github.com/repos/rails/rails/issues/48930"}]`))
	}))
	defer server.Close()

	issues, _ := FetchIssues(server.URL)
	if len(issues) != 1 {
		t.Errorf("Expected 1 issue, got %d", len(issues))
	}

	if issues[0].ID != 1847156211 {
		t.Errorf("Expected id 1847156211, got %d", issues[0].ID)
	}
}
