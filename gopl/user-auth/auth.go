package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

var userLoginDetails = map[string]string{}

func newUser() {
	var username, password string

	// get user name
	fmt.Println("Enter new user name:")
	fmt.Scanln(&username)

	// get user password
	fmt.Println("Enter password:")
	fmt.Scanln(&password)

	shaUserName := sha256.New().Sum([]byte(username))
	shaPassword := sha256.New().Sum([]byte(password))
	userLoginDetails[string(shaUserName)] = string(shaPassword)
}

func checkLogin() {
	var username, password string

	// get user name
	fmt.Println("Enter existing user name:")
	fmt.Scanln(&username)

	// get user password
	fmt.Println("Enter password:")
	fmt.Scanln(&password)

	enterUserName := sha256.New().Sum([]byte(username))
	enterPassword := sha256.New().Sum([]byte(password))

	// check user exist
	if v, ok := userLoginDetails[string(enterUserName)]; ok {
		if v == string(enterPassword) {
			fmt.Println("Login Success")
		} else {
			fmt.Println("Login failed")
		}
	} else {
		fmt.Println("User not exist")
	}
}

func main() {

	// store user login details
	n := "N"
	for true {
		fmt.Println("New user (y/N)?")
		fmt.Scanln(&n)
		if strings.Compare(n, "y") == 0 {
			newUser()
		}

		// check user login
		checkLogin()
	}
}
