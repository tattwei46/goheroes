package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetHeroes(t *testing.T) {
	req, err := http.NewRequest("GET", "/heroes", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetHeroes)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status ok; got %v", rr.Code)
	}

	expected := `[{"id":"1","name":"Ironman"}]`
	result := strings.TrimSpace(rr.Body.String())
	assert.Equal(t, expected, result, "handle returned unexpected body, got %v, want %v", result, expected)

}

func TestGetHeroById(t *testing.T) {
	req, err := http.NewRequest("GET", "/hero", nil)

	vars := map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetHero)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status ok; got %v", rr.Code)
	}

	expected := `{"id":"1","name":"Ironman"}`
	result := strings.TrimSpace(rr.Body.String())
	assert.Equal(t, expected, result, "handle returned unexpected body, got %v, want %v", result, expected)

}
