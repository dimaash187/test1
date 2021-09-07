package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"test1/handlers"
	H "test1/handlers"
	"testing"
)

func TestMaskLeftFunc(t *testing.T) {
	testFunc := handlers.MaskLeftFunc

	actual := testFunc("1234567890")
	if actual != "******7890" {
		t.Errorf("expected '******7890', got '%s'", actual)
	}
}

func TestOrderByTimeStampFunc(t *testing.T) {
	testFunc := handlers.OrderByTimeStampFunc

	var transactions []H.Transaction
	transactions = append(transactions, H.Transaction{ID: 1, Amount: 200, MessageType: "Debit", CreatedAt: "2020-06-11T19:11:24+00:00", TransactionID: 10, PAN: "4080230386144446", TransactionCategory: "Grocery", PostedTimeStamp: "2020-06-11T19:11:24+00:00", TransactionType: "POS", SendingAccount: 39203, ReceivingAccount: 993020, TransactionNote: "Merchant 00308281"}, H.Transaction{ID: 2, Amount: 499, MessageType: "Credit", CreatedAt: "2020-06-11T19:11:24+00:00", TransactionID: 12, PAN: "5166697943434128", TransactionCategory: "Food and Beverage", PostedTimeStamp: "2020-06-11T19:11:24+00:00", TransactionType: "POS", SendingAccount: 39400, ReceivingAccount: 9233020, TransactionNote: "Jimmys Corn and Cheese refund"}, H.Transaction{ID: 3, Amount: 20000, MessageType: "Debit", CreatedAt: "2020-06-11T19:11:24+00:00", TransactionID: 17, PAN: "5488452462266852", TransactionCategory: "ATM", PostedTimeStamp: "2020-06-11T19:11:24+00:00", TransactionType: "POS", SendingAccount: 99302, ReceivingAccount: 11209, TransactionNote: "ATM #39902 Burrard Street"}, H.Transaction{ID: 4, Amount: 8839, MessageType: "Debit", CreatedAt: "2020-06-11T19:11:24+00:00", TransactionID: 10, PAN: "4954335252282726", TransactionCategory: "Automotive", PostedTimeStamp: "2020-06-11T19:11:24+00:00", TransactionType: "POS", SendingAccount: 83839, ReceivingAccount: 9233020, TransactionNote: "Muffler Bearings Inc."}, H.Transaction{ID: 5, Amount: 6173, MessageType: "Debit", CreatedAt: "2020-06-21T20:11:24+00:00", TransactionID: 21, PAN: "4844085301308048", TransactionCategory: "Household", PostedTimeStamp: "2020-06-21T20:11:24+00:00", TransactionType: "POS", SendingAccount: 9018, ReceivingAccount: 9222020, TransactionNote: "Amazon Inc."}, H.Transaction{ID: 6, Amount: 9018, MessageType: "Debit", CreatedAt: "2020-06-14T21:11:24+00:00", TransactionID: 33, PAN: "4090070794938361", TransactionCategory: "Electronics", PostedTimeStamp: "2020-06-14T21:11:24+00:00", TransactionType: "POS", SendingAccount: 83339, ReceivingAccount: 9233021, TransactionNote: "Apple Store"}, H.Transaction{ID: 7, Amount: 1275, MessageType: "Debit", CreatedAt: "2020-06-09T11:11:24+00:00", TransactionID: 12, PAN: "4807678678904632", TransactionCategory: "Cryptocurrency", PostedTimeStamp: "2020-06-09T11:11:24+00:00", TransactionType: "POS", SendingAccount: 83839, ReceivingAccount: 9233020, TransactionNote: "Bittrex Inc."}, H.Transaction{ID: 8, Amount: 2167, MessageType: "Debit", CreatedAt: "2020-06-18T15:11:24+00:00", TransactionID: 44, PAN: "4673062314928753", TransactionCategory: "Food and Beverage", PostedTimeStamp: "2020-06-18T15:11:24+00:00", TransactionType: "POS", SendingAccount: 8569, ReceivingAccount: 533020, TransactionNote: "Giorgios Pizza Ltd."}, H.Transaction{ID: 9, Amount: 3178, MessageType: "Debit", CreatedAt: "2020-05-01T14:11:24+00:00", TransactionID: 90, PAN: "5109473381765575", TransactionCategory: "Internet Services", PostedTimeStamp: "2020-05-01T14:11:24+00:00", TransactionType: "POS", SendingAccount: 63639, ReceivingAccount: 4233010, TransactionNote: "Google"}, H.Transaction{ID: 10, Amount: 7718, MessageType: "Debit", CreatedAt: "2020-07-11T19:11:24+00:00", TransactionID: 109, PAN: "5158563621617519", TransactionCategory: "Health Services", PostedTimeStamp: "2020-07-11T19:11:25+00:00", TransactionType: "POS", SendingAccount: 13839, ReceivingAccount: 244020, TransactionNote: "Vons Nails"})
	want := []H.Transaction{
		{
			ID:                  10,
			Amount:              7718,
			MessageType:         "Debit",
			CreatedAt:           "2020-07-11T19:11:24+00:00",
			TransactionID:       109,
			PAN:                 "5158563621617519",
			TransactionCategory: "Health Services",
			PostedTimeStamp:     "2020-07-11T19:11:25+00:00",
			TransactionType:     "POS",
			SendingAccount:      13839,
			ReceivingAccount:    244020,
			TransactionNote:     "Vons Nails",
		},
		{
			ID:                  5,
			Amount:              6173,
			MessageType:         "Debit",
			CreatedAt:           "2020-07-11T19:11:24+00:00",
			TransactionID:       109,
			PAN:                 "4844085301308048",
			TransactionCategory: "Household",
			PostedTimeStamp:     "2020-06-21T20:11:24+00:00",
			TransactionType:     "POS",
			SendingAccount:      9018,
			ReceivingAccount:    9222020,
			TransactionNote:     "Amazon Inc.",
		},
		{
			ID:                  8,
			Amount:              2167,
			MessageType:         "Debit",
			CreatedAt:           "2020-07-11T19:11:24+00:00",
			TransactionID:       109,
			PAN:                 "4673062314928753",
			TransactionCategory: "Food and Beverage",
			PostedTimeStamp:     "2020-06-18T15:11:24+00:00",
			TransactionType:     "POS",
			SendingAccount:      8569,
			ReceivingAccount:    533020,
			TransactionNote:     "Giorgios Pizza Ltd.",
		},
		{
			ID:                  6,
			Amount:              9018,
			MessageType:         "Debit",
			CreatedAt:           "2020-07-11T19:11:24+00:00",
			TransactionID:       109,
			PAN:                 "4090070794938361",
			TransactionCategory: "Electronics",
			PostedTimeStamp:     "2020-06-14T21:11:24+00:00",
			TransactionType:     "POS",
			SendingAccount:      83339,
			ReceivingAccount:    9233021,
			TransactionNote:     "Apple Store",
		},
		{
			ID:                  1,
			Amount:              200,
			MessageType:         "Debit",
			CreatedAt:           "2020-07-11T19:11:24+00:00",
			TransactionID:       109,
			PAN:                 "4080230386144446",
			TransactionCategory: "Grocery",
			PostedTimeStamp:     "2020-06-11T19:11:24+00:00",
			TransactionType:     "POS",
			SendingAccount:      39203,
			ReceivingAccount:    993020,
			TransactionNote:     "Merchant 00308281",
		},
		{
			ID:                  3,
			Amount:              20000,
			MessageType:         "Debit",
			CreatedAt:           "2020-07-11T19:11:24+00:00",
			TransactionID:       109,
			PAN:                 "5488452462266852",
			TransactionCategory: "ATM",
			PostedTimeStamp:     "2020-06-11T19:11:24+00:00",
			TransactionType:     "POS",
			SendingAccount:      99302,
			ReceivingAccount:    11209,
			TransactionNote:     "ATM #39902 Burrard Street",
		},
		{
			ID:                  2,
			Amount:              499,
			MessageType:         "Credit",
			CreatedAt:           "2020-07-11T19:11:24+00:00",
			TransactionID:       109,
			PAN:                 "5166697943434128",
			TransactionCategory: "Food and Beverage",
			PostedTimeStamp:     "2020-06-11T19:11:24+00:00",
			TransactionType:     "POS",
			SendingAccount:      39400,
			ReceivingAccount:    9233020,
			TransactionNote:     "Jimmys Corn and Cheese refund",
		},
		{
			ID:                  4,
			Amount:              8839,
			MessageType:         "Debit",
			CreatedAt:           "2020-07-11T19:11:24+00:00",
			TransactionID:       109,
			PAN:                 "4954335252282726",
			TransactionCategory: "Automotive",
			PostedTimeStamp:     "2020-06-11T19:11:24+00:00",
			TransactionType:     "POS",
			SendingAccount:      83839,
			ReceivingAccount:    9233020,
			TransactionNote:     "Muffler Bearings Inc.",
		},
		{
			ID:                  7,
			Amount:              1275,
			MessageType:         "Debit",
			CreatedAt:           "2020-07-11T19:11:24+00:00",
			TransactionID:       109,
			PAN:                 "4807678678904632",
			TransactionCategory: "Cryptocurrency",
			PostedTimeStamp:     "2020-06-09T11:11:24+00:00",
			TransactionType:     "POS",
			SendingAccount:      83839,
			ReceivingAccount:    9233020,
			TransactionNote:     "Bittrex Inc.",
		},
		{
			ID:                  9,
			Amount:              3178,
			MessageType:         "Debit",
			CreatedAt:           "2020-07-11T19:11:24+00:00",
			TransactionID:       109,
			PAN:                 "5109473381765575",
			TransactionCategory: "Internet Services",
			PostedTimeStamp:     "2020-05-01T14:11:24+00:00",
			TransactionType:     "POS",
			SendingAccount:      63639,
			ReceivingAccount:    4233010,
			TransactionNote:     "Google",
		},
	}
	testFunc(transactions)

	for i := range want {
		if transactions[i].ID != want[i].ID {
			t.Errorf("ID expected '%d', got '%d'", want[i].ID, transactions[i].ID)
		}
		if transactions[i].Amount != want[i].Amount {
			t.Errorf("Amount expected '%d', got '%d'", want[i].Amount, transactions[i].Amount)
		}
		if transactions[i].MessageType != want[i].MessageType {
			t.Errorf("MessageType expected '%s', got '%s'", want[i].MessageType, transactions[i].MessageType)
		}
		// if transdata[i].TransactionID != want[i].TransactionID {
		// 	t.Errorf("TransactionID expected '%d', got '%d'", want[i].TransactionID, transdata[i].TransactionID)
		// }
		if transactions[i].PAN != want[i].PAN {
			t.Errorf("PAN expected '%s', got '%s'", want[i].PAN, transactions[i].PAN)
		}
		if transactions[i].TransactionCategory != want[i].TransactionCategory {
			t.Errorf("TransactionCategory expected '%s', got '%s'", want[i].TransactionCategory, transactions[i].TransactionCategory)
		}
		if transactions[i].TransactionNote != want[i].TransactionNote {
			t.Errorf("TransactionNote expected '%s', got '%s'", want[i].TransactionNote, transactions[i].TransactionNote)
		}
		if transactions[i].SendingAccount != want[i].SendingAccount {
			t.Errorf("SendingAccount expected '%d', got '%d'", want[i].SendingAccount, transactions[i].SendingAccount)
		}
		if transactions[i].ReceivingAccount != want[i].ReceivingAccount {
			t.Errorf("ReceivingAccount expected '%d', got '%d'", want[i].ReceivingAccount, transactions[i].ReceivingAccount)
		}
		if transactions[i].TransactionType != want[i].TransactionType {
			t.Errorf("TransactionType expected '%s', got '%s'", want[i].TransactionType, transactions[i].TransactionType)
		}
		if transactions[i].PostedTimeStamp != want[i].PostedTimeStamp {
			t.Errorf("PostedTimeStamp expected '%s', got '%s'", want[i].PostedTimeStamp, transactions[i].PostedTimeStamp)
		}
	}

}

func TestGetTransactions(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/transactions_by_timestamp", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(H.GetTransactionsByTimeStamp("data.json"))

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	fmt.Println("GOT: " + rr.Body.String())
	var transdata []H.Transaction
	if err := json.Unmarshal([]byte(rr.Body.String()), &transdata); err != nil {
		t.Errorf("error unmarshing json %s",
			err)
	}

	want := []H.Transaction{
		{
			ID:                  10,
			Amount:              7718,
			MessageType:         "Debit",
			CreatedAt:           "2020-07-11T19:11:24+00:00",
			TransactionID:       109,
			PAN:                 "************7519",
			TransactionCategory: "Health Services",
			PostedTimeStamp:     "2020-07-11T19:11:25+00:00",
			TransactionType:     "POS",
			SendingAccount:      13839,
			ReceivingAccount:    244020,
			TransactionNote:     "Vons Nails",
		},
		{
			ID:                  5,
			Amount:              6173,
			MessageType:         "Debit",
			CreatedAt:           "2020-07-11T19:11:24+00:00",
			TransactionID:       109,
			PAN:                 "************8048",
			TransactionCategory: "Household",
			PostedTimeStamp:     "2020-06-21T20:11:24+00:00",
			TransactionType:     "POS",
			SendingAccount:      9018,
			ReceivingAccount:    9222020,
			TransactionNote:     "Amazon Inc.",
		},
		{
			ID:                  8,
			Amount:              2167,
			MessageType:         "Debit",
			CreatedAt:           "2020-07-11T19:11:24+00:00",
			TransactionID:       109,
			PAN:                 "************8753",
			TransactionCategory: "Food and Beverage",
			PostedTimeStamp:     "2020-06-18T15:11:24+00:00",
			TransactionType:     "POS",
			SendingAccount:      8569,
			ReceivingAccount:    533020,
			TransactionNote:     "Giorgios Pizza Ltd.",
		},
		{
			ID:                  6,
			Amount:              9018,
			MessageType:         "Debit",
			CreatedAt:           "2020-07-11T19:11:24+00:00",
			TransactionID:       109,
			PAN:                 "************8361",
			TransactionCategory: "Electronics",
			PostedTimeStamp:     "2020-06-14T21:11:24+00:00",
			TransactionType:     "POS",
			SendingAccount:      83339,
			ReceivingAccount:    9233021,
			TransactionNote:     "Apple Store",
		},
		{
			ID:                  1,
			Amount:              200,
			MessageType:         "Debit",
			CreatedAt:           "2020-07-11T19:11:24+00:00",
			TransactionID:       109,
			PAN:                 "************4446",
			TransactionCategory: "Grocery",
			PostedTimeStamp:     "2020-06-11T19:11:24+00:00",
			TransactionType:     "POS",
			SendingAccount:      39203,
			ReceivingAccount:    993020,
			TransactionNote:     "Merchant 00308281",
		},
		{
			ID:                  3,
			Amount:              20000,
			MessageType:         "Debit",
			CreatedAt:           "2020-07-11T19:11:24+00:00",
			TransactionID:       109,
			PAN:                 "************6852",
			TransactionCategory: "ATM",
			PostedTimeStamp:     "2020-06-11T19:11:24+00:00",
			TransactionType:     "POS",
			SendingAccount:      99302,
			ReceivingAccount:    11209,
			TransactionNote:     "ATM #39902 Burrard Street",
		},
		{
			ID:                  2,
			Amount:              499,
			MessageType:         "Credit",
			CreatedAt:           "2020-07-11T19:11:24+00:00",
			TransactionID:       109,
			PAN:                 "************4128",
			TransactionCategory: "Food and Beverage",
			PostedTimeStamp:     "2020-06-11T19:11:24+00:00",
			TransactionType:     "POS",
			SendingAccount:      39400,
			ReceivingAccount:    9233020,
			TransactionNote:     "Jimmys Corn and Cheese refund",
		},
		{
			ID:                  4,
			Amount:              8839,
			MessageType:         "Debit",
			CreatedAt:           "2020-07-11T19:11:24+00:00",
			TransactionID:       109,
			PAN:                 "************2726",
			TransactionCategory: "Automotive",
			PostedTimeStamp:     "2020-06-11T19:11:24+00:00",
			TransactionType:     "POS",
			SendingAccount:      83839,
			ReceivingAccount:    9233020,
			TransactionNote:     "Muffler Bearings Inc.",
		},
		{
			ID:                  7,
			Amount:              1275,
			MessageType:         "Debit",
			CreatedAt:           "2020-07-11T19:11:24+00:00",
			TransactionID:       109,
			PAN:                 "************4632",
			TransactionCategory: "Cryptocurrency",
			PostedTimeStamp:     "2020-06-09T11:11:24+00:00",
			TransactionType:     "POS",
			SendingAccount:      83839,
			ReceivingAccount:    9233020,
			TransactionNote:     "Bittrex Inc.",
		},
		{
			ID:                  9,
			Amount:              3178,
			MessageType:         "Debit",
			CreatedAt:           "2020-07-11T19:11:24+00:00",
			TransactionID:       109,
			PAN:                 "************5575",
			TransactionCategory: "Internet Services",
			PostedTimeStamp:     "2020-05-01T14:11:24+00:00",
			TransactionType:     "POS",
			SendingAccount:      63639,
			ReceivingAccount:    4233010,
			TransactionNote:     "Google",
		},
	}

	for i := range want {
		if transdata[i].ID != want[i].ID {
			t.Errorf("ID expected '%d', got '%d'", want[i].ID, transdata[i].ID)
		}
		if transdata[i].Amount != want[i].Amount {
			t.Errorf("Amount expected '%d', got '%d'", want[i].Amount, transdata[i].Amount)
		}
		if transdata[i].MessageType != want[i].MessageType {
			t.Errorf("MessageType expected '%s', got '%s'", want[i].MessageType, transdata[i].MessageType)
		}
		// if transdata[i].TransactionID != want[i].TransactionID {
		// 	t.Errorf("TransactionID expected '%d', got '%d'", want[i].TransactionID, transdata[i].TransactionID)
		// }
		if transdata[i].PAN != want[i].PAN {
			t.Errorf("PAN expected '%s', got '%s'", want[i].PAN, transdata[i].PAN)
		}
		if transdata[i].TransactionCategory != want[i].TransactionCategory {
			t.Errorf("TransactionCategory expected '%s', got '%s'", want[i].TransactionCategory, transdata[i].TransactionCategory)
		}
		if transdata[i].TransactionNote != want[i].TransactionNote {
			t.Errorf("TransactionNote expected '%s', got '%s'", want[i].TransactionNote, transdata[i].TransactionNote)
		}
		if transdata[i].SendingAccount != want[i].SendingAccount {
			t.Errorf("SendingAccount expected '%d', got '%d'", want[i].SendingAccount, transdata[i].SendingAccount)
		}
		if transdata[i].ReceivingAccount != want[i].ReceivingAccount {
			t.Errorf("ReceivingAccount expected '%d', got '%d'", want[i].ReceivingAccount, transdata[i].ReceivingAccount)
		}
		if transdata[i].TransactionType != want[i].TransactionType {
			t.Errorf("TransactionType expected '%s', got '%s'", want[i].TransactionType, transdata[i].TransactionType)
		}
		if transdata[i].PostedTimeStamp != want[i].PostedTimeStamp {
			t.Errorf("PostedTimeStamp expected '%s', got '%s'", want[i].PostedTimeStamp, transdata[i].PostedTimeStamp)
		}

		// if !reflect.DeepEqual(want[i], transdata[i]) {
		// 	t.Errorf("expected Transaction struct is not equal to actual")
		// }

	}
}
