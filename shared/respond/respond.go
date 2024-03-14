package respond

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// Respond writes the given status code, content type and body to the given http.ResponseWriter
func Respond(w http.ResponseWriter, status int, contentType string, body []byte) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(status)
	_, _ = w.Write(body)
}

// JSON writes the given status code and value as a JSON response
func JSON(w http.ResponseWriter, status int, v any) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)

	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Respond(w, status, "application/json", buf.Bytes())
}

// Err writes an error response
func Err(w http.ResponseWriter, err error) {
	if rerr, ok := err.(Error); ok {
		JSON(w, rerr.StatusCode(), map[string]any{"error": rerr.Error()})
		return
	}

	log.Printf("Unhandled error: %s", err)

	JSON(w, 500, map[string]any{"error": "Something went wrong. Please retry later"})
}
