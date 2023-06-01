package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBarCodeRoute_Return_Ok(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	body := []byte(`{
		"content": "Test request",
		"width": 100,
		"height": 100
	  }`)
	req, _ := http.NewRequest("POST", "/generate/barcode", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestQrCodeRoute_Return_Ok(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	body := []byte(`{
		"content": "Test request",
		"width": 100,
		"height": 100
	  }`)
	req, _ := http.NewRequest("POST", "/generate/qrcode", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestBarCodeRoute_Return_BadRequest_Content_IsEmpty(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	body := []byte(`{
		"content": "",
		"width": 100,
		"height": 100
	  }`)
	req, _ := http.NewRequest("POST", "/generate/barcode", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

}
func TestBarCodeRoute_Return_BadRequest_Body_IsEmpty(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	body := []byte(`{}`)
	req, _ := http.NewRequest("POST", "/generate/barcode", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

}

func TestQrCodeRoute_Return_BadRequest_Content_IsEmpty(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	body := []byte(`{
		"content": "",
		"width": 100,
		"height": 100
	  }`)
	req, _ := http.NewRequest("POST", "/generate/qrcode", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

}
func TestQrCodeRoute_Return_BadRequest_Body_IsEmpty(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	body := []byte(`{ }`)
	req, _ := http.NewRequest("POST", "/generate/qrcode", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

}
