package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"userAPI/database"
)

//func notFoundStatus(res http.ResponseWriter, isCourse bool) {
//	// write status code to header
//	res.WriteHeader(http.StatusNotFound)
//
//	if isCourse {
//		// write to the response
//		// Write() takes in []byte as input
//		// response code 404 means not found
//		res.Write([]byte("404 - No course found"))
//		fmt.Println(http.StatusNotFound, "- No course found.")
//	} else {
//		// write to the response
//		// Write() takes in []byte as input
//		// response code 404 means not found
//		res.Write([]byte("404 - No module found"))
//		fmt.Println(http.StatusNotFound, "- No module found.")
//	}
//
//}
//
func acceptedStatus(res http.ResponseWriter) {

	// write status code to header
	res.WriteHeader(http.StatusAccepted)

	res.Write([]byte("202 - User Found"))
	fmt.Println(http.StatusAccepted, "user found")

}

func unprocessableEntityStatus(res http.ResponseWriter) {

	// write status code to header
	res.WriteHeader(http.StatusUnprocessableEntity)

	res.Write([]byte("422 - Please supply course information with course code and course name in JSON format."))
	fmt.Println(http.StatusUnprocessableEntity, "Please supply course information with course code and course name in JSON format.")
	//if isCourse {
	// write to the response
	// Write() takes in []byte as input
	// response code 422 means unprocessable entity
	//res.Write([]byte("422 - Please supply course information with course code and course name in JSON format."))
	//fmt.Println(http.StatusUnprocessableEntity, "Please supply course information with course code and course name in JSON format.")
	//} else {
	//	// write to the response
	//	// Write() takes in []byte as input
	//	// response code 422 means unprocessable entity
	//	res.Write([]byte("422 - Please supply course information with module code and module name in JSON format."))
	//	fmt.Println(http.StatusUnprocessableEntity, "Please supply module information with module code and module name in JSON format.")
	//}

}

func createdStatus(res http.ResponseWriter) {

	// write status code to header
	res.WriteHeader(http.StatusCreated)

	res.Write([]byte("201 - User added"))
	fmt.Println(http.StatusCreated, "User added.")

	//if isCourse {
	// write to the response
	// Write() takes in []byte as input
	// response code 201 means created successfully
	//res.Write([]byte("201 - Course added:" + urlParams["courseid"]))
	//fmt.Println(http.StatusCreated, urlParams["courseid"], "added.")
	//} else {
	//	// write to the response
	//	// Write() takes in []byte as input
	//	// response code 201 means created successfully
	//	res.Write([]byte("201 - Module added:" + urlParams["moduleid"]))
	//	fmt.Println(http.StatusCreated, urlParams["moduleid"], "added.")
	//}
	//
}

//
//func conflictStatus(res http.ResponseWriter, isCourse bool) {
//	// write status code to header
//	res.WriteHeader(http.StatusConflict)
//
//	if isCourse {
//		// write to the response
//		// Write() takes in []byte as input
//		// response code 409 means conflict
//		res.Write([]byte("409 - Duplicate course ID. Please use PUT method if you want to update the content."))
//		fmt.Println(http.StatusConflict, "Duplicate course ID. Please use PUT method if you want to update the content.")
//	} else {
//		// write to the response
//		// Write() takes in []byte as input
//		// response code 409 means conflict
//		res.Write([]byte("409 - Duplicate module ID. Please use PUT method if you want to update the content."))
//		fmt.Println(http.StatusConflict, "Duplicate module ID. Please use PUT method if you want to update the content.")
//	}
//}

//func userConflictStatus(res http.ResponseWriter) {
//	// write status code to header
//	res.WriteHeader(http.StatusConflict)
//
//	// write to the response
//	// Write() takes in []byte as input
//	// response code 409 means conflict
//	res.Write([]byte("409 - Duplicate user. You already have one api key."))
//	fmt.Println(http.StatusConflict, "Duplicate user. You already have one api key.")
//}
//
//func userAPIProcessableEntityStatus(res http.ResponseWriter) {
//
//	// write status code to header
//	res.WriteHeader(http.StatusUnprocessableEntity)
//
//	// write to the response
//	// Write() takes in []byte as input
//	// response code 422 means unprocessable entity
//	res.Write([]byte("422 - Please supply user and key information in JSON format."))
//	fmt.Println(http.StatusUnprocessableEntity, "Please supply user and key information in JSON format.")
//
//}
//
//func createdStatusKey(res http.ResponseWriter, user string) {
//
//	// write status code to header
//	res.WriteHeader(http.StatusCreated)
//
//	res.Write([]byte("201 - Key added to " + user))
//	fmt.Println(http.StatusCreated, "Key added to "+user)
//
//}

//func notFoundStatusUser(res http.ResponseWriter) {
//	// write status code to header
//	res.WriteHeader(http.StatusNotFound)
//
//	// write to the response
//	// Write() takes in []byte as input
//	// response code 404 means not found
//	res.Write([]byte("404 - User not found"))
//	fmt.Println(http.StatusNotFound, "- User not found.")
//
//}
//
//func acceptedStatusKey(res http.ResponseWriter, user string, isUpdate bool) {
//	// write status code to header
//	res.WriteHeader(http.StatusAccepted)
//
//	if isUpdate {
//		// write to the response
//		// Write() takes in []byte as input
//		// response code 202 means accepted for processing
//		res.Write([]byte("202 - Key updated to " + user))
//		fmt.Println(http.StatusAccepted, "Key updated to ", user)
//	} else {
//		// write to the response
//		// Write() takes in []byte as input
//		// response code 202 means accepted for processing
//		res.Write([]byte("202 -  Key deleted for " + user))
//		fmt.Println(http.StatusAccepted, "Key deleted for ", user)
//	}
//}

func unprocessableEntityStatusJSON(res http.ResponseWriter, ok bool, msg string, data map[string]interface{}) {

	// set header to context-type of json format
	res.Header().Set("Content-Type", "application/json")

	// write status code to header
	res.WriteHeader(http.StatusUnprocessableEntity)

	resp := make(map[string]interface{})

	resp = jsonFormat(ok, "[MS-Users]- "+msg, data)
	//resp["ok"] = true
	//resp["msg"] = "[MS-Users]- API key generated successfully."
	//resp["data"] = apiMap

	//// returns apikey in JSON converted from GO data
	json.NewEncoder(res).Encode(resp)

	//// write status code to header
	//res.WriteHeader(http.StatusUnprocessableEntity)
	//
	//res.Write([]byte("422 - Please supply course information with course code and course name in JSON format."))
	//fmt.Println(http.StatusUnprocessableEntity, "Please supply course information with course code and course name in JSON format.")

}

func acceptedStatusJSON(res http.ResponseWriter, ok bool, msg string, data map[string]interface{}) {

	// set header to context-type of json format
	res.Header().Set("Content-Type", "application/json")

	// write status code to header
	res.WriteHeader(http.StatusAccepted)

	resp := make(map[string]interface{})

	resp = jsonFormat(ok, "[MS-Users]- "+msg, data)
	//resp["ok"] = true
	//resp["msg"] = "[MS-Users]- API key generated successfully."
	//resp["data"] = apiMap

	//// returns apikey in JSON converted from GO data
	json.NewEncoder(res).Encode(resp)

}

func createdStatusJSON(res http.ResponseWriter, ok bool, msg string, data map[string]interface{}) {

	// set header to context-type of json format
	res.Header().Set("Content-Type", "application/json")

	// write status code to header
	res.WriteHeader(http.StatusCreated)

	resp := make(map[string]interface{})

	resp = jsonFormat(ok, "[MS-Users]- "+msg, data)
	//resp["ok"] = true
	//resp["msg"] = "[MS-Users]- API key generated successfully."
	//resp["data"] = apiMap

	//// returns apikey in JSON converted from GO data
	json.NewEncoder(res).Encode(resp)

	//// write status code to header
	//res.WriteHeader(http.StatusCreated)
	//
	//res.Write([]byte("201 - User added"))
	//fmt.Println(http.StatusCreated, "User added.")

	//if isCourse {
	// write to the response
	// Write() takes in []byte as input
	// response code 201 means created successfully
	//res.Write([]byte("201 - Course added:" + urlParams["courseid"]))
	//fmt.Println(http.StatusCreated, urlParams["courseid"], "added.")
	//} else {
	//	// write to the response
	//	// Write() takes in []byte as input
	//	// response code 201 means created successfully
	//	res.Write([]byte("201 - Module added:" + urlParams["moduleid"]))
	//	fmt.Println(http.StatusCreated, urlParams["moduleid"], "added.")
	//}
	//
}

func conflictStatusJSON(res http.ResponseWriter, ok bool, msg string, data map[string]interface{}) {

	// set header to context-type of json format
	res.Header().Set("Content-Type", "application/json")

	// write status code to header
	res.WriteHeader(http.StatusConflict)

	resp := make(map[string]interface{})

	resp = jsonFormat(ok, "[MS-Users]- "+msg, data)
	//resp["ok"] = true
	//resp["msg"] = "[MS-Users]- API key generated successfully."
	//resp["data"] = apiMap

	//// returns apikey in JSON converted from GO data
	json.NewEncoder(res).Encode(resp)

	//// write status code to header
	//res.WriteHeader(http.StatusConflict)
	//
	//res.Write([]byte("409 - Duplicate course ID. Please use PUT method if you want to update the content."))
	//fmt.Println(http.StatusConflict, "Duplicate course ID. Please use PUT method if you want to update the content.")

}

func notAcceptableStatusJSON(res http.ResponseWriter, ok bool, msg string, data map[string]interface{}) {

	// set header to context-type of json format
	res.Header().Set("Content-Type", "application/json")

	// write status code to header
	res.WriteHeader(http.StatusNotAcceptable)

	resp := make(map[string]interface{})

	resp = jsonFormat(ok, "[MS-Users]- "+msg, data)
	//resp["ok"] = true
	//resp["msg"] = "[MS-Users]- API key generated successfully."
	//resp["data"] = apiMap

	//// returns apikey in JSON converted from GO data
	json.NewEncoder(res).Encode(resp)

	//// write status code to header
	//res.WriteHeader(http.StatusConflict)
	//
	//res.Write([]byte("409 - Duplicate course ID. Please use PUT method if you want to update the content."))
	//fmt.Println(http.StatusConflict, "Duplicate course ID. Please use PUT method if you want to update the content.")

}

func badRequestStatusUser(res http.ResponseWriter, ok bool, msg string, data map[int]interface{}) {

	// set header to context-type of json format
	res.Header().Set("Content-Type", "application/json")

	// write status code to header
	res.WriteHeader(http.StatusBadRequest)

	resp := make(map[string]interface{})

	resp = jsonFormatUser(ok, "[MS-Users]- "+msg, data)
	//resp["ok"] = true
	//resp["msg"] = "[MS-Users]- API key generated successfully."
	//resp["data"] = apiMap

	//// returns apikey in JSON converted from GO data
	json.NewEncoder(res).Encode(resp)

	//// write status code to header
	//res.WriteHeader(http.StatusConflict)
	//
	//res.Write([]byte("409 - Duplicate course ID. Please use PUT method if you want to update the content."))
	//fmt.Println(http.StatusConflict, "Duplicate course ID. Please use PUT method if you want to update the content.")

}

func acceptedStatusUser(res http.ResponseWriter, ok bool, msg string, data map[int]interface{}) {

	// set header to context-type of json format
	res.Header().Set("Content-Type", "application/json")

	// write status code to header
	res.WriteHeader(http.StatusAccepted)

	resp := make(map[string]interface{})

	resp = jsonFormatUser(ok, "[MS-Users]- "+msg, data)
	//resp["ok"] = true
	//resp["msg"] = "[MS-Users]- API key generated successfully."
	//resp["data"] = apiMap

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
	//resp["ok"] = true
	//resp["msg"] = "[MS-Users]- API key generated successfully."
	//resp["data"] = apiMap

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
	//resp["ok"] = true
	//resp["msg"] = "[MS-Users]- API key generated successfully."
	//resp["data"] = apiMap

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
