package api

import (
	"bytes"
	"mime/multipart"
	"net/http/httptest"
	"testing"
)

func TestHandleIngestAcceptsFile(t *testing.T) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	part, _ := writer.CreateFormFile("batch", "test.csv")
	part.Write([]byte("id,note\n1,15"))
	writer.Close()

	req := httptest.NewRequest("POST", "/ingest", &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()

	handleIngest(w, req)

	if w.Code != 202 {
		t.Errorf("attendu 202, obtenu %d", w.Code)
	}
}