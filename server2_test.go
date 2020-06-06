package main

import (
	"log"
	"os"
	"testing"
)

func TestServer(t *testing.T) {
	file, err := os.OpenFile("tmp.txt", os.O_APPEND|os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove("tmp.txt")

	service := Service{
		filename: file.Name(),
		tail:     file,
	}

	loginArgs := LoginArgs{
		Username: "username1",
		Password: "password1",
	}
	loginResult, err := service.Login(&loginArgs)
	if err != nil {
		t.Fatal(err)
	}
	if loginResult.GetErr() != nil {
		t.Fatal("login failed")
	}

	token := Token{
		Token: loginResult.GetToken(),
	}
	getUserInfoResult, err := service.GetUserInfo(&token)
	if err != nil {
		t.Fatal(err)
	}
	if getUserInfoResult.GetErr() != nil {
		t.Fatal("get user info failed")
	}
	if getUserInfoResult.GetUserInfo().GetUsername() != "username1" {
		t.Fatal("non-match username")
	}

	setUserInfoArgs := SetUserInfoArgs{
		Token:    loginResult.GetToken(),
		Head:     "avatar/head1.png",
	}
	setUserInfoResult, err := service.SetUserInfo(&setUserInfoArgs)
	if err != nil {
		t.Fatal(err)
	}
	if setUserInfoResult.GetErr() != nil {
		t.Fatal("set user info failed")
	}

	getUserInfoResult, err = service.GetUserInfo(&token)
	if err != nil {
		t.Fatal(err)
	}
	if getUserInfoResult.GetErr() != nil {
		t.Fatal("get user info failed")
	}
	if getUserInfoResult.GetUserInfo().GetHead() != "avatar/head1.png" {
		t.Fatal("non-match avatar")
	}

	logoutArgs := LogoutArgs{
		Token: loginResult.GetToken(),
	}
	logoutResult, err := service.Logout(&logoutArgs)
	if err != nil {
		t.Fatal(err)
	}
	if logoutResult.GetErr() != nil {
		t.Fatal("log out failed")
	}
}
