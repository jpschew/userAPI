package api

import (
	"encoding/json"
	"net/http"
	"sync"
	"userAPI/database"
)

func unprocessableEntityStatus(res http.ResponseWriter) {

	// set header to context-type of json format
	res.Header().Set("Content-Type", "application/json")

	// write status code to header
	res.WriteHeader(http.StatusUnprocessableEntity)

	resp := make(map[string]interface{})

	resp = jsonFormat(false, "[MS-Users]- Please supply information in JSON format.", make(map[string]interface{}))

	//// returns apikey in JSON converted from GO data
	json.NewEncoder(res).Encode(resp)
}

func unprocessableEntityStatusJSON(res http.ResponseWriter, ok bool, msg string, data map[string]interface{}) {

	// set header to context-type of json format
	res.Header().Set("Content-Type", "application/json")

	// write status code to header
	res.WriteHeader(http.StatusUnprocessableEntity)

	resp := make(map[string]interface{})

	resp = jsonFormat(ok, "[MS-Users]- "+msg, data)

	//// returns apikey in JSON converted from GO data
	json.NewEncoder(res).Encode(resp)
}

func acceptedStatusJSON(res http.ResponseWriter, ok bool, msg string, data map[string]interface{}) {

	// set header to context-type of json format
	res.Header().Set("Content-Type", "application/json")

	// write status code to header
	res.WriteHeader(http.StatusAccepted)

	resp := make(map[string]interface{})

	resp = jsonFormat(ok, "[MS-Users]- "+msg, data)

	//// returns apikey in JSON converted from GO data
	json.NewEncoder(res).Encode(resp)

}

func createdStatusJSON(res http.ResponseWriter, ok bool, msg string, data map[string]interface{}, wg *sync.WaitGroup) {

	defer wg.Done()

	// set header to context-type of json format
	res.Header().Set("Content-Type", "application/json")

	// write status code to header
	res.WriteHeader(http.StatusCreated)

	resp := make(map[string]interface{})

	resp = jsonFormat(ok, "[MS-Users]- "+msg, data)

	//// returns apikey in JSON converted from GO data
	json.NewEncoder(res).Encode(resp)
}

func conflictStatusJSON(res http.ResponseWriter, ok bool, msg string, data map[string]interface{}) {

	// set header to context-type of json format
	res.Header().Set("Content-Type", "application/json")

	// write status code to header
	res.WriteHeader(http.StatusConflict)

	resp := make(map[string]interface{})

	resp = jsonFormat(ok, "[MS-Users]- "+msg, data)

	//// returns apikey in JSON converted from GO data
	json.NewEncoder(res).Encode(resp)
}

func notAcceptableStatusJSON(res http.ResponseWriter, ok bool, msg string, data map[string]interface{}) {

	// set header to context-type of json format
	res.Header().Set("Content-Type", "application/json")

	// write status code to header
	res.WriteHeader(http.StatusNotAcceptable)

	resp := make(map[string]interface{})

	resp = jsonFormat(ok, "[MS-Users]- "+msg, data)

	//// returns apikey in JSON converted from GO data
	json.NewEncoder(res).Encode(resp)
}

func badRequestStatusUser(res http.ResponseWriter, ok bool, msg string, data map[int]interface{}) {

	// set header to context-type of json format
	res.Header().Set("Content-Type", "application/json")

	// write status code to header
	res.WriteHeader(http.StatusBadRequest)

	resp := make(map[string]interface{})

	resp = jsonFormatUser(ok, "[MS-Users]- "+msg, data)

	//// returns apikey in JSON converted from GO data
	json.NewEncoder(res).Encode(resp)

}

func acceptedStatusUser(res http.ResponseWriter, ok bool, msg string, data map[int]interface{}) {

	// set header to context-type of json format
	res.Header().Set("Content-Type", "application/json")

	// write status code to header
	res.WriteHeader(http.StatusAccepted)

	resp := make(map[string]interface{})

	resp = jsonFormatUser(ok, "[MS-Users]- "+msg, data)

	//// returns apikey in JSON converted from GO data
	json.NewEncoder(res).Encode(resp)

}

func createdStatusUser(res http.ResponseWriter, ok bool, msg string, data map[int]database.UserInfo) {

	// set header to context-type of json format
	res.Header().Set("Content-Type", "application/json")

	// write status code to header
	res.WriteHeader(http.StatusCreated)

	resp := make(map[string]interface{})

	resp = jsonFormatAllUsers(ok, "[MS-Users]- "+msg, data)

	//// returns apikey in JSON converted from GO data
	json.NewEncoder(res).Encode(resp)
}

func jsonFormatUser(ok bool, msg string, data map[int]interface{}) map[string]interface{} {
	resp := make(map[string]interface{})
	resp["ok"] = ok
	resp["msg"] = msg
	resp["data"] = data
	return resp
}

func notFoundStatusJSON(res http.ResponseWriter, ok bool, msg string, data map[string]interface{}) {

	res.Header().Set("Content-Type", "application/json")

	// write status code to header
	res.WriteHeader(http.StatusNotFound)

	resp := make(map[string]interface{})

	resp = jsonFormat(ok, "[MS-Users]- "+msg, data)

	json.NewEncoder(res).Encode(resp)
}

func jsonFormat(ok bool, msg string, data map[string]interface{}) map[string]interface{} {
	resp := make(map[string]interface{})
	resp["ok"] = ok
	resp["msg"] = msg
	resp["data"] = data
	return resp
}

func acceptedStatusAllUsers(res http.ResponseWriter, ok bool, msg string, data map[int]database.UserInfo) {

	// set header to context-type of json format
	res.Header().Set("Content-Type", "application/json")

	// write status code to header
	res.WriteHeader(http.StatusAccepted)

	resp := make(map[string]interface{})

	resp = jsonFormatAllUsers(ok, "[MS-Users]- "+msg, data)

	//// returns apikey in JSON converted from GO data
	json.NewEncoder(res).Encode(resp)

}

func jsonFormatAllUsers(ok bool, msg string, data map[int]database.UserInfo) map[string]interface{} {
	resp := make(map[string]interface{})
	resp["ok"] = ok
	resp["msg"] = msg
	resp["data"] = data
	return resp
}

func acceptedStatusAllTransactions(res http.ResponseWriter, ok bool, msg string, data map[int]database.Transactions) {

	// set header to context-type of json format
	res.Header().Set("Content-Type", "application/json")

	// write status code to header
	res.WriteHeader(http.StatusAccepted)

	resp := make(map[string]interface{})

	resp = jsonFormatAllTransactions(ok, "[MS-Users]- "+msg, data)

	//// returns apikey in JSON converted from GO data
	json.NewEncoder(res).Encode(resp)

}

func jsonFormatAllTransactions(ok bool, msg string, data map[int]database.Transactions) map[string]interface{} {
	resp := make(map[string]interface{})
	resp["ok"] = ok
	resp["msg"] = msg
	resp["data"] = data
	return resp
}
