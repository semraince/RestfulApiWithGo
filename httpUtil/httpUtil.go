package httpUtil

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type contextKey struct {
	name string
}

// StatusCtxKey is a context key to record a future HTTP response status code.
var StatusCtxKey = &contextKey{"Status"}

func GenerateResponse(w http.ResponseWriter, r *http.Request, err error, result interface{}) {
	if nil != err {
		w.WriteHeader(http.StatusServiceUnavailable)
		JSON(w, r, err)
		return
	}
	if nil == result {
		result = struct {
			Message string
		}{Message: "success"}
	}
	JSON(w, r, result)
}

func GetQueryParam(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

func GetBody(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(&v)
}

func JSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if status, ok := r.Context().Value(StatusCtxKey).(int); ok {
		w.WriteHeader(status)
	}
	w.Write(buf.Bytes())
}
