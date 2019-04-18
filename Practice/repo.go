package main

var currentId int

var users Users

func init() {
	RepoCreateUser(User{Name: "Test"})
}

func RepoGetUser(id int) User {
	for _, user := range users {
		if user.Id == id {
			return user
		}
	}
	// return empty Todo if not found
	return User{}
}

func RepoCreateUser(user User) User {
	currentId += 1
	user.Id = currentId
	users = append(users, user)
	return user
}
