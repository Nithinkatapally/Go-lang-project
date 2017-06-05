package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
		Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"AllContacts",
		"GET",
		"/contacts",
		AllContactsGet,
	},
	Route{
		"Contacts",
		"GET",
		"/contacts/{contactId}",
		GetContact,
	},
	Route{
		"AddContact",
		"POST",
		"/contacts",
		CreateContact,
	},
	Route{
		"AddContact",
		"PUT",
		"/contacts/{contactId:[0-9]+}",
		UpdateContact,
	},
	Route{
			"DeleteContact",
			"DELETE",
			"/contacts/{contactId:[0-9]+}",
			DeleteContact,
		},

}
