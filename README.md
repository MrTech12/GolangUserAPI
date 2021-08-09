# GolangUserAPI

## Overview
An API application, writting in Golang. The API can **retrieve**, **add** & **delete** users.

The application does not communicate with a database. Users are stored in-memory. There are 3 test users in place to experiment with.

The application listens to port 8085, as configured by the following line
```golang
log.Fatal(http.ListenAndServe(":8085", muxRouter))
```

The file **API Documentation.md** contains the endpoints, expected request and return data of the application.

The application is build and runs on **go 1.16.5**.

To run the application:
1. Open a terminal window in the working directory. (Directory where the project is stored on the local system.)
2. run the following command: `go run .`

## Project structure

`main.go`
* Contains API logic and is the starting point of the application.

`UserRepository.go`
* Contains logic to manipulate user objects.

`User.go`
* Contains the structure of a user object.

`go.mod`& `go.sum`
* Contains information about used packages. 