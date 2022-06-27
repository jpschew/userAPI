package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	// implementation of GO's database/sql/driver interface
	// only need to import and use the GO's database/sql API
	_ "github.com/go-sql-driver/mysql"
)

type UserInfo struct {
	UserID    int
	Phone     string
	Name      string
	Password  string
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

type RedeemVoucherInfo struct {
	UserID   int            `json:"userID"`
	Vouchers map[string]int `json:"vouchers"`
	Amount   int            `json:"amount"`
}

const (
	YYYYMMDDhhmmss = "2006-01-02 15:04:05"
)

// CreateDBConn creates a connection to mysql database given the driver name, dsn and db name.
func CreateDBConn(driver string, dsn string, dbName string) *sql.DB {

	// Use mysql as driver Name and a valid DSN as data SourceName:
	source := dsn + "/" + dbName
	db, err := sql.Open(driver, source)

	// handle error
	if err != nil {
		log.Panicln(err.Error())
	}

	return db
}

func InitAllUsers(db *sql.DB) map[string]string {

	users := make(map[string]string)

	var phone, password string

	query := fmt.Sprintf(`
								SELECT phone, password
								FROM Users
								`)

	if results, err := db.Query(query); err != nil {
		log.Panicln(err.Error())
	} else {

		for results.Next() {
			err = results.Scan(&phone, &password)

			if err != nil {
				log.Panicln(err.Error())
			}

			users[phone] = password
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

	transactions := make(map[int]Transactions)

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

			if err != nil {
				log.Panicln(err.Error())
			}

			transactions[trans.TransID] = trans
		}
	}
	return transactions
}

func GetUserTransactionsByItem(db *sql.DB, pageIndex int, recordsPerPage int, userID int, item string) map[int][]Transactions {

	var trans Transactions

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

			if err != nil {
				log.Panicln(err.Error())
			}

			transactions[trans.UserID] = append(transactions[trans.UserID], trans)
		}
	}
	return transactions
}

func GetUserTransactions(db *sql.DB, pageIndex int, recordsPerPage int, userID int) map[int][]Transactions {

	var trans Transactions

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

			if err != nil {
				log.Panicln(err.Error())
			}

			transactions[trans.UserID] = append(transactions[trans.UserID], trans)
		}
	}
	return transactions
}

func GetAllUsers(db *sql.DB, pageIndex int, recordsPerPage int) map[int]UserInfo {

	var user UserInfo

	users := make(map[int]UserInfo)

	query := fmt.Sprintf(`
								SELECT id, phone, name, password, points, last_login
								FROM Users
								ORDER BY id
								LIMIT %d OFFSET %d
								`, recordsPerPage, pageIndex*recordsPerPage)

	if results, err := db.Query(query); err != nil {
		log.Panicln(err.Error())
	} else {
		for results.Next() {
			// Scan() copy each row of data from db and assign to the address specified
			err = results.Scan(&user.UserID, &user.Phone, &user.Name, &user.Password, &user.Points, &user.LastLogin)
			if err != nil {
				log.Panicln(err.Error())
			}

			users[user.UserID] = user
		}
	}

	return users
}

func AddUser(db *sql.DB, name string, phone string, password string) map[int]UserInfo {

	// add user first before getting the user info to pass back
	// cannot use goroutines here
	addUser(db, name, phone, password)
	user := getUser(db, phone)

	return user
}

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

func getUser(db *sql.DB, phone string) map[int]UserInfo {
	var user UserInfo

	userInfo := make(map[int]UserInfo)

	query := fmt.Sprintf(`
								SELECT id, phone, name, password, points, last_login
								FROM Users
								WHERE phone = '%s'
								`, phone)
	if err := db.QueryRow(query).Scan(&user.UserID, &user.Phone, &user.Name, &user.Password, &user.Points, &user.LastLogin); err != nil {
		log.Panicln(err.Error())
	}

	userInfo[user.UserID] = user

	return userInfo
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

	if err := db.QueryRow(query).Scan(&userPoints); err != nil {
		log.Panicln(err.Error())
	}

	return userPoints
}

func AddTransaction(db *sql.DB, phone string, item string, points int, weight int, wg *sync.WaitGroup) {

	defer wg.Done()

	itemID := getItemID(db, item)
	userID := getUserID(db, phone)

	var wgDB sync.WaitGroup

	defer wgDB.Wait()

	wgDB.Add(2)
	go addToTransaction(db, userID, itemID, weight, &wgDB)
	// after a transaction will add points to the user
	go updateUserPoints(db, userID, points, false, &wgDB)

}

func getUserID(db *sql.DB, phone string) int {
	var userID int

	query := fmt.Sprintf(`
						SELECT id
						FROM Users
						WHERE phone = '%s'
						`, phone)

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

	if err := db.QueryRow(query).Scan(&itemID); err != nil {
		log.Panicln(err.Error())
	}

	return itemID
}

func addToTransaction(db *sql.DB, userID int, itemID int, weight int, wg *sync.WaitGroup) {

	defer wg.Done()

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

func updateUserPoints(db *sql.DB, id int, points int, isExchange bool, wg *sync.WaitGroup) {

	defer wg.Done()

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
	}
	return userID, redeem, valid
}

func updateVoucher(db *sql.DB, id int, voucherID string) {
	query := fmt.Sprintf(`
						UPDATE Vouchers
						SET redeem = 1, time_updated = '%s'
						WHERE user_id = '%d' AND voucher_id = '%s'
						`, time.Now().Format(YYYYMMDDhhmmss), id, voucherID)
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
		//log.Panicln(err.Error())
		return redeem, valid
	}

	if vID == voucherID {
		valid = true
	}

	return redeem, valid
}

func AddVoucher(db *sql.DB, phone string, voucherID string, amount int, points int) (int, string, int, int) {

	var wg sync.WaitGroup

	defer wg.Wait()

	userID := getUserID(db, phone)

	wg.Add(2)
	go addVoucherUser(db, userID, voucherID, amount, &wg)
	go updateUserPoints(db, userID, points, true, &wg)

	return userID, voucherID, amount, points
}

func addVoucherUser(db *sql.DB, id int, voucherID string, amount int, wg *sync.WaitGroup) {

	defer wg.Done()

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

	if err := db.QueryRow(query).Scan(&redeem); err != nil { // no voucher ID
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

func GetUserValidVoucher(db *sql.DB, phone string) RedeemVoucherInfo {
	userID := getUserID(db, phone)
	return getValidVoucher(db, userID)
}

func getValidVoucher(db *sql.DB, userID int) RedeemVoucherInfo {

	var voucherInfo RedeemVoucherInfo

	voucherInfo.Vouchers = make(map[string]int)

	info := struct {
		vID string
		amt int
	}{}

	query := fmt.Sprintf(`
								SELECT user_id, voucher_id, voucher_amt
								FROM Vouchers
								WHERE redeem = 1 AND user_id = '%d'
								`, userID)

	if results, err := db.Query(query); err != nil {
		log.Panicln(err.Error())
	} else {
		for results.Next() {
			// Scan() copy each row of data from db and assign to the address specified
			err = results.Scan(&voucherInfo.UserID, &info.vID, &info.amt)
			if err != nil {
				log.Panicln(err.Error())
			}

			voucherInfo.Amount += info.amt
			voucherInfo.Vouchers[info.vID] = info.amt

		}
		//fmt.Println(voucherInfo.UserID)
		if voucherInfo.UserID == 0 { // means no records
			voucherInfo.UserID = userID
		}
	}

	return voucherInfo
}
