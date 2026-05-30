package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"src/projector/db"
	"src/projector/models"
)

func setupTestDB(t *testing.T) {
	t.Helper()
	var err error
	db.DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}
	if err := db.DB.AutoMigrate(&models.Order{}); err != nil {
		t.Fatalf("failed to migrate test db: %v", err)
	}
}

func newRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/orders", CreateOrder)
	return r
}

func TestCreateOrder_Success(t *testing.T) {
	setupTestDB(t)

	body, _ := json.Marshal(map[string]any{
		"client_name":   "Acme Corp",
		"project_type":  "Web",
		"delivery_date": "2026-12-01T00:00:00Z",
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	newRouter().ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d: %s", w.Code, w.Body.String())
	}

	var resp map[string]any
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp["id"] == "" {
		t.Error("expected non-empty id")
	}
	if resp["status"] != "Pending" {
		t.Errorf("expected status Pending, got %v", resp["status"])
	}
}

func TestCreateOrder_MissingClientName(t *testing.T) {
	setupTestDB(t)

	body, _ := json.Marshal(map[string]any{
		"project_type":  "Web",
		"delivery_date": "2026-12-01T00:00:00Z",
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	newRouter().ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d: %s", w.Code, w.Body.String())
	}
}

func TestCreateOrder_PastDeliveryDate(t *testing.T) {
	setupTestDB(t)

	body, _ := json.Marshal(map[string]any{
		"client_name":   "Acme Corp",
		"project_type":  "Web",
		"delivery_date": "2020-01-01T00:00:00Z",
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	newRouter().ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d: %s", w.Code, w.Body.String())
	}

	var resp map[string]any
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp["error"] != "delivery_date must be in the future" {
		t.Errorf("unexpected error message: %v", resp["error"])
	}
}

func TestCreateOrder_Duplicate(t *testing.T) {
	setupTestDB(t)

	body, _ := json.Marshal(map[string]any{
		"client_name":   "Acme Corp",
		"project_type":  "Web",
		"delivery_date": "2026-12-01T00:00:00Z",
	})

	newRouter().ServeHTTP(httptest.NewRecorder(), func() *http.Request {
		r := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(body))
		r.Header.Set("Content-Type", "application/json")
		return r
	}())

	body, _ = json.Marshal(map[string]any{
		"client_name":   "Acme Corp",
		"project_type":  "Web",
		"delivery_date": "2026-12-01T00:00:00Z",
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	newRouter().ServeHTTP(w, req)

	if w.Code != http.StatusConflict {
		t.Errorf("expected 409, got %d: %s", w.Code, w.Body.String())
	}
}
