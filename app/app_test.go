package app

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWindowControlsAreDirectlyWired(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	Handler(rec, req)
	body := rec.Body.String()

	for _, action := range []string{"close", "minimize", "maximize"} {
		needle := `class="window-control-button ` + action + `"`
		if !strings.Contains(body, needle) {
			t.Fatalf("missing %s window control", action)
		}
	}
	if count := strings.Count(body, "CharlexWindowControls.action(this,event)"); count < 18 {
		t.Fatalf("expected direct control handlers on static windows, found %d", count)
	}
}
