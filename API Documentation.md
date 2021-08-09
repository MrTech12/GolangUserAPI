## The following endpoints are available:

<br>

# GET `/users`
* Retrieves all users.
* Does not require any parameter or body data.

## Responses
| HTTP status code      | description  |
|:---------------------:|:------------:|
| 200                   | all users    |

<br>

The data that is returned has the following format:

| name      | datatype  |
| ----------|:---------:|
| ID        | string    |
| mail      | string    |
| phone     | int       |
| password  | string    |

<br/>

An example of the data:

![example of data](https://i.imgur.com/QyzE1S1.png)

<br/>

# POST `/users`
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

<br>

## Responses
| HTTP status code      | description                             |
|:---------------------:|:---------------------------------------:|
| 201                   | User has been created successfully.     |
| 400                   | Not all field have been entered.        |
| 400                   | A valid mailadress has not been entered.|

<br>

# GET `/users/{id}`
* Retrieves a single user, based on the given ID.
* Requires an ID value at the {id} placeholder. An example: `/users/reuRUE473`

<br>

## Responses
| HTTP status code      | description                             |
|:---------------------:|:---------------------------------------:|
| 200                   | one user                                |
| 404                   | There is no user with that ID.          |

<br>

# GET `users/mail/{mail}`
* Retrieves a single user, based on the given mailadress.
* Requires a mailadress at the {mail} placeholder. An example: `/users/mail/dvorak@gmail.com`

<br>

## Responses
| HTTP status code      | description                               |
|:---------------------:|:-----------------------------------------:|
| 200                   | one user                                  |
| 400                   | A valid mailadress has not been entered.  |
| 404                   | There is no user with that mailadress.    |

<br>

# DELETE `users/{id}`
* Deletes a user, based on the given ID.
* Requires an ID value at the {id} placeholder. An example: `/users/reuRUE473`

<br>

## Responses
| HTTP status code      | description                               |
|:---------------------:|:-----------------------------------------:|
| 200                   | User successfully deleted.                |
| 404                   | There is no user with that ID.            |

<br>