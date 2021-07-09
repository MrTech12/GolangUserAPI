package main

// The append method creates a copy of the slice instead of directly adding a new element.
// As a result, the copy with the new element will be returned.
func AddUser(newUser User, currentUsers []User) []User {
	currentUsers = append(currentUsers, newUser)
	return currentUsers
}

func LookupUserByMail(mail string, currentUsers []User) User {
	var matchingIndex int
	indexFound := false

	for index, value := range currentUsers {
		if value.Mail == mail {
			matchingIndex = index
			indexFound = true
		}
	}

	if indexFound {
		return currentUsers[matchingIndex]
	} else {
		return User{}
	}
}

func LookupUserByID(ID int, currentUsers []User) User {
	var matchingIndex int
	indexFound := false

	for index, value := range currentUsers {
		if value.ID == ID {
			matchingIndex = index
			indexFound = true
		}
	}

	if indexFound {
		return currentUsers[matchingIndex]
	} else {
		return User{}
	}
}

func RemoveUser(ID int, currentUsers []User) ([]User, bool) {
	var matchingIndex int
	indexFound := false

	for index, value := range currentUsers {
		if value.ID == ID {
			matchingIndex = index
			indexFound = true
		}
	}

	if indexFound {
		currentUsers = append(currentUsers[:matchingIndex], currentUsers[matchingIndex+1:]...)
		return currentUsers, true
	} else {
		return currentUsers, false
	}
}

