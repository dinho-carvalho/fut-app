package httprespond

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJSON_WritesStatusHeaderAndBody(t *testing.T) {
	rr := httptest.NewRecorder()
	payload := map[string]string{"ok": "yes"}

	if err := JSON(rr, http.StatusCreated, payload); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected %d, got %d", http.StatusCreated, rr.Code)
	}
	if ct := rr.Header().Get("Content-Type"); ct != "application/json" {
		t.Fatalf("expected application/json, got %q", ct)
	}
	var got map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &got); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	if got["ok"] != "yes" {
		t.Fatalf("unexpected body: %v", got)
	}
}
