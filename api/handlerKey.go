package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"userAPI/apikey"
	"userAPI/database"
)

// GenKey is the handler function for generating an API key using SHA256 hashing algorithm.
func GenKey(res http.ResponseWriter, req *http.Request) {
	//fmt.Fprintf(res, "Welcome to the BEST UNIVERSITY COURSES API!")

	apiKey := apikey.GenerateAPIKey()

	//// display the message to user
	//fmt.Fprintf(res, "generated key: %s\n", apiKey)

	//// set header to context-type of json format
	//res.Header().Set("Content-Type", "application/json")
	//
	//// write status code to header
	//res.WriteHeader(http.StatusAccepted)

	apiMap := make(map[string]interface{})
	apiMap["apiKey"] = apiKey

	createdStatusJSON(res, true, "API key generated successfully.", apiMap)

	//resp := make(map[string]interface{})
	//resp["ok"] = true
	//resp["msg"] = "[MS-Users]- API key generated successfully."
	//resp["data"] = apiMap
	//
	////// returns apikey in JSON converted from GO data
	//json.NewEncoder(res).Encode(resp)

	//return apiKey

}

// AddUpdateKey adds/updates the API key to the overall map and database.
func AddUpdateKey(res http.ResponseWriter, req *http.Request) {

	//fmt.Println(req.Header)
	if req.Header.Get("Content-Type") == "application/json" {
		//if req.Method == "POST" {
		//
		//	// connect to database to add key to key table
		//	db := database.CreateDBConn(sqlDriver, dsn, dbName)
		//	defer db.Close()
		//
		//	defer req.Body.Close()
		//
		//	if reqBody, err := ioutil.ReadAll(req.Body); err != nil { // error when reading request body
		//
		//		userAPIProcessableEntityStatus(res)
		//
		//	} else { // no error when reading request body
		//
		//		newKey := database.Key{"", ""}
		//		json.Unmarshal(reqBody, &newKey)
		//		//fmt.Println(string(reqBody), newKey)
		//
		//		if _, ok := userAPIKey[newKey.UserName]; ok { // user exists
		//			userConflictStatus(res)
		//		} else {
		//
		//			userAPIKey[newKey.UserName] = newKey.APIKey
		//			database.AddKey(db, newKey)
		//			createdStatusKey(res, newKey.UserName)
		//		}
		//	}
		//}

		if req.Method == "PUT" {

			// connect to database to add key to key table
			db := database.CreateDBConn(sqlDriver, dsn, dbName)
			defer db.Close()

			defer req.Body.Close()

			if reqBody, err := ioutil.ReadAll(req.Body); err != nil { // error when reading request body

				unprocessableEntityStatus(res)

			} else { // no error when reading request body

				newKey := struct {
					Username string `json:"username"`
					APIKey   string `json:"apiKey"`
				}{}

				json.Unmarshal(reqBody, &newKey)
				//fmt.Println(string(reqBody), newKey)

				//if !validPhoneNum(newKey.Phone) {
				//	unprocessableEntityStatusJSON(res, false,
				//		"Phone must be 8 digits integer.", data)
				//	return
				//}

				if _, ok := userAPI[newKey.Username]; ok {
					database.UpdateKey(db, newKey.Username, newKey.APIKey)

					//database.AddUpdateKey(db, newKey.Phone, newKey.APIKey)
					//userAPI[newKey.Phone] = newKey.APIKey
					acceptedStatusJSON(res, true, "API key updated.", make(map[string]interface{}))
				} else {
					database.AddKey(db, newKey.Username, newKey.APIKey)
					acceptedStatusJSON(res, true, "API key added.", make(map[string]interface{}))

					//notFoundStatusJSON(res, false, "User not found.", make(map[string]interface{}))
				}
				userAPI[newKey.Username] = newKey.APIKey
				
			}
		}
	}
}

//
//// DeleteKey deletes the API key from the overall map and database.
//func DeleteKey(res http.ResponseWriter, req *http.Request) {
//
//	if req.Header.Get("Content-Type") == "application/json" {
//		if req.Method == "POST" {
//			// connect to database to delete course/module from course/module table
//			db := database.CreateDBConn(sqlDriver, dsn, dbName)
//			defer db.Close()
//
//			defer req.Body.Close()
//
//			if reqBody, err := ioutil.ReadAll(req.Body); err != nil { // error when reading body
//				userAPIProcessableEntityStatus(res)
//			} else {
//				newKey := database.Key{"", ""}
//
//				json.Unmarshal(reqBody, &newKey)
//
//				if _, ok := userAPIKey[newKey.UserName]; !ok {
//					notFoundStatusUser(res)
//				} else {
//					// this will delete the key from a map
//					// first arg is the map
//					// second arg is the key
//					delete(userAPIKey, newKey.UserName)
//
//					database.DeleteKey(db, newKey)
//					acceptedStatusKey(res, newKey.UserName, false)
//				}
//			}
//
//		}
//	}
//}
