package main

var users []string

func main() {
	AddUser("Alice")
	AddUser("Bob")
}

func AddUser(s string) {
	users = append(users, s)
}
