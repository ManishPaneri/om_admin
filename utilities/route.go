package utilities

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"net/http"
)

func Mapping(r *http.Request, method string) bool {
	ID := r.URL.Query().Get("id")
	export := r.URL.Query().Get("export")
	Import := r.URL.Query().Get("import")
	switch true {
	case r.Method == "GET" && (ID == "" || ID == "0") && method == "GETALL" && (export == "" || export == "false"):
		spew.Dump("================ GET ALL API REQUEST ================", r, "======== END REQUEST ========")
		return true
		break

	case r.Method == "GET" && ID == "0" && method == "EXPORT" && export == "true":
		spew.Dump("================ EXPORT API REQUEST ================", r, "======== END REQUEST ========")
		return true
		break

	case r.Method == "GET" && ID != "" && method == "GETONE":
		spew.Dump("================ GET API REQUEST ================", r, "======== END REQUEST ========")
		return true
		break

	case r.Method == "POST" && method == "POST" && (Import == "" || Import == "false"):
		spew.Dump("================ POST API REQUEST ================", r, "======== END REQUEST ========")
		return true
		break

	case r.Method == "POST" && method == "IMPORT" && Import == "true":
		spew.Dump("================ IMPORT API REQUEST ================", r, "======== END REQUEST ========")
		return true
		break

	case r.Method == "PUT" && ID == "" && method == "PUT":
		spew.Dump("================ PUT API REQUEST ================", r, "======== END REQUEST ========")
		return true
		break
	}

	return false
}

func URLReturnResponseJson(w http.ResponseWriter, data interface{}) {
	permission := w.Header().Get("currentUser")
	UpdateLoggedApiByRequestID(data, permission)
	w.Header().Del("currentUser")
	spew.Dump("================ API RESPONSE ================", data, "======== END RESPONSE ========")
	returnJson, _ := json.Marshal(data)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(returnJson)
}

func URLReturnResponseCSV(w http.ResponseWriter, data interface{}) {
	permission := w.Header().Get("currentUser")
	UpdateLoggedApiByRequestID(data, permission)
	w.Header().Del("currentUser")
}
