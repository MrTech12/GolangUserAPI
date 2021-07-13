# GolangUserAPI

## Overview
An API application, writting in Golang. The API can **retrieve**, **add** & **delete** users.

The application does not communicate with a database. Users are stored in-memory. There are 3 test users in place to experiment with.

The application listens to port 8085, as configured by the following line
```golang
log.Fatal(http.ListenAndServe(":8085", muxRouter))
```

## Endpoints
The following endpoints are available

* GET `/users`
    * Retrieves all users.
    * Does not require any parameter or body data.
    * The data that is returned has the following format:

| name      | datatype  |
| ----------|:---------:|
| ID        | string    |
| mail      | string    |
| phone     | int       |
| password  | string    |

An example of the data:

![example of the data](https://i.imgur.com/QyzE1S1.png)

<br/>

* POST `/users`
    * Creates a new user.
    * Requires body data in the form of JSON. The data needs to be in the following format:

| name      | datatype  |
| ----------|:---------:|
| mail      | string    |
| phone     | int       |
| password  | string    |

<br/>

An example of body data:

![example of body data](https://i.imgur.com/DedmLx3.png)

<br/>

* GET `/users/{id}`
    * Retrieves a single user, based on the given ID.
    * Requires an ID value at the {id} placeholder. An example: `/users/reuRUE473`
* GET `users/mail/{mail}`
    * Retrieves a single user, based on the given mailadress.
    * Requires a mailadress at the {mail} placeholder. An example: `/users/mail/dvorak@gmail.com`

* DELETE `users/{id}`
    * Deletes a user, based on the given ID.
    * Requires an ID value at the {id} placeholder. An example: `/users/reuRUE473`

## Project structure

`main.go`
* Contains API logic and is the starting point of the application.

`UserRepository.go`
* Contains logic to manipulate user objects.

`User.go`
* Contains the structure of a user object.

`go.mod`& `go.sum`
* Contains information about used packages. 