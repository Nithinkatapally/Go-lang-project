package main

type Contact struct {
	Id 	       int         `json:id`
	FirstName      string      `json:"firstName"`
	LastName       string      `json:"lastName"`
	Email          string	   `json:"email"`
	PhoneNumber    string	   `json:"phoneNumber"`
}

type Contacts []Contact
