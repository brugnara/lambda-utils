package proxy

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
)

func TestGetProxyResponseCORS(t *testing.T) {
	body := fmt.Sprintf("body - %v", uuid.New())

	// ok
	responseOK := NewResponse(nil, body, true)
	if responseOK.Body != body {
		t.Error("Wrong body")
		return
	}
	if responseOK.Headers["Access-Control-Allow-Origin"] != "*" {
		t.Error("Wrong header CORS")
		return
	}
	if responseOK.StatusCode != http.StatusOK {
		t.Error("Wrong StatusCode")
		return
	}

	// ko
	responseKO := NewResponse(errors.New("random error"), body, true)
	if responseKO.Body != body {
		t.Error("Wrong body")
		return
	}
	if responseKO.Headers["Access-Control-Allow-Origin"] != "*" {
		t.Error("Wrong header CORS")
		return
	}
	if responseKO.StatusCode != http.StatusInternalServerError {
		t.Error("Wrong StatusCode")
		return
	}
}

func TestGetProxyResponseNoCORS1(t *testing.T) {
	body := fmt.Sprintf("body - %v", uuid.New())

	// ok
	responseOK := NewResponse(nil, body)
	if responseOK.Body != body {
		t.Error("Wrong body")
		return
	}
	if responseOK.Headers["Access-Control-Allow-Origin"] == "*" {
		t.Error("Unexpected CORS header")
		return
	}
	if responseOK.StatusCode != http.StatusOK {
		t.Error("Wrong StatusCode")
		return
	}

	// ko
	responseKO := NewResponse(errors.New("random error"), body)
	if responseKO.Body != body {
		t.Error("Wrong body")
		return
	}
	if responseKO.Headers["Access-Control-Allow-Origin"] == "*" {
		t.Error("Unexpected CORS header")
		return
	}
	if responseKO.StatusCode != http.StatusInternalServerError {
		t.Error("Wrong StatusCode")
		return
	}
}
func TestGetProxyResponseNoCORS2(t *testing.T) {
	body := fmt.Sprintf("body - %v", uuid.New())

	// ok
	responseOK := NewResponse(nil, body, false)
	if responseOK.Body != body {
		t.Error("Wrong body")
		return
	}
	if responseOK.Headers["Access-Control-Allow-Origin"] == "*" {
		t.Error("Unexpected CORS header")
		return
	}
	if responseOK.StatusCode != http.StatusOK {
		t.Error("Wrong StatusCode")
		return
	}

	// ko
	responseKO := NewResponse(errors.New("random error"), body, false)
	if responseKO.Body != body {
		t.Error("Wrong body")
		return
	}
	if responseKO.Headers["Access-Control-Allow-Origin"] == "*" {
		t.Error("Unexpected CORS header")
		return
	}
	if responseKO.StatusCode != http.StatusInternalServerError {
		t.Error("Wrong StatusCode")
		return
	}
}
