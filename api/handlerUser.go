package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
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

				info := struct { // fields need to be caps for first letter to take in json input
					Phone  string `json:"phone"`
					Item   string `json:"item"`
					Points int    `json:"points"`
					Weight int    `json:"weight"`
				}{}

				json.Unmarshal(reqBody, &info)
				//fmt.Println(reqBody, string(reqBody))
				//fmt.Println(info)

				//  check if it is a valid phone number
				if !validPhoneNum(info.Phone) {
					unprocessableEntityStatusJSON(res, false,
						"Phone must be 8 digits integer.", data)
					return
				}

				// check if all information is there
				if !validTransInfo(info.Phone, info.Item, info.Points, info.Weight) {
					unprocessableEntityStatusJSON(res, false,
						"All fields need to be filled.", data)
					return
				}

				// if phone in userMap
				// add transaction to database
				if _, ok := usersMap[info.Phone]; ok {
					data["phone"] = info.Phone
					data["item"] = info.Item
					data["points"] = info.Points
					data["weight"] = info.Weight

					// add transaction to database
					database.AddTransaction(db, info.Phone, info.Item, info.Points, info.Weight)
					createdStatusJSON(res, true, "Transactions added.", data)
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

func RetrieveUserPoints(res http.ResponseWriter, req *http.Request) {

	data := make(map[string]interface{})

	// get the header from request
	// check if body is in json format
	// use Content-Type to check for the resource type
	// for POST and PUT, information is sent via the request body
	//&& req.Header.Get("APIkey") != ""
	if req.Header.Get("Content-Type") == "application/json" {

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
					Phone string `json:"phone"`
				}{}
				json.Unmarshal(reqBody, &info)
				//fmt.Println(info)

				// check if phone number is valid
				if !validPhoneNum(info.Phone) {
					unprocessableEntityStatusJSON(res, false,
						"Phone must be 8 digits integer.", data)
					return
				}

				// if phone is in userMap
				// retrieve points from database
				if _, ok := usersMap[info.Phone]; ok {

					// retrieve points from database
					userID, points := database.RetrievePoints(db, info.Phone)

					data["userID"] = userID
					data["points"] = points

					acceptedStatusJSON(res, true, "Points retrieved.", data)
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

func VoucherStatus(res http.ResponseWriter, req *http.Request) {

	data := make(map[string]interface{})

	// get the header from request
	// check if body is in json format
	// use Content-Type to check for the resource type
	// for POST and PUT, information is sent via the request body
	//&& req.Header.Get("APIkey") != ""
	if req.Header.Get("Content-Type") == "application/json" {

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
				}{}
				json.Unmarshal(reqBody, &info)
				//fmt.Println(info)

				// check if phone number is valid
				if !validPhoneNum(info.Phone) {
					unprocessableEntityStatusJSON(res, false,
						"Phone must be 8 digits integer.", data)
					return
				}

				// check if all information is present
				if !validCheckVoucher(info.Phone, info.VoucherID) {
					unprocessableEntityStatusJSON(res, false,
						"All fields must be filled.", data)
					return
				}

				// if phone in userMap
				// retrieve voucher status from database
				if _, ok := usersMap[info.Phone]; ok {

					// retrieve voucher status from database
					userID, vID, redeem, validVoucherID := database.RetrieveVoucherStatus(db, info.Phone, info.VoucherID)

					// check if voucherID is valid
					if !validVoucherID {
						notFoundStatusJSON(res, false, "Voucher not found.", data)
						return
					}

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
	if req.Header.Get("Content-Type") == "application/json" {

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

				info := struct {
					Phone     string `json:"phone"`
					VoucherID string `json:"vID"`
				}{}
				json.Unmarshal(reqBody, &info)
				//fmt.Println(info)

				// check if phone number is valid
				if !validPhoneNum(info.Phone) {
					unprocessableEntityStatusJSON(res, false,
						"Phone must be 8 digits integer.", data)
					return
				}

				// check if all information is present
				if !validCheckVoucher(info.Phone, info.VoucherID) {
					unprocessableEntityStatusJSON(res, false,
						"All fields must be filled.", data)
					return
				}

				// if phone number in userMap
				// redeem voucher from database
				if _, ok := usersMap[info.Phone]; ok {

					// redeem voucher from database
					userID, redeem, validVoucherID := database.RedeemVoucher(db, info.Phone, info.VoucherID)

					// check if voucherID is valid
					if !validVoucherID {
						notFoundStatusJSON(res, false, "Voucher not found.", data)
						return
					}
					// check if voucher has been redeemed
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

				// check if all voucher information are present
				if !validVoucherInfo(info.Phone, info.VoucherID, info.Points, info.Amount) {
					unprocessableEntityStatusJSON(res, false,
						"All fields need to be filled.", data)
					return
				}

				// if phone number in userMap
				// add voucher to user in database
				if _, ok := usersMap[info.Phone]; ok {
					// add voucher to user in database
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
	uTrans := make(map[int][]database.Transactions)

	data := make(map[int]interface{})

	// get query strings
	// should have page_index and records_per_page
	// values will be a map
	values := req.URL.Query()

	// check if query strings are present
	if _, ok := values["page_index"]; !ok {
		badRequestStatusUser(res, false, "Both page_index and records_per_page need to be provided.", data)
		return
	}
	if _, ok := values["records_per_page"]; !ok {
		badRequestStatusUser(res, false, "Both page_index and records_per_page need to be provided.", data)
		return
	}

	page, records, valid := getQueryStrings(values)

	// check if query string values are provided correctly
	if !valid {
		badRequestStatusUser(res, false, "Both page_index and records_per_page need to be integer.", data)
		return
	}
	if !positiveInt(page, records) {
		badRequestStatusUser(res, false, "Page_index and records_per_page need to be equal or larger "+
			"than 0 and 1 respectively to be able to retrieve transactions.", data)
		return
	}

	// Vars returns the route variables for the current request, if any from the gorilla mux
	// return a map with key string and value string
	// params is a map with key(userid, itemid) as string and value(specified in url) as string
	params := mux.Vars(req)

	// Get method - retrieve and request data from a specified recourse (url)
	// does not require the body
	if req.Method == "GET" {
		db := database.CreateDBConn(sqlDriver, dsn, dbName)
		defer db.Close()

		//trans = database.GetAllTransactions(db, page, records)
		//fmt.Println(trans, uTrans)

		if params["userid"] != "" { // userid exists in url

			// check if value is integer
			if uID, err := strconv.Atoi(params["userid"]); err != nil {
				badRequestStatusUser(res, false, "User id need to be integer.", data)
				return
			} else { // if value is integer

				var msg string

				if params["itemid"] != "" { // itemid exists in url

					// convert to Title case
					item := cases.Title(language.Und, cases.NoLower).String(strings.ToLower(params["itemid"]))

					// get user transactions by item from database
					uTrans = database.GetUserTransactionsByItem(db, page, records, uID, item)

					msg = fmt.Sprintf("Get %d %s transcations for userID %d.",
						len(uTrans[uID]), item, uID)

				} else { // itemid does not exists in url

					// get user transactions from database
					uTrans = database.GetUserTransactions(db, page, records, uID)

					msg = fmt.Sprintf("Get %d transcations for userID %d.",
						len(uTrans[uID]), uID)

				}

				// check if userid in uTrans
				if _, ok := uTrans[uID]; !ok { // if uID not in uTrans means no transactions retrieved from database

					if page == 0 { // if no offset and no transactions
						msg = fmt.Sprintf("No transactions in database userID %d.",
							uID)
					} else { // 1 or more pages
						msg = fmt.Sprintf("No more transactions available in database userID %d.",
							uID)
					}
					notFoundStatusJSON(res, false, msg, make(map[string]interface{}))

					return
				} else { // if uID in uTrans means transactions retrieved from database

					data[uID] = uTrans[uID]

					acceptedStatusUser(res, true, msg, data)
				}

			}

		} else { // userid does not exist in url

			// get all transactions from database
			trans = database.GetAllTransactions(db, page, records)

			// if no transactions retrieved from database
			if len(trans) == 0 {
				if page == 0 { // if no offset and no transactions
					notFoundStatusJSON(res, false, "No transactions in database.",
						make(map[string]interface{}))
				} else { // 1 or more pages
					notFoundStatusJSON(res, false, "No more transactions available in database.",
						make(map[string]interface{}))
				}

			} else { // if transactions retrieved from database
				msg := fmt.Sprintf("Get transcation id %d to %d .",
					(page*records)+1, (page*records)+len(trans))
				if (page*records)+1 == (page*records)+len(trans) {
					msg = fmt.Sprintf("Get transaction id %d.", (page*records)+1)
				}
				acceptedStatusAllTransactions(res, true, msg, trans)
			}
		}
	}

}

func GetAllUsers(res http.ResponseWriter, req *http.Request) {

	//var results []map[int]database.UserInfo
	users := make(map[int]database.UserInfo)
	data := make(map[int]interface{})

	// get query strings
	// should have page_index and records_per_page
	// values will be a map
	values := req.URL.Query()

	// check if query strings are present
	if _, ok := values["page_index"]; !ok {
		badRequestStatusUser(res, false, "Both page_index and records_per_page need to be provided.", data)
		return
	}
	if _, ok := values["records_per_page"]; !ok {
		badRequestStatusUser(res, false, "Both page_index and records_per_page need to be provided.", data)
		return
	}

	page, records, valid := getQueryStrings(values)

	// check if query string values are provided correctly
	if !valid {
		badRequestStatusUser(res, false, "Both page_index and records_per_page need to be integer.", data)
		return
	}
	if !positiveInt(page, records) {
		badRequestStatusUser(res, false, "Page_index and records_per_page need to be equal or larger "+
			"than 0 and 1 respectively to be able to retrieve users.", data)
		return
	}

	// Get method - retrieve and request data from a specified recourse (url)
	// does not require the body
	if req.Method == "GET" {
		db := database.CreateDBConn(sqlDriver, dsn, dbName)
		defer db.Close()

		users = database.GetAllUsers(db, page, records)

		// if no users retrieved from database
		if len(users) == 0 {

			if page == 0 { // if no offset and no transactions
				notFoundStatusJSON(res, false, "No users in database.",
					make(map[string]interface{}))
			} else { // 1 or more pages
				notFoundStatusJSON(res, false, "No more users available in database.",
					make(map[string]interface{}))
			}
		} else { // users retrieved from database

			msg := fmt.Sprintf("Get user id %d to %d.", (page*records)+1, (page*records)+len(users))
			if (page*records)+1 == (page*records)+len(users) {
				msg = fmt.Sprintf("Get user id %d.", (page*records)+1)
			}

			acceptedStatusAllUsers(res, true, msg, users)
		}
	}
}

func AddUser(res http.ResponseWriter, req *http.Request) {

	data := make(map[string]interface{})

	// get the header from request
	// check if body is in json format
	// use Content-Type to check for the resource type
	// for POST and PUT, information is sent via the request body
	if req.Header.Get("Content-Type") == "application/json" {

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
					Password string `json:"password"`
				}{}

				json.Unmarshal(reqBody, &user)
				userPW, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
				//fmt.Println(user, userPW)

				// check if phone number is valid
				if !validPhoneNum(user.Phone) {
					unprocessableEntityStatusJSON(res, false,
						"Phone must be 8 digits integer.", data)
					return
				}

				// check if all user information are present
				if !validUserInfo(user.Name, user.Phone, user.Password) {
					unprocessableEntityStatusJSON(res, false,
						"All fields need to be filled.", data)
					return
				}

				// if phone number exists in userMap
				if _, ok := usersMap[user.Phone]; !ok { // new user, add user to database
					usersMap[user.Phone] = true
					data["name"] = user.Name
					data["phone"] = user.Phone

					// add user to database
					database.AddUser(db, user.Name, user.Phone, string(userPW))
					createdStatusJSON(res, true, "User added.", data)

				} else { // existing user, duplicate user
					conflictStatusJSON(res, false, "Duplicate user.", data)
				}
			}
		}
	} else { // content type not json

		notAcceptableStatusJSON(res, false, "Content-type is not JSON format for POST/PUT method.", data)
	}
}

func getQueryStrings(queryString map[string][]string) (int, int, bool) {

	pageIndex := queryString["page_index"][0]
	recordsPerPage := queryString["records_per_page"][0]

	page, records, valid := validParams(pageIndex, recordsPerPage)

	return page, records, valid
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

func validParams(pageIndex string, recordsPerPage string) (int, int, bool) {

	var page, records int
	var err error

	if page, err = strconv.Atoi(pageIndex); err != nil {
		return 0, 0, false
	}
	if records, err = strconv.Atoi(recordsPerPage); err != nil {
		return 0, 0, false
	}

	return page, records, true
}

func positiveInt(pageIndex int, recordsPerPage int) bool {

	if pageIndex < 0 || recordsPerPage < 1 {
		return false
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
