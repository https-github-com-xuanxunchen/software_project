package main

import (
	"fmt"
	"testing"
)

func TestServer(t *testing.T) {
	var server Service

	fmt.Println("--------Test Login start--------")
	loginArgs := LoginArgs{ //测试Login
		Username: "username1",
		Password: "password1",
	}
	loginResult, err := server.Login(&loginArgs)
	if err != nil {
		panic(err)
	}
	fmt.Println(loginResult.GetToken())
	fmt.Println(loginResult.GetErr().GetCode())
	fmt.Println(loginResult.GetErr().GetMsg())
	fmt.Println("--------Test Login end--------")

	fmt.Println("--------Test GetUserInfo start--------")
	token := Token{ //测试GetUserInfo
		Token: "2",
	}
	getUserInfoResult, err := server.GetUserInfo(&token)
	if err != nil {
		panic(err)
	}
	fmt.Println(getUserInfoResult.GetUserInfo().GetId())
	fmt.Println(getUserInfoResult.GetUserInfo().GetUsername())
	fmt.Println(getUserInfoResult.GetUserInfo().GetPassword())
	fmt.Println(getUserInfoResult.GetErr().GetCode())
	fmt.Println(getUserInfoResult.GetErr().GetMsg())
	fmt.Println("--------Test GetUserInfo end--------")

	fmt.Println("--------Test GetOthersInfoArgs start--------")
	getOthersInfoArgs := GetOthersInfoArgs{ //测试GetOthersInfo
		Id: "2",
	}
	getOthersInfoResult, err := server.GetOthersInfo(&getOthersInfoArgs)
	if err != nil {
		panic(err)
	}
	fmt.Println(getOthersInfoResult.GetUserInfo().GetId())
	fmt.Println(getOthersInfoResult.GetUserInfo().GetUsername())
	fmt.Println(getOthersInfoResult.GetUserInfo().GetPassword())
	fmt.Println(getOthersInfoResult.GetErr().GetCode())
	fmt.Println(getOthersInfoResult.GetErr().GetMsg())
	fmt.Println("--------Test GetOthersInfoArgs end--------")

	fmt.Println("--------Test SetUserInfo start--------")
	setUserInfoArgs := SetUserInfoArgs{
		Token:    "2",
		Username: "User2",
		Password: "Password2",
		Head:     "avatar/head1.png",
	}
	setUserInfoResult, err := server.SetUserInfo(&setUserInfoArgs)
	if err != nil {
		panic(err)
	}
	fmt.Println(setUserInfoResult.GetErr().GetCode())
	fmt.Println(setUserInfoResult.GetErr().GetMsg())
	fmt.Println("--------Test SetUserInfo end--------")

	fmt.Println("--------Test Logout start--------")
	logoutArgs := LogoutArgs{ //测试Logout
		Token: "2",
	}
	logoutResult, err := server.Logout(&logoutArgs)
	if err != nil {
		panic(err)
	}
	fmt.Println(logoutResult.GetErr().GetCode())
	fmt.Println(logoutResult.GetErr().GetMsg())
	fmt.Println("--------Test Logout end--------")
}
