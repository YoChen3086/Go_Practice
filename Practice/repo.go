package main

var currentId int

var users map[int]User

func init() {
	users = make(map[int]User)
}

func RepoGetUser(id int) User {
	user, ok := users[id]

	if ok {
		return user
	} else {
		return User{}
	}
}

func RepoCreateUser(user User) User {
	currentId += 1
	user.Id = currentId
	users[user.Id] = user
	return user
}

func RepoUpdateUser(id int, name string) User {
	user, ok := users[id]

	if ok {
		user.Name = name
		users[id] = user
		return user
	} else {
		return RepoCreateUser(User{Id: id, Name: name})
	}
}

func RepoDeleteUser(id int) User {
	user, ok := users[id]

	if ok {
		delete(users, id)
		return user
	} else {
		return User{}
	}
}
