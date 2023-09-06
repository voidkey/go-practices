package main

import (
	"IM/client/processes"
	"fmt"
)

var (
	key  int
	loop = true

	userId   int
	userPwd  string
	userName string
)

func main() {
	for {
		fmt.Println("------------------Instant Messaging------------------")
		fmt.Println("\t\t\t1.Log in")
		fmt.Println("\t\t\t2.Sign up")
		fmt.Println("\t\t\t3.Exit")
		fmt.Println("Choose(1~3):")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			loginMenu()
			loop = false
		case 2:
			signupMenu()
			loop = false
		case 3:
			fmt.Println("Bye bye!")
			loop = false
		default:
			fmt.Println("Unexcepted Option, Please input correctly!")
		}
		if !loop {
			break
		}
	}
}

func loginMenu() {
	fmt.Println("------------------Login Menu------------------")
	fmt.Println("Please input userID:")
	fmt.Scanf("%d\n", &userId)
	fmt.Println("Please input password:")
	fmt.Scanf("%s\n", &userPwd)
	up := &processes.UserProcess{}
	err := up.Login(userId, userPwd)
	if err != nil {
		fmt.Printf("Log in failed!\nError:%v\n", err)
	}
	fmt.Println("------------------Login Menu------------------")
}

func signupMenu() {
	fmt.Println("------------------Signup Menu------------------")
	fmt.Println("Please input userID:")
	fmt.Scanf("%d\n", &userId)
	fmt.Println("Please input password:")
	fmt.Scanf("%s\n", &userPwd)
	fmt.Println("Please input userName:")
	fmt.Scanf("%s\n", &userName)
	up := &processes.UserProcess{}
	err := up.Signup(userId, userPwd, userName)
	if err != nil {
		fmt.Printf("Sign up failed!\nError:%v\n", err)
	}
	fmt.Println("------------------Signup Menu------------------")
}
