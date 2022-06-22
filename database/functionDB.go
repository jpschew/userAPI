package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	// implementation of GO's database/sql/driver interface
	// only need to import and use the GO's database/sql API
	_ "github.com/go-sql-driver/mysql"
)

type UserInfo struct {
	UserID int
	Phone  string
	Name   string
	//ApiKey    string
	Points    int
	LastLogin time.Time
}

type Transactions struct {
	TransID   int
	UserID    int
	Name      string
	Weight    int
	Item      string
	TransDate time.Time
}

//type UTransactions struct {
//	UserID  int
//	TransID int
//	//Name      string
//	Weight    int
//	Item      string
//	TransDate time.Time
//}

// CreateDBConn creates a connection to mysql database given the driver name, dsn and db name.
func CreateDBConn(driver string, dsn string, dbName string) *sql.DB {

	// Use mysql as driver Name and a valid DSN as data SourceName:
	source := dsn + "/" + dbName
	db, err := sql.Open(driver, source)

	// handle error
	if err != nil {
		//panic(err.Error())
		log.Panicln(err.Error())
	}

	return db
}

func InitAllUsers(db *sql.DB) map[string]bool {

	users := make(map[string]bool)

	//user := struct {
	//	Phone string
	//	API   string
	//}{}
	var phone string

	query := fmt.Sprintf(`
								SELECT phone
								FROM Users
								`)

	if results, err := db.Query(query); err != nil {
		log.Panicln(err.Error())
	} else {

		for results.Next() {
			err = results.Scan(&phone)

			if err != nil {
				log.Panicln(err.Error())
			}

			users[phone] = true
		}
	}

	return users
}

func GetAllKeys(db *sql.DB) map[string]string {

	userAPI := make(map[string]string)

	user := struct {
		Username string
		API      string
	}{}

	query := fmt.Sprintf(`
								SELECT username, api_key
								FROM Api
								`)

	if results, err := db.Query(query); err != nil {
		log.Panicln(err.Error())
	} else {

		for results.Next() {
			err = results.Scan(&user.Username, &user.API)

			if err != nil {
				log.Panicln(err.Error())
			}

			userAPI[user.Username] = user.API
		}
	}

	return userAPI
}

func GetAllTransactions(db *sql.DB, pageIndex int, recordsPerPage int) map[int]Transactions {

	var trans Transactions
	//var uTrans Transactions
	//var tID, uID int

	transactions := make(map[int]Transactions)
	//userTrans := make(map[int][]Transactions)

	query := fmt.Sprintf(`
								SELECT t.id, u.id, u.name, t.trans_date, t.weight, i.name
								FROM my_db.Users u
								JOIN my_db.Transactions t
								ON u.id = t.user_id
								JOIN my_db.Items i
								ON i.id = t.item_id
								ORDER BY t.id
								LIMIT %d OFFSET %d
								`, recordsPerPage, pageIndex*recordsPerPage)

	if results, err := db.Query(query); err != nil {
		log.Panicln(err.Error())
	} else {
		for results.Next() {
			// Scan() copy each row of data from db and assign to the address specified
			err = results.Scan(&trans.TransID, &trans.UserID, &trans.Name, &trans.TransDate, &trans.Weight, &trans.Item)
			//fmt.Println(user.ApiKey, "db")
			//fmt.Println(input.ApiKey, "input")
			if err != nil {
				//fmt.Println(course.CourseCode, course.CourseName, course.Description, module.ModuleCode, module.ModuleName, module.Description)
				log.Panicln(err.Error())

			}

			transactions[trans.TransID] = trans

			//uTrans := Transactions{
			//	trans.TransID,
			//	trans.UserID,
			//	trans.Name,
			//	trans.Weight,
			//	trans.Item,
			//	trans.TransDate,
			//}

			//if _, ok := userTrans[trans.UserID]; !ok {
			//	userTrans[trans.UserID] = append(make([]Transactions, 0), trans)
			//} else {
			//	userTrans[trans.UserID] = append(userTrans[trans.UserID], trans)
			//}

		}
	}
	return transactions //, userTrans
}

func GetUserTransactionsByItem(db *sql.DB, pageIndex int, recordsPerPage int, userID int, item string) map[int][]Transactions {

	var trans Transactions
	//var id int

	transactions := make(map[int][]Transactions)

	query := fmt.Sprintf(`
								SELECT t.id, u.id, u.name, t.trans_date, t.weight, i.name
								FROM my_db.Users u
								JOIN my_db.Transactions t
								ON u.id = t.user_id
								JOIN my_db.Items i
								ON i.id = t.item_id
								WHERE u.id = '%d' AND i.name = '%s'
								ORDER BY t.id
								LIMIT %d OFFSET %d
								`, userID, item, recordsPerPage, pageIndex*recordsPerPage)

	if results, err := db.Query(query); err != nil {
		log.Panicln(err.Error())
	} else {
		for results.Next() {
			// Scan() copy each row of data from db and assign to the address specified
			err = results.Scan(&trans.TransID, &trans.UserID, &trans.Name, &trans.TransDate, &trans.Weight, &trans.Item)
			//fmt.Println(user.ApiKey, "db")
			//fmt.Println(input.ApiKey, "input")
			if err != nil {
				//fmt.Println(course.CourseCode, course.CourseName, course.Description, module.ModuleCode, module.ModuleName, module.Description)
				log.Panicln(err.Error())

			}

			//if _, ok := transactions[trans.UserID]; !ok {
			//	transactions[trans.UserID] = append(make([]Transactions, 0), trans)
			//} else {
			//	transactions[trans.UserID] = append(userTrans[trans.UserID], trans)
			//}
			//transactions[trans.userID] = trans
			transactions[trans.UserID] = append(transactions[trans.UserID], trans)

		}
	}
	return transactions
}

func GetUserTransactions(db *sql.DB, pageIndex int, recordsPerPage int, userID int) map[int][]Transactions {

	var trans Transactions
	//var id int

	transactions := make(map[int][]Transactions)

	query := fmt.Sprintf(`
								SELECT t.id, u.id, u.name, t.trans_date, t.weight, i.name
								FROM my_db.Users u
								JOIN my_db.Transactions t
								ON u.id = t.user_id
								JOIN my_db.Items i
								ON i.id = t.item_id
								WHERE u.id = '%d'
								ORDER BY t.id
								LIMIT %d OFFSET %d
								`, userID, recordsPerPage, pageIndex*recordsPerPage)

	if results, err := db.Query(query); err != nil {
		log.Panicln(err.Error())
	} else {
		for results.Next() {
			// Scan() copy each row of data from db and assign to the address specified
			err = results.Scan(&trans.TransID, &trans.UserID, &trans.Name, &trans.TransDate, &trans.Weight, &trans.Item)
			//fmt.Println(user.ApiKey, "db")
			//fmt.Println(input.ApiKey, "input")
			if err != nil {
				//fmt.Println(course.CourseCode, course.CourseName, course.Description, module.ModuleCode, module.ModuleName, module.Description)
				log.Panicln(err.Error())

			}

			//if _, ok := transactions[trans.UserID]; !ok {
			//	transactions[trans.UserID] = append(make([]Transactions, 0), trans)
			//} else {
			//	transactions[trans.UserID] = append(userTrans[trans.UserID], trans)
			//}
			//transactions[trans.userID] = trans
			transactions[trans.UserID] = append(transactions[trans.UserID], trans)

		}
	}
	return transactions
}

func GetAllUsers(db *sql.DB, pageIndex int, recordsPerPage int) map[int]UserInfo {

	var user UserInfo
	//var id int
	//var results []map[int]UserInfo

	users := make(map[int]UserInfo)

	query := fmt.Sprintf(`
								SELECT id, phone, name, points, last_login
								FROM Users
								ORDER BY id
								LIMIT %d OFFSET %d
								`, recordsPerPage, pageIndex*recordsPerPage)

	if results, err := db.Query(query); err != nil {
		log.Panicln(err.Error())
	} else {
		for results.Next() {
			// Scan() copy each row of data from db and assign to the address specified
			err = results.Scan(&user.UserID, &user.Phone, &user.Name, &user.Points, &user.LastLogin)
			//fmt.Println(user.ApiKey, "db")
			//fmt.Println(input.ApiKey, "input")
			if err != nil {
				//fmt.Println(course.CourseCode, course.CourseName, course.Description, module.ModuleCode, module.ModuleName, module.Description)
				log.Panicln(err.Error())

			}
			users[user.UserID] = user
		}
	}
	//results = append(results, users)
	return users
}

func AddUser(db *sql.DB, name string, phone string, password string) {
	//userExist := checkExistUser(db, phone)
	//if !userExist {
	addUser(db, name, phone, password)
	//}

	//return userExist

}

//func checkExistUser(db *sql.DB, phone string) bool {
//
//	user := struct {
//		Name   string
//		ID     string
//		APIKey string
//	}{}
//
//	userExist := false
//
//	query := fmt.Sprintf(`
//								SELECT name, nric, api_key
//								FROM Users
//								WHERE nric='%s'
//								`, id)
//	if err := db.QueryRow(query).Scan(&user.Name, &user.ID, &user.APIKey); err != nil {
//		log.Panicln(err.Error())
//	} else {
//		if user.ID != "" && user.APIKey != apiKey { // existing user with wrong apikey
//			log.Panicln(errors.New("invalid api key"))
//		} else {
//			userExist = true
//		}
//	}
//
//	//if results, err := db.Query(query); err != nil {
//	//	log.Panicln(err.Error())
//	//} else {
//	//	if results.Next() {
//	//		// Scan() copy each row of data from db and assign to the address specified
//	//		err = results.Scan(&user.Name, &user.ID, &user.APIKey)
//	//		//fmt.Println(user.ApiKey, "db")
//	//		//fmt.Println(input.ApiKey, "input")
//	//		if err != nil {
//	//			//fmt.Println(course.CourseCode, course.CourseName, course.Description, module.ModuleCode, module.ModuleName, module.Description)
//	//			log.Panicln(err.Error())
//	//
//	//		} else {
//	//			if user.ID != "" && user.ApiKey != input.ApiKey { // existing user with wrong apikey
//	//				log.Panicln(errors.New("invalid api key"))
//	//			} else {
//	//				userExist = true
//	//			}
//	//		}
//	//	}
//	//}
//
//	//if !userExist {
//	//	addUser(db, input)
//	//}
//
//	return userExist
//
//}

func addUser(db *sql.DB, name string, phone string, password string) {
	query := fmt.Sprintf(`
						INSERT INTO Users 
						(phone, name, password) VALUES
						('%s', '%s', '%s')
						`, phone, name, password)
	_, err := db.Query(query)
	if err != nil {
		log.Panicln(err.Error())
	}
}

func RetrievePoints(db *sql.DB, phone string) (int, int) {
	userID := getUserID(db, phone)
	return userID, retrievePointsUser(db, userID)
}

func retrievePointsUser(db *sql.DB, userID int) int {

	var userPoints int
	query := fmt.Sprintf(`
						SELECT points
						FROM Users
						WHERE id = '%d'
						`, userID)
	//fmt.Println(query)

	if err := db.QueryRow(query).Scan(&userPoints); err != nil {
		log.Panicln(err.Error())
	}

	//if results, err := db.Query(query); err != nil {
	//	log.Panicln(err.Error())
	//} else {
	//	if results.Next() {
	//
	//		// Scan() copy each row of data from db and assign to the address specified
	//		err = results.Scan(&UserPoints)
	//
	//		if err != nil {
	//			log.Panicln(err.Error())
	//
	//		}
	//	}
	//}
	return userPoints
}

func AddTransaction(db *sql.DB, phone string, item string, points int, weight int) {
	itemID := getItemID(db, item)
	userID := getUserID(db, phone)
	addToTransaction(db, userID, itemID, weight)
	// after a transaction will add points to the user
	updateUserPoints(db, userID, points, false)
}

func getUserID(db *sql.DB, phone string) int {
	var userID int

	query := fmt.Sprintf(`
						SELECT id
						FROM Users
						WHERE phone = '%s'
						`, phone)
	//_, err := db.Query(query)
	if err := db.QueryRow(query).Scan(&userID); err != nil {
		log.Panicln(err.Error())
	}

	return userID
}

func getItemID(db *sql.DB, item string) int {

	var itemID int

	query := fmt.Sprintf(`
						SELECT id
						FROM Items
						WHERE name = '%s'
						`, item)
	//_, err := db.Query(query)
	if err := db.QueryRow(query).Scan(&itemID); err != nil {
		log.Panicln(err.Error())
	}

	//if results, err := db.Query(query); err != nil {
	//	log.Panicln(err.Error())
	//} else {
	//
	//	if results.Next() {
	//		// Scan() copy each row of data from db and assign to the address specified
	//		err = results.Scan(&itemID)
	//		//fmt.Println(user.ApiKey, "db")
	//		//fmt.Println(input.ApiKey, "input")
	//		if err != nil {
	//			//fmt.Println(course.CourseCode, course.CourseName, course.Description, module.ModuleCode, module.ModuleName, module.Description)
	//			log.Panicln(err.Error())
	//
	//		}
	//	}
	//}

	return itemID
}

func addToTransaction(db *sql.DB, userID int, itemID int, weight int) {
	query := fmt.Sprintf(`
						INSERT INTO Transactions
						(user_id, item_id, weight) VALUES
						('%d', '%d', '%d')
						`, userID, itemID, weight)
	_, err := db.Query(query)
	if err != nil {
		log.Panicln(err.Error())
	}
}

func updateUserPoints(db *sql.DB, id int, points int, isExchange bool) {

	var newPoints int
	currPoints := retrievePointsUser(db, id)

	if isExchange {
		newPoints = currPoints - points
	} else {
		newPoints = points + currPoints
	}

	query := fmt.Sprintf(`
						UPDATE Users
						SET points = '%d'
						WHERE id = '%d'
						`, newPoints, id)
	_, err := db.Query(query)
	if err != nil {
		log.Panicln(err.Error())
	}
}

func RedeemVoucher(db *sql.DB, phone string, voucherID string) (int, int, bool) {

	userID := getUserID(db, phone)
	redeem, valid := validVoucher(db, userID, voucherID)
	if valid { // if valid voucher update db
		updateVoucher(db, userID, voucherID)
		//updateUserPoints(db, id, points, true)
	}
	return userID, redeem, valid
}

func updateVoucher(db *sql.DB, id int, voucherID string) {
	query := fmt.Sprintf(`
						UPDATE Vouchers
						SET redeem = 1
						WHERE user_id = '%d' AND voucher_id = '%s'
						`, id, voucherID)
	_, err := db.Query(query)
	if err != nil {
		log.Panicln(err.Error())
	}

}

func validVoucher(db *sql.DB, id int, voucherID string) (int, bool) {

	valid := false
	query := fmt.Sprintf(`
						SELECT voucher_id, redeem
						FROM Vouchers
						WHERE user_id = '%d' && voucher_id = '%s'
						`, id, voucherID)

	var vID string
	var redeem int

	if err := db.QueryRow(query).Scan(&vID, &redeem); err != nil {
		log.Panicln(err.Error())
	}

	if vID == voucherID {
		valid = true
	}

	//if results, err := db.Query(query); err != nil {
	//	log.Panicln(err.Error())
	//} else {
	//	if results.Next() {
	//
	//		// Scan() copy each row of data from db and assign to the address specified
	//		err = results.Scan(&vID)
	//		//fmt.Println(user.ApiKey, "db")
	//		//fmt.Println(input.ApiKey, "input")
	//		if err != nil {
	//			//fmt.Println(course.CourseCode, course.CourseName, course.Description, module.ModuleCode, module.ModuleName, module.Description)
	//			log.Panicln(err.Error())
	//
	//		}
	//		if vID == voucherID {
	//			valid = true
	//		}
	//	}
	//}

	return redeem, valid
}

//func IssueVoucher(db *sql.DB, id int, voucherID string, points int) bool {
//	//valid := validVoucher(db, id, voucherID)
//	//if valid { // if valid voucher update db
//	//	updateVoucher(db, id, voucherID)
//	updateUserPoints(db, id, points, true)
//	//}
//	//return valid
//}

func AddVoucher(db *sql.DB, phone string, voucherID string, amount int, points int) (int, string, int, int) {
	userID := getUserID(db, phone)
	addVoucherUser(db, userID, voucherID, amount)
	updateUserPoints(db, userID, points, true)
	return userID, voucherID, amount, points
}

func addVoucherUser(db *sql.DB, id int, voucherID string, amount int) {
	query := fmt.Sprintf(`
						INSERT INTO Vouchers
						(user_id, voucher_amt, voucher_id) VALUES
						('%d', '%d', '%s')
						`, id, amount, voucherID)
	_, err := db.Query(query)
	if err != nil {
		log.Panicln(err.Error())
	}
}

func RetrieveVoucherStatus(db *sql.DB, phone string, voucherID string) (int, string, bool, bool) {
	userID := getUserID(db, phone)
	//redeem, valid := validVoucher(db, userID, voucherID)
	//if valid { // if valid voucher update db
	valid, redeem := voucherStatus(db, userID, voucherID)
	return userID, voucherID, redeem, valid
}

func voucherStatus(db *sql.DB, id int, voucherID string) (bool, bool) {
	var redeem int

	query := fmt.Sprintf(`
						SELECT redeem
						FROM Vouchers
						WHERE voucher_id = '%s' AND user_id = '%d'
						`, voucherID, id)
	//_, err := db.Query(query)
	if err := db.QueryRow(query).Scan(&redeem); err != nil { // no voucher ID
		//log.Panicln(err.Error())
		return false, false

	}

	if redeem == 1 {
		return true, true
	}

	return true, false
}

func UpdateKey(db *sql.DB, username string, apiKey string) {
	query := fmt.Sprintf(`
						UPDATE Api
						SET api_key = '%s'
						WHERE username = '%s'
						`, apiKey, username)
	_, err := db.Query(query)
	if err != nil {
		log.Panicln(err.Error())
	}
}

func AddKey(db *sql.DB, username string, apiKey string) {
	query := fmt.Sprintf(`
						INSERT INTO Api
						(username, api_key) VALUES
						('%s','%s')
						`, username, apiKey)
	_, err := db.Query(query)
	if err != nil {
		log.Panicln(err.Error())
	}
}

//func UpdateVoucherStatus(db *sql.DB, id int, voucherID string) bool {
//	return voucherStatus(db, id, voucherID)
//}
