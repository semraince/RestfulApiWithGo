package pairs

import (
	"GoDatabaseAssignment/httpUtil"
	"fmt"
	"net/http"
)

func KeyValuePairs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getKeyValuePair(w, r)
	case http.MethodPost:
		saveKeyValuePair(w, r)
	default:
		// unknown method
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
func getKeyValuePair(w http.ResponseWriter, r *http.Request) {
	key := httpUtil.GetQueryParam(r, "key")
	result, err := GetKeyValuePair(key)
	if err != nil {
		httpUtil.GenerateResponse(w, r, nil, createError(err))
		return
	}

	httpUtil.GenerateResponse(w, r, err, result)
}

func saveKeyValuePair(w http.ResponseWriter, r *http.Request) {
	var data KeyValue
	err := httpUtil.GetBody(r, &data)
	if nil != err {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := SavePair(data)
	if err != nil {
		httpUtil.GenerateResponse(w, r, nil, createError(err))
		return
	}
	httpUtil.GenerateResponse(w, r, nil, result)
}

func createError(err error) ErrorMessage {
	errorMessage := ErrorMessage{Message: fmt.Sprint(err)}
	return errorMessage
}
