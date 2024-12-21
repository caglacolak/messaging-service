package tests

import (
	"net/http"
	"net/http/httptest"
	"project/internal/api"
	"testing"
)

func TestGetSentMessages(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/sent-messages", nil)
	w := httptest.NewRecorder()
	api.GetSentMessages(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %v", w.Code)
	}

	// Test getting sent messages again should return bad request
	w = httptest.NewRecorder()
	api.GetSentMessages(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status BadRequest, got %v", w.Code)
	}
}

func TestStartSchedulerOnce(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/start-messages", nil)
	w := httptest.NewRecorder()
	api.StartScheduler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %v", w.Code)
	}

	// Test starting the scheduler again should return bad request
	w = httptest.NewRecorder()
	api.StartScheduler(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status BadRequest, got %v", w.Code)
	}
}

func TestStopScheduler(t *testing.T) {
	// Start the scheduler first
	req := httptest.NewRequest(http.MethodGet, "/start-messages", nil)
	w := httptest.NewRecorder()
	api.StartScheduler(w, req)

	// Now stop the scheduler
	req = httptest.NewRequest(http.MethodGet, "/stop-messages", nil)
	w = httptest.NewRecorder()
	api.StopScheduler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %v", w.Code)
	}

	// Test stopping the scheduler again should return bad request
	w = httptest.NewRecorder()
	api.StopScheduler(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status BadRequest, got %v", w.Code)
	}
}

func TestStopSchedulerWithoutStarting(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/stop-messages", nil)
	w := httptest.NewRecorder()
	api.StopScheduler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status BadRequest, got %v", w.Code)
	}
}
