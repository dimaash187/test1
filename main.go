package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	H "test1/handlers"

	"github.com/gorilla/mux"
)

func main() {

	argsPtr := flag.String("transactions", "data.json", "")
	flag.Parse()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/transactions", H.GetTransactions(*argsPtr)).Methods("GET")
	router.HandleFunc("/transactions_by_timestamp", H.GetTransactionsByTimeStamp(*argsPtr)).Methods("GET")
	router.HandleFunc("/transactions_by_timestamp", H.PostTransactionsByTimeStamp).Methods("POST")

	fmt.Println("Serving transactions on port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
