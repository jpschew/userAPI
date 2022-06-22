package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"userAPI/api"
)

func main() {

	router := mux.NewRouter()

	// route and handler functions for user profile
	// without specifying the methods using .Method()
	// it will accept all methods
	router.HandleFunc("/api/v1/", api.Home).Methods("GET")

	// add user
	router.HandleFunc("/api/v1/adduser", api.AddUser).Methods("POST")
	// add transaction to user
	router.HandleFunc("/api/v1/addtransaction", api.AddTransaction).Methods("POST")
	// retrieve points from user
	router.HandleFunc("/api/v1/retrieve", api.RetrieveUserPoints).Methods("POST")
	// add voucher to user
	router.HandleFunc("/api/v1/addvoucher", api.AddUserVoucher).Methods("POST")
	// check voucher status of an user using voucher id
	router.HandleFunc("/api/v1/checkvoucher", api.VoucherStatus).Methods("POST")
	// redeem voucher of an user
	router.HandleFunc("/api/v1/redeemvoucher", api.RedeemVoucher).Methods("PUT")

	// get all users
	router.HandleFunc("/api/v1/getusers", api.GetAllUsers).Methods("GET")
	// get all transactions
	router.HandleFunc("/api/v1/gettransactions", api.GetAllTransactions).Methods("GET")
	// get all user transactions
	router.HandleFunc("/api/v1/gettransactions/{userid}", api.GetAllTransactions).Methods("GET")
	// get all user transactions filtered by item
	router.HandleFunc("/api/v1/gettransactions/{userid}/{itemid}", api.GetAllTransactions).Methods("GET")

	// route and handler functions for api key
	// generate apikey to user
	router.HandleFunc("/api/v1/genkey", api.GenKey).Methods("GET")
	// add apikey to user
	router.HandleFunc("/api/v1/addkey", api.AddUpdateKey).Methods("PUT")

	//log.Fatalln(http.ListenAndServe(":5001", router))
	log.Fatalln(http.ListenAndServeTLS(":5000", "./ssl/localhost.cert.pem", "./ssl/localhost.key.pem", router))
}
