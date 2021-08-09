package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/mail"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	receivedUser, _ := ioutil.ReadAll(r.Body)
	var newUser User
	json.Unmarshal(receivedUser, &newUser) //Converting the JSON request body to an User struct
	validEntry := false

	if len(newUser.Mail) == 0 || len(newUser.Password) == 0 || newUser.Phone == 0 { //Checking if the User struct has the required values.
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Not all field have been entered.")
	} else {
		validEntry = true
	}

	if validEntry { //The User struct has the required values.
		if VerifyMail(newUser.Mail) {
			Users = AddUser(newUser, Users)

			w.WriteHeader(http.StatusCreated)
			fmt.Fprint(w, "User has been created successfully.")
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "A valid mailadress has not been entered.")
		}
	}
}

// Checking if the entered mail, has a valid RFC 5322 format.
func VerifyMail(enteredMail string) bool {
	_, err := mail.ParseAddress(enteredMail)
	if err == nil {
		return true
	} else {
		return false
	}
}

func FindUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	routeValue := mux.Vars(r) //Returning route values.
	receivedID := routeValue["id"]

	tracedUser := LookupUserByID(receivedID, Users)
	if len(tracedUser.ID) != 0 {
		json.NewEncoder(w).Encode(tracedUser)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "There is no user with that ID.")
	}
}

func FindUserByMail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	routeValue := mux.Vars(r) //Returning route values.
	receivedMail := routeValue["mail"]

	if VerifyMail(receivedMail) {
		tracedUser := LookupUserByMail(receivedMail, Users)
		if len(tracedUser.Password) != 0 {
			json.NewEncoder(w).Encode(tracedUser)
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "There is no user with that mailadress.")
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "A valid mailadress has not been entered.")
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	routeValue := mux.Vars(r) //Returning route values.
	receivedID := routeValue["id"]

	removed := false
	Users, removed = RemoveUser(receivedID, Users)
	if removed {
		fmt.Fprint(w, "User successfully deleted.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "There is no user with that ID.")
	}
}

func RequestHandler() {
	muxRouter := mux.NewRouter().StrictSlash(true)

	muxRouter.HandleFunc("/users", GetUsers).Methods("GET")                   //Getting all users
	muxRouter.HandleFunc("/users", CreateUser).Methods("POST")                //Creating a new user
	muxRouter.HandleFunc("/users/{id}", FindUserByID).Methods("GET")          //Getting a specific user by id
	muxRouter.HandleFunc("/users/mail/{mail}", FindUserByMail).Methods("GET") //Getting a specific user by mail
	muxRouter.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")         //Deleting a specific user by id

	log.Fatal(http.ListenAndServe(":8085", muxRouter))
}

func main() {
	fmt.Println("Application started on port 8085.")

	Users = []User{
		{ID: "dwRQAc68PhHQh4BUnrNsoS", Mail: "qwerty@gmail.com", Phone: 1234567, Password: "qaerrer#@##"},
		{ID: "dwEREWE234EEW4BUnrNsEW", Mail: "azerty@gmail.com", Phone: 8905323, Password: "URuttu3848R"},
		{ID: "RUUEuewueueEUEU233UEUE", Mail: "dvorak@gmail.com", Phone: 9999999, Password: "EEEU88838#4"}}
	RequestHandler()
}
