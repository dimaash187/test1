package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
)

type Transaction struct {
	ID                  int    `json:"id"`
	Amount              int    `json:"amount"`
	MessageType         string `json:"conversation_type"`
	CreatedAt           string `json:"created_at"`
	TransactionID       int    `json:"created_at"`
	PAN                 string `json:"pan"`
	TransactionCategory string `json:"transaction_category"`
	PostedTimeStamp     string `json:"posted_timestamp"`
	TransactionType     string `json:"transaction_type"`
	SendingAccount      int    `json:"sending_account"`
	ReceivingAccount    int    `json:"receiving_account"`
	TransactionNote     string `json:"transaction_note"`
}

var transactions []Transaction

// Orders transactions by PostedTimeStamp in DESCdending order
func orderByTimeStamp(transdata []Transaction) []Transaction {
	sort.Slice(transdata, func(i, j int) bool {
		return transdata[j].PostedTimeStamp < transdata[i].PostedTimeStamp
	})
	return transdata
}

// Masks the string from Left to Right leaving out last 4 characters
func maskLeft(s string) string {
	rs := []rune(s)
	for i := 0; i < len(s)-4; i++ {
		rs[i] = '*'
	}
	return string(rs)
}

// Custom UnmarshalJSON function which uses a Transaction alias
// to convert the "PAN" of type inteeger into string, so that it can further be masked
// In case if byte array contains "PAN" in form of a string ( case being unit tests when testing output from handler )
// then will simply use the string as is without attempting a int to string conversion
func (t *Transaction) UnmarshalJSON(data []byte) error {
	type Alias Transaction
	aux := &struct {
		PAN int `json:"pan"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}
	aux2 := &struct {
		PAN string `json:"pan"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}

	if err := json.Unmarshal(data, &aux2); err != nil {
		if err.Error() == "json: cannot unmarshal number into Go struct field .pan of type string" {
			if err2 := json.Unmarshal(data, &aux); err2 != nil {
				log.Print(err2)
				return err2
			}
			t.PAN = maskLeft(strconv.Itoa(aux.PAN))
			return nil
		}
		log.Print(err)
		return err
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		if err.Error() == "json: cannot unmarshal string into Go struct field .pan of type int" {
			// log.Print(err.Error())
			if err2 := json.Unmarshal(data, &aux2); err2 != nil {
				log.Print(err2)
				return err2
			}
			t.PAN = aux2.PAN
			return nil
		}
		log.Print(err)
		return err
	}
	t.PAN = maskLeft(strconv.Itoa(aux.PAN))
	return nil
}

// Handler which takes 
func GetTransactions(datafile string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// mock data
		// Lets use mock data from a file "data.json" instead of inline
		// transactions = append(transactions, Transaction{ID: 1, Amount: 200, MessageType: "Debit", CreatedAt: "2020-06-11T19:11:24+00:00", TransactionID: 10, PAN: 4080230386144446, TransactionCategory: "Grocery", PostedTimeStamp: "2020-06-11T19:11:24+00:00", TransactionType: "POS", SendingAccount: 39203, ReceivingAccount: 993020, TransactionNote: "Merchant 00308281"}, Transaction{ID: 2, Amount: 499, MessageType: "Credit", CreatedAt: "2020-06-11T19:11:24+00:00", TransactionID: 12, PAN: 5166697943434128, TransactionCategory: "Food and Beverage", PostedTimeStamp: "2020-06-11T19:11:24+00:00", TransactionType: "POS", SendingAccount: 39400, ReceivingAccount: 9233020, TransactionNote: "Jimmys Corn and Cheese refund"}, Transaction{ID: 3, Amount: 20000, MessageType: "Debit", CreatedAt: "2020-06-11T19:11:24+00:00", TransactionID: 17, PAN: 5488452462266852, TransactionCategory: "ATM", PostedTimeStamp: "2020-06-11T19:11:24+00:00", TransactionType: "POS", SendingAccount: 99302, ReceivingAccount: 11209, TransactionNote: "ATM #39902 Burrard Street"}, Transaction{ID: 4, Amount: 8839, MessageType: "Debit", CreatedAt: "2020-06-11T19:11:24+00:00", TransactionID: 10, PAN: 4954335252282726, TransactionCategory: "Automotive", PostedTimeStamp: "2020-06-11T19:11:24+00:00", TransactionType: "POS", SendingAccount: 83839, ReceivingAccount: 9233020, TransactionNote: "Muffler Bearings Inc."}, Transaction{ID: 5, Amount: 6173, MessageType: "Debit", CreatedAt: "2020-06-21T20:11:24+00:00", TransactionID: 21, PAN: 4844085301308048, TransactionCategory: "Household", PostedTimeStamp: "2020-06-21T20:11:24+00:00", TransactionType: "POS", SendingAccount: 9018, ReceivingAccount: 9222020, TransactionNote: "Amazon Inc."}, Transaction{ID: 6, Amount: 9018, MessageType: "Debit", CreatedAt: "2020-06-14T21:11:24+00:00", TransactionID: 33, PAN: 4090070794938361, TransactionCategory: "Electronics", PostedTimeStamp: "2020-06-14T21:11:24+00:00", TransactionType: "POS", SendingAccount: 83339, ReceivingAccount: 9233021, TransactionNote: "Apple Store"}, Transaction{ID: 7, Amount: 1275, MessageType: "Debit", CreatedAt: "2020-06-09T11:11:24+00:00", TransactionID: 12, PAN: 4807678678904632, TransactionCategory: "Cryptocurrency", PostedTimeStamp: "2020-06-09T11:11:24+00:00", TransactionType: "POS", SendingAccount: 83839, ReceivingAccount: 9233020, TransactionNote: "Bittrex Inc."}, Transaction{ID: 8, Amount: 2167, MessageType: "Debit", CreatedAt: "2020-06-18T15:11:24+00:00", TransactionID: 44, PAN: 4673062314928753, TransactionCategory: "Food and Beverage", PostedTimeStamp: "2020-06-18T15:11:24+00:00", TransactionType: "POS", SendingAccount: 8569, ReceivingAccount: 533020, TransactionNote: "Giorgios Pizza Ltd."}, Transaction{ID: 9, Amount: 3178, MessageType: "Debit", CreatedAt: "2020-05-01T14:11:24+00:00", TransactionID: 90, PAN: 5109473381765575, TransactionCategory: "Internet Services", PostedTimeStamp: "2020-05-01T14:11:24+00:00", TransactionType: "POS", SendingAccount: 63639, ReceivingAccount: 4233010, TransactionNote: "Google"}, Transaction{ID: 10, Amount: 7718, MessageType: "Debit", CreatedAt: "2020-07-11T19:11:24+00:00", TransactionID: 109, PAN: 5158563621617519, TransactionCategory: "Health Services", PostedTimeStamp: "2020-07-11T19:11:25+00:00", TransactionType: "POS", SendingAccount: 13839, ReceivingAccount: 244020, TransactionNote: "Vons Nails"})

		// Open data.json
		jsonFile, err := os.Open(datafile)
		// if we os.Open returns an error then handle it
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()

		// Lets sort by PostedTimeStamp field in descing order
		orderByTimeStamp(transactions)

		log.Println("Current Transactions:")
		json.NewEncoder(w).Encode(transactions)

	}
}

func GetTransactionsByTimeStamp(datafile string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Open data.json
		jsonFile, err := os.Open(datafile)
		// if we os.Open returns an error then handle it
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()

		// read jsonFile as a byte array.
		byteValue, _ := ioutil.ReadAll(jsonFile)

		// Lets create the array of Transactions locally within the scope of this func
		var transdata []Transaction

		// we unmarshal byteValue which contains
		// jsonFile's content into 'transdata' which we defined above
		if err := json.Unmarshal([]byte(byteValue), &transdata); err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Lets sort by PostedTimeStamp field in descing order
		orderByTimeStamp(transdata)

		log.Println("Current Transactions:")
		json.NewEncoder(w).Encode(transdata)
	}
}

func PostTransactionsByTimeStamp(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var transdata []Transaction
	if err := decoder.Decode(&transdata); err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Lets sort by PostedTimeStamp field in descing order
	orderByTimeStamp(transdata)

	log.Println("Current Transactions:")
	json.NewEncoder(w).Encode(transdata)
}
