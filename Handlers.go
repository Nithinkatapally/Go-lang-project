package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"strconv"
)


func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status OK", html.EscapeString(r.URL.Path))
}

func AllContactsGet(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(contacts); err != nil {
		panic(err)
	}
}

func GetContact(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)
	contactId, err := strconv.Atoi(vars["contactId"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	sContact:= findContact(contactId)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(sContact); err != nil {
		panic(err)
	}


}


func CreateContact(w http.ResponseWriter, r *http.Request)  {
	var contact Contact
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &contact); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := createContact(contact)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
	
}

func UpdateContact(w http.ResponseWriter, r *http.Request)  {
	var contact Contact
	vars := mux.Vars(r)
	contactId,err := strconv.Atoi(vars["contactId"])
	if err != nil {
		panic( "Invalid product ID")
		return
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &contact); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}


	t := updateContact(contactId,contact)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}

}

func DeleteContact(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	contactId, err := strconv.Atoi(vars["contactId"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	t:=deleteContact(contactId)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}

}


