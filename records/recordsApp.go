package records

import (
	"GoDatabaseAssignment/httpUtil"
	"net/http"
)

func Records(w http.ResponseWriter, r *http.Request) {
	var data RequestRecord
	err := httpUtil.GetBody(r, &data)
	if nil != err {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	httpUtil.GenerateResponse(w, r, err, GetRecords(data))
}
