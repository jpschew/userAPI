package api

import (
	"net/http"
	"strconv"
)

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
		} else {
			if !(string(phone[0]) == "8" || string(phone[0]) == "9") { // if first number is not 8 or 9
				return false
			}
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

// validKey checks for a valid key to secure the REST API
// so that only authenticated user can use the REST API
func validKey(req *http.Request) (string, bool) {

	// check if apikey and username field is present
	if req.Header.Get("apiKey") == "" || req.Header.Get("username") == "" {
		//	notFoundStatusJSON(res, false, "Both API Key and username need to be provided", make(map[string]interface{}))
		//	return
		return "Both API Key and username need to be provided.", false
	}

	// check for valid access token
	if !(req.Header.Get("apiKey") == userAPI[req.Header.Get("username")]) {
		//notFoundStatusJSON(res, false, "Invalid API Key", make(map[string]interface{}))
		return "Invalid API Key.", false
	}

	return "", true
}
