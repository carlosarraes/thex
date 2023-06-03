package main

import "fmt"

type User struct {
	firstName string
	lastName  string
	age       int
	email     string
	password  string
}

func main() {
	user := User{
		firstName: "John",
		lastName:  "Doe",
		age:       25,
		email:     "johndoe@gmail.com",
		password:  "123456",
	}

	fmt.Println(user)
	changeName(&user, "Jane", "Doe")
	changePassword(&user, "654321")
	changeEmail(&user, "janedoe@gmail.com")
	changeAge(&user, 26)
	fmt.Println(user)
}

func changeName(u *User, firstName string, lastName string) {
	u.firstName = firstName
	u.lastName = lastName
}

func changePassword(u *User, password string) {
	u.password = password
}

func changeEmail(u *User, email string) {
	u.email = email
}

func changeAge(u *User, age int) {
	u.age = age
}
