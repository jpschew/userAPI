package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
	"strconv"
	"userAPI/config"
	"userAPI/database"
)

var (
	sqlDriver = "mysql"
	dbPath    = "database"
)

var (
	usersMap map[string]bool   // map to contain user phone
	userAPI  map[string]string // map phone to api
	dsn      string
	dbName   string
)

func init() {
	dbConfig, err := config.LoadDBConfig(dbPath, "db")
	if err != nil {
		log.Fatalln(err.Error())
	}
	dsn = dbConfig.UserName + ":" + dbConfig.Password + "@tcp(" + dbConfig.DBIP + ":" + dbConfig.DBPort + ")"
	// need ?parseTime=true after dbName to map timestamp in mysql to time.Time in golang
	dbName = dbConfig.DBName + "?parseTime=true"
	fmt.Println(dsn, dbName)

	usersMap = make(map[string]bool)
	// retrieve data from database and assign to the userAPI map
	db := database.CreateDBConn(sqlDriver, dsn, dbName)
	defer db.Close()
	usersMap = database.InitAllUsers(db)
	userAPI = database.GetAllKeys(db)
	//fmt.Println(usersMap)
	//fmt.Println(userAPI)
}

func Home(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(res, "Welcome to user profile API!")
}

func AddTransaction(res http.ResponseWriter, req *http.Request) {
	//fmt.Println("get all users")

	//// check for valid access token
	//if !validKey(req) {
	//	res.WriteHeader(http.StatusNotFound)
	//	res.Write([]byte("401 -Invalid key"))
	//	return
	//}

	data := make(map[string]interface{})

	// get the header from request
	// check if body is in json format
	// use Content-Type to check for the resource type
	// for POST and PUT, information is sent via the request body

	if req.Header.Get("Content-Type") == "application/json" {
		fmt.Println("json type")

		// POST is for creating new course
		// if duplicate course or module is created
		// will ask user to use PUT to update the content instead
		if req.Method == "POST" {
			db := database.CreateDBConn(sqlDriver, dsn, dbName)
			defer db.Close()

			// close the request body at the end of the function after reading
			defer req.Body.Close()

			// read in data from request body
			if reqBody, err := io.ReadAll(req.Body); err != nil { // error when reading request body
				//fmt.Println("body error")
				unprocessableEntityStatus(res)
			} else { // no error when reading request body

				//fmt.Println("no body error")

				//user := database.UserInfo{"", "", ""}
				info := struct { // fields need to be caps for first letter to take in json input
					Phone  string `json:"phone"`
					Item   string `json:"item"`
					Points int    `json:"points"`
					Weight int    `json:"weight""`
				}{}

				json.Unmarshal(reqBody, &info)
				//fmt.Println(reqBody, string(reqBody))
				//fmt.Println(info)

				if !validPhoneNum(info.Phone) {
					unprocessableEntityStatusJSON(res, false,
						"Phone must be 8 digits integer.", data)
					return
				}

				if !validTransInfo(info.Phone, info.Item, info.Points, info.Weight) {
					unprocessableEntityStatusJSON(res, false,
						"All fields need to be filled.", data)
					return
				}

				if _, ok := usersMap[info.Phone]; ok {
					data["phone"] = info.Phone
					data["item"] = info.Item
					data["points"] = info.Points
					data["weight"] = info.Weight

					database.AddTransaction(db, info.Phone, info.Item, info.Points, info.Weight)
					createdStatusJSON(res, true, "Transactions added.", data)
				} else {
					notFoundStatusJSON(res, false, "User not found.", data)
				}
			}
		}
	} else {
		fmt.Println("no content type")
		notAcceptableStatusJSON(res, false, "Content-type is not JSON format for POST/PUT method.", data)
	}

}

func RetrieveUserPoints(res http.ResponseWriter, req *http.Request) {

	data := make(map[string]interface{})

	// get the header from request
	// check if body is in json format
	// use Content-Type to check for the resource type
	// for POST and PUT, information is sent via the request body
	//&& req.Header.Get("APIkey") != ""
	if req.Header.Get("Content-Type") == "application/json" {
		//fmt.Println("json type")

		// POST is for creating new course
		// if duplicate course or module is created
		// will ask user to use PUT to update the content instead
		if req.Method == "POST" {
			db := database.CreateDBConn(sqlDriver, dsn, dbName)
			defer db.Close()

			// close the request body at the end of the function after reading
			defer req.Body.Close()

			// read in data from request body
			if reqBody, err := io.ReadAll(req.Body); err != nil { // error when reading request body
				//fmt.Println("body error")
				unprocessableEntityStatus(res)
			} else { // no error when reading request body

				//fmt.Println("no body error")

				//user := database.UserInfo{"", "", ""}
				info := struct {
					Phone string `json:"phone"`
				}{}
				json.Unmarshal(reqBody, &info)
				//fmt.Println(info)

				if !validPhoneNum(info.Phone) {
					unprocessableEntityStatusJSON(res, false,
						"Phone must be 8 digits integer.", data)
					return
				}

				if _, ok := usersMap[info.Phone]; ok {

					userID, points := database.RetrievePoints(db, info.Phone)

					data["userID"] = userID
					data["points"] = points

					acceptedStatusJSON(res, true, "Points retrieved.", data)
				} else {
					notFoundStatusJSON(res, false, "User not found.", data)
				}

				//// write status code to header
				//res.WriteHeader(http.StatusAccepted)
				//
				//res.Header().Set("Content-Type", "application/json")
				//resp := make(map[string]int)
				//resp["points"] = points
				//if jsonResp, err := json.Marshal(resp); err != nil {
				//	log.Panicln(err.Error())
				//} else {
				//	res.Write(jsonResp)
				//}
				//res.Write([]byte("202 - Retrieve points:" + string(points)))
				//fmt.Println(http.StatusAccepted, "retrieve points:", points)

			}
		}
	} else {
		//fmt.Println("no content type")
		notAcceptableStatusJSON(res, false, "Content-type is not JSON format for POST/PUT method.", data)
	}
}

func VoucherStatus(res http.ResponseWriter, req *http.Request) {

	data := make(map[string]interface{})

	// get the header from request
	// check if body is in json format
	// use Content-Type to check for the resource type
	// for POST and PUT, information is sent via the request body
	//&& req.Header.Get("APIkey") != ""
	if req.Header.Get("Content-Type") == "application/json" {
		//fmt.Println("json type")

		// POST is for creating new course
		// if duplicate course or module is created
		// will ask user to use PUT to update the content instead
		if req.Method == "POST" {
			db := database.CreateDBConn(sqlDriver, dsn, dbName)
			defer db.Close()

			// close the request body at the end of the function after reading
			defer req.Body.Close()

			// read in data from request body
			if reqBody, err := io.ReadAll(req.Body); err != nil { // error when reading request body
				//fmt.Println("body error")
				unprocessableEntityStatus(res)
			} else { // no error when reading request body

				//fmt.Println("no body error")

				//user := database.UserInfo{"", "", ""}
				info := struct {
					Phone     string `json:"phone"`
					VoucherID string `json:"vID"`
				}{}
				json.Unmarshal(reqBody, &info)
				//fmt.Println(info)

				if !validPhoneNum(info.Phone) {
					unprocessableEntityStatusJSON(res, false,
						"Phone must be 8 digits integer.", data)
					return
				}

				if !validCheckVoucher(info.Phone, info.VoucherID) {
					unprocessableEntityStatusJSON(res, false,
						"All fields must be filled.", data)
					return
				}

				if _, ok := usersMap[info.Phone]; ok {
					// check add user
					userID, vID, redeem := database.RetrieveVoucherStatus(db, info.Phone, info.VoucherID)

					data["userID"] = userID
					data["voucherID"] = vID
					data["redeemed"] = redeem

					acceptedStatusJSON(res, true, "Retrieved voucher status.", data)
				} else {
					notFoundStatusJSON(res, false, "User not found.", data)
				}

			}
		}
	} else {
		//fmt.Println("no content type")
		notAcceptableStatusJSON(res, false, "Content-type is not JSON format for POST/PUT method.", data)
	}
}

func RedeemVoucher(res http.ResponseWriter, req *http.Request) {

	data := make(map[string]interface{})

	// get the header from request
	// check if body is in json format
	// use Content-Type to check for the resource type
	// for POST and PUT, information is sent via the request body
	//&& req.Header.Get("APIkey") != ""
	if req.Header.Get("Content-Type") == "application/json" {
		//fmt.Println("json type")

		// POST is for creating new course
		// if duplicate course or module is created
		// will ask user to use PUT to update the content instead
		if req.Method == "PUT" {
			db := database.CreateDBConn(sqlDriver, dsn, dbName)
			defer db.Close()

			// close the request body at the end of the function after reading
			defer req.Body.Close()

			// read in data from request body
			if reqBody, err := io.ReadAll(req.Body); err != nil { // error when reading request body
				//fmt.Println("body error")
				unprocessableEntityStatus(res)
			} else { // no error when reading request body

				//fmt.Println("no body error")

				//user := database.UserInfo{"", "", ""}
				info := struct {
					Phone     string `json:"phone"`
					VoucherID string `json:"vID"`
				}{}
				json.Unmarshal(reqBody, &info)
				//fmt.Println(info)

				if !validPhoneNum(info.Phone) {
					unprocessableEntityStatusJSON(res, false,
						"Phone must be 8 digits integer.", data)
					return
				}

				if !validCheckVoucher(info.Phone, info.VoucherID) {
					unprocessableEntityStatusJSON(res, false,
						"All fields must be filled.", data)
					return
				}

				// check add user
				if _, ok := usersMap[info.Phone]; ok {
					userID, redeem, validVoucherID := database.RedeemVoucher(db, info.Phone, info.VoucherID)
					if !validVoucherID {
						notFoundStatusJSON(res, false, "Voucher not found.", data)
						return
					}
					if redeem == 1 {
						notAcceptableStatusJSON(res, false, "Voucher has been redeemed.", data)
						return
					}
					data["userID"] = userID
					data["voucherID"] = info.VoucherID
					data["redeemed"] = true

					acceptedStatusJSON(res, true, "Voucher redeemed.", data)

				} else {
					notFoundStatusJSON(res, false, "User not found.", data)
				}
			}
		}
	} else {
		//fmt.Println("no content type")
		notAcceptableStatusJSON(res, false, "Content-type is not JSON format for POST/PUT method.", data)
	}
}

func AddUserVoucher(res http.ResponseWriter, req *http.Request) {

	data := make(map[string]interface{})

	// get the header from request
	// check if body is in json format
	// use Content-Type to check for the resource type
	// for POST and PUT, information is sent via the request body
	//&& req.Header.Get("APIkey") != ""
	if req.Header.Get("Content-Type") == "application/json" {
		//fmt.Println("json type")

		// POST is for creating new course
		// if duplicate course or module is created
		// will ask user to use PUT to update the content instead
		if req.Method == "POST" {
			db := database.CreateDBConn(sqlDriver, dsn, dbName)
			defer db.Close()

			// close the request body at the end of the function after reading
			defer req.Body.Close()

			// read in data from request body
			if reqBody, err := io.ReadAll(req.Body); err != nil { // error when reading request body
				//fmt.Println("body error")
				unprocessableEntityStatus(res)
			} else { // no error when reading request body

				//fmt.Println("no body error")

				info := struct {
					Phone     string `json:"phone"`
					VoucherID string `json:"vID"`
					Amount    int    `json:"amount"`
					Points    int    `json:"points"`
				}{}

				json.Unmarshal(reqBody, &info)
				//fmt.Println(info)

				if !validPhoneNum(info.Phone) {
					unprocessableEntityStatusJSON(res, false,
						"Phone must be 8 digits integer.", data)
					return
				}

				if !validVoucherInfo(info.Phone, info.VoucherID, info.Points, info.Amount) {
					unprocessableEntityStatusJSON(res, false,
						"All fields need to be filled.", data)
					return
				}

				if _, ok := usersMap[info.Phone]; ok {
					// add voucher
					userID, vID, amount, points := database.AddVoucher(db, info.Phone, info.VoucherID, info.Amount, info.Points)

					data["id"] = userID
					data["voucherID"] = vID
					data["amount"] = amount
					data["pointsExchange"] = points

					acceptedStatusJSON(res, true, "Voucher added.", data)
				} else {
					notFoundStatusJSON(res, false, "User not found.", data)
				}

			}
		}
	} else {
		//fmt.Println("no content type")
		notAcceptableStatusJSON(res, false, "Content-type is not JSON format for POST/PUT method.", data)
	}
}

func GetAllTransactions(res http.ResponseWriter, req *http.Request) {

	trans := make(map[int]database.Transactions)
	uTrans := make(map[int][]database.UTransactions)

	data := make(map[int]interface{})

	// Vars returns the route variables for the current request, if any from the gorilla mux
	// return a map with key string and value string
	// params is a map with key(courseid) as string and value(specified in url) as string
	params := mux.Vars(req)

	// Get method - retrieve and request data from a specified recourse (url)
	// does not require the body
	if req.Method == "GET" {
		db := database.CreateDBConn(sqlDriver, dsn, dbName)
		defer db.Close()

		trans, uTrans = database.GetAllTransactions(db)
		//fmt.Println(trans, uTrans)

		if params["userid"] != "" { // userid exists in url

			if uID, err := strconv.Atoi(params["userid"]); err != nil {
				badRequestStatusUser(res, false, "User id need to be integer.", data)
				return
			} else {
				if _, ok := uTrans[uID]; !ok {
					notFoundStatusJSON(res, false, "User not found.", make(map[string]interface{}))
					return
				} else {
					data[uID] = uTrans[uID]
					acceptedStatusUser(res, true, "Retrieved user transactions.", data)
					//// write status code to header
					//res.WriteHeader(http.StatusAccepted)
					//json.NewEncoder(res).Encode(uTrans[uID])
				}
			}

		} else { // userid does not exist in url

			if len(trans) == 0 {
				acceptedStatusAllTransactions(res, true, "No users in database.", trans)
			} else {
				acceptedStatusAllTransactions(res, true, "Get all users.", trans)
			}
			//// write status code to header
			//res.WriteHeader(http.StatusAccepted)
			//
			////res.Header().Set("Content-Type", "application/json")
			////resp := make(map[string]bool)
			////resp["redeem"] = status
			////if jsonResp, err := json.Marshal(resp); err != nil {
			////	log.Panicln(err.Error())
			////} else {
			////	res.Write(jsonResp)
			////}
			//
			//json.NewEncoder(res).Encode(trans)
		}

	}

}

func GetAllUsers(res http.ResponseWriter, req *http.Request) {

	//var results []map[int]database.UserInfo
	users := make(map[int]database.UserInfo)
	data := make(map[int]interface{})

	params := mux.Vars(req)

	// Get method - retrieve and request data from a specified recourse (url)
	// does not require the body
	if req.Method == "GET" {
		db := database.CreateDBConn(sqlDriver, dsn, dbName)
		defer db.Close()

		users = database.GetAllUsers(db)

		//data = users
		if params["userid"] != "" { // userid exists in url

			if uID, err := strconv.Atoi(params["userid"]); err != nil {
				badRequestStatusUser(res, false, "User id need to be integer.", data)
				return
			} else {
				if _, ok := users[uID]; !ok {
					notFoundStatusJSON(res, false, "User not found.", make(map[string]interface{}))
					return
				} else {

					data[uID] = users[uID]
					acceptedStatusUser(res, true, "Retrieved user data.", data)
				}
			}

		} else { // userid does not exist in url

			if len(users) == 0 {
				acceptedStatusAllUsers(res, true, "No users in database.", users)
			} else {
				acceptedStatusAllUsers(res, true, "Get all users.", users)
			}
		}

	}

}

func AddUser(res http.ResponseWriter, req *http.Request) {

	data := make(map[string]interface{})

	// get the header from request
	// check if body is in json format
	// use Content-Type to check for the resource type
	// for POST and PUT, information is sent via the request body
	//&& req.Header.Get("APIkey") != ""
	if req.Header.Get("Content-Type") == "application/json" {
		//fmt.Println("json type")

		// POST is for creating new course
		// if duplicate course or module is created
		// will ask user to use PUT to update the content instead
		if req.Method == "POST" {
			db := database.CreateDBConn(sqlDriver, dsn, dbName)
			defer db.Close()

			// close the request body at the end of the function after reading
			defer req.Body.Close()

			// read in data from request body
			if reqBody, err := io.ReadAll(req.Body); err != nil { // error when reading request body
				//fmt.Println("body error")
				unprocessableEntityStatus(res)
			} else { // no error when reading request body

				//fmt.Println("no body error")

				user := struct {
					Name     string `json:"name"`
					Phone    string `json:"phone"`
					Password string `json:"password""`
					//APIKey string `json:"key"`
				}{}

				//userMap := make(map[string]string)

				json.Unmarshal(reqBody, &user)
				userPW, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
				//fmt.Println(user, userPW)

				if !validPhoneNum(user.Phone) {
					unprocessableEntityStatusJSON(res, false,
						"Phone must be 8 digits integer.", data)
					return
				}

				if !validUserInfo(user.Name, user.Phone, user.Password) {
					unprocessableEntityStatusJSON(res, false,
						"All fields need to be filled.", data)
					return
				}

				if _, ok := usersMap[user.Phone]; !ok { // existing user
					usersMap[user.Phone] = true
					data["name"] = user.Name
					data["phone"] = user.Phone

					database.AddUser(db, user.Name, user.Phone, string(userPW))
					createdStatusJSON(res, true, "User added.", data)

				} else {
					conflictStatusJSON(res, false, "Duplicate user.", data)
				}

				//// check add user
				//userExist := database.CheckAddUser(db, user.Name, user.Phone)
				//if userExist {
				//	acceptedStatus(res)
				//} else {
				//	createdStatus(res)
				//}
			}
		}
	} else {

		notAcceptableStatusJSON(res, false, "Content-type is not JSON format for POST/PUT method.", data)
	}
}

func validUserInfo(name string, phone string, password string) bool {
	if name == "" || phone == "" || password == "" {
		return false
	}
	return true
}

func validTransInfo(phone string, item string, points int, weight int) bool {
	if phone == "" || item == "" || points == 0 || weight == 0 {
		return false
	}
	return true
}

func validVoucherInfo(phone string, vID string, points int, amount int) bool {
	if phone == "" || vID == "" || points == 0 || amount == 0 {
		return false
	}
	return true
}

func validCheckVoucher(phone string, vID string) bool {
	if phone == "" || vID == "" {
		return false
	}
	return true
}

func validPhoneNum(phone string) bool {
	if _, err := strconv.Atoi(phone); err != nil {
		return false
	} else {
		if len(phone) != 8 {
			return false
		}
	}
	return true
}

//// validKey checks for a valid key to secure the REST API
//// so that only authenticated user can use the REST API
//func validKey(r *http.Request) bool {
//	v := r.URL.Query()
//	//fmt.Println(v)
//	//fmt.Println(userAPIKey)
//	// check if user exists
//	if user, ok := v["user"]; ok && user[0] != "" {
//		//fmt.Println(user[0])
//		// check if key exists
//		if key, ok := v["key"]; ok && key[0] != "" {
//			//fmt.Println(user[0], key[0])
//			// check if key tagger to user is correct
//			if userAPIKey[user[0]] == key[0] {
//				return true
//			} else {
//				return false
//			}
//		}
//		return false
//	}
//	return false
//}
