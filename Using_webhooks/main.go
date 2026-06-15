package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type payReq struct {
	AmountPaid int `json:"amountPaid"`
	AmountLeft int `json:"balance"`
	StatusCode int `json:"statuscode"`
}

func main() {
	// client := &http.Client{}

	r := mux.NewRouter()

	// this is webhook
	r.HandleFunc("/api/v1/paymentstatus", paymentStatus).Methods("POST")
	fmt.Println("Starting the server at 4000 port ")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func paymentStatus(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please fill the required field")
		return
	}

	var pay payReq

	json.NewDecoder(r.Body).Decode(&pay)

	a := pay.AmountLeft
	b:= pay.AmountPaid
	c:= pay.StatusCode

	if(c == 402){
		fmt.Println("Payment required!!!")
		json.NewEncoder(w).Encode("notified to client server")
		return;
	}

	fmt.Printf("%v\n", c);
	

	fmt.Printf("Payment Done and access granted\nPaid = %v \nBalance = %v\n", b,a);
	
	json.NewEncoder(w).Encode("notified to client server")

	return
}
