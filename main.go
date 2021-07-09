package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("This is text")
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// response, _ := json.Marshal(Users)
	// w.Write(response)
	// fmt.Fprintf(w, "%+v", string(createdUser))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Users)
}

//TODO: Add Regex verification for mail.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	receivedUser, _ := ioutil.ReadAll(r.Body)
	var newUser User
	json.Unmarshal(receivedUser, &newUser) //Convert the JSON request body to an User struct
	Users = AddUser(newUser, Users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "User has been created successfully.")
}

func FindUserByID(w http.ResponseWriter, r *http.Request) {
	routeValue := mux.Vars(r)
	receivedID := routeValue["id"]
	numericID, err := strconv.Atoi(receivedID)
	if err == nil {
		tracedUser := LookupUserByID(numericID, Users)
		if len(tracedUser.Password) != 0 {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(tracedUser)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "There is no user with that ID")
		}
	}
}

func FindUserByMail(w http.ResponseWriter, r *http.Request) {
	routeValue := mux.Vars(r)
	receivedMail := routeValue["mail"]
	tracedUser := LookupUserByMail(receivedMail, Users)
	if len(tracedUser.Password) != 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tracedUser)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "There is no user with that mailadress.")
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	routeValue := mux.Vars(r)
	receivedID := routeValue["id"]
	numericID, err := strconv.Atoi(receivedID)
	if err == nil {
		removed := false
		Users, removed = RemoveUser(numericID, Users)
		if removed {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, "User successfully deleted.")
		} else if !removed {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "There is no user with that ID.")
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "A valid numeric ID has not been send.")
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
		{ID: 0, Mail: "qwerty@gmail.com", Phone: 1234567, Password: "qaerrer#@##"},
		{ID: 1, Mail: "azerty@gmail.com", Phone: 8905323, Password: "URuttu3848R"},
		{ID: 2, Mail: "dvorak@gmail.com", Phone: 9999999, Password: "EEEU88838#4"}}
	RequestHandler()
}
