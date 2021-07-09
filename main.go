package main

import (
	"fmt"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	//TODO: create method to add a user, with Regex verification.
	//TODO: create method to lookup an user, via mail.
	//TODO: create method to lookup an user, via ID.
	//TODO: create method to delete an user, via ID.

	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("This is text")
}

func main() {
	fmt.Println("Hello World!")

	Users = []User{
		{ID: 0, mail: "test1@gmail.com", phone: 12345, password: "q"},
		{ID: 1, mail: "test2@gmail.com", phone: 6789, password: "w"},
		{ID: 2, mail: "test3@gmail.com", phone: 9999, password: "s"}}

	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8085", nil))
}
