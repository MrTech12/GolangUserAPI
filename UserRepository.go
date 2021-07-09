package main

func AddUser(newUser User, currentUsers []User) {
	currentUsers = append(currentUsers, newUser)
}

//TODO: Find out how to only have 1 lookup method
func LookupUserByMail(mail string, currentUsers []User) User {
	var matchingIndex int

	for index, value := range currentUsers {
		if value.mail == mail {
			matchingIndex = index
		} else {
			return User{}
		}
	}
	return currentUsers[matchingIndex]
}

func LookupUserByID(ID int, currentUsers []User) User {
	var matchingIndex int

	for index, value := range currentUsers {
		if value.ID == ID {
			matchingIndex = index
		} else {
			return User{}
		}
	}
	return currentUsers[matchingIndex]
}

func DeleteUser(ID int, currentUsers []User) int {
	var matchingIndex int

	for index, value := range currentUsers {
		if value.ID == ID {
			matchingIndex = index
		} else {
			return -1
		}
	}
	currentUsers = append(currentUsers[:matchingIndex], currentUsers[matchingIndex+1:]...)
	return 0
}
