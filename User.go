package main

type User struct {
	ID       string `json:"ID"`
	Mail     string `json:"mail"`
	Phone    int    `json:"phone"`
	Password string `json:"password"`
}

var Users []User
