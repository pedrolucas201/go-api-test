package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, recebido %d", res.StatusCode)
	}
}

func TestMessageHandler(t *testing.T) {
	body := Message{
		Text: "Teste unitário",
	}

	jsonBody, err := json.Marshal(body); 
	
	if err != nil { t.Fatal(err) }

	req := httptest.NewRequest(
		http.MethodPost,
		"/message",
		bytes.NewBuffer(jsonBody),
	)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	messageHandler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Errorf("esperado status 201, recebido %d", res.StatusCode)
	}

	var response Message
	json.NewDecoder(res.Body).Decode(&response)

	if response.Text != body.Text {
		t.Errorf("esperado %s, recebido %s", body.Text, response.Text)
	}
}
