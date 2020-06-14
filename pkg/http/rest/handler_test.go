package rest

import (
	"bytes"
	"github.com/zilohumberto/notification-sns-go/pkg/sending"
	"io"
	"net/http"

	"net/http/httptest"
	"testing"
)

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestPing(t *testing.T){
	sender := sending.NewService()
	mapUrls(sender)
	w := performRequest(router, "GET", "/ping", nil)
	if http.StatusOK != w.Code{
		t.Errorf("Status code error. expected 200. got %v", w.Code)
		return
	}
	if w.Body.String() != "pong"{
		t.Errorf("Body response error. expected ping. got %v", w.Body)
	}
}

func TestHandler(t *testing.T) {
	var rawBody = []byte(`{
		"target": "arn:aws:sns:region:device-arg",
		"message": "{\"default\": \"we are in live\"}",
		"subject": "SOME_ACTION"
	}`)
	body := bytes.NewBuffer(rawBody)
	w := performRequest(router, "POST", "/push", body)

	if w.Code != 200{
		t.Errorf("Error Status code. Excepted 200. got %d", w.Code)
		return
	}
	if w.Body.String() != "\"Notification sent\""{
		t.Errorf("Body response error. expected \"Notification sent\". got %v", w.Body.String())
	}
}
