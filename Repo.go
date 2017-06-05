package main

import (
	"fmt"
)

var currentId int

var contacts Contacts


func init() {
	createContact(Contact{ Id:1,FirstName: "Carrick", LastName: "Michael", Email: "michael@gmail.com", PhoneNumber:"1234664"})
	createContact(Contact{ Id:2,FirstName: "Edvin", LastName: "Vandersar", Email: "edwin@gmail.com", PhoneNumber:"1234664"},)

}

func findContact(id int) Contact {
	for _, t := range contacts {
		if t.Id == id {
			return t
		}
	}
	// return empty Contact if not found
	return Contact{}
}

func createContact(t Contact) Contact {
	currentId += 1
	t.Id = currentId
	contacts = append(contacts, t)
	return t
}


func deleteContact(id int) error {
	for i, t := range contacts {
		if t.Id == id {
			contacts = append(contacts[:i], contacts[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Contact with id of %d to delete", id)
}

func updateContact(id int, t Contact) Contact  {
	oldContact:= findContact(id)
	if (Contact{}) == oldContact  {
		fmt.Println("is zero value")
	}

	t.Id = id
	deleteContact(id)
	contacts = append(contacts,t)

	return  t;

}
