package server

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRegistrationClientAllowsSlowJobStart(t *testing.T) {
	sidecar := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second)
		_ = json.NewEncoder(w).Encode(map[string]any{"ok": true, "batch_id": "batch-1"})
	}))
	defer sidecar.Close()

	client := registrationClient(Options{RegistrationURL: sidecar.URL})
	result, err := client.Start(context.Background(), map[string]any{"count": 1}, "")
	if err != nil || result["batch_id"] != "batch-1" {
		t.Fatalf("result=%#v err=%v", result, err)
	}
}
