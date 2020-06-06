package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var UserInfos map[string]string

type Entry struct {
	Token    string
	Id       string
	Username string
	Password string
	Url      string
	Mark     string
}

type server2 struct {
	//helloworld.UnimplementedGreeterServer
}

func (s *server2) Login(ctx context.Context, in *LoginArgs) (*LoginResult, error) {
	file1, err := os.OpenFile("table1.txt", os.O_RDWR, os.ModePerm) //打开文件
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file1) //逐行扫描文件
	var cnt = 0
	cntChar := 0
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), " ") //每行的内容提取出来
		//fmt.Println(scanner.Text())
		entry := Entry{
			Token:    strs[0],
			Id:       strs[1],
			Username: strs[2],
			Password: strs[3],
			Url:      strs[4],
			Mark:     strs[5],
		}
		cnt++
		cntChar += len(scanner.Text()) //记录总共经过了多少个字符
		fmt.Println("cntChar = ", cntChar)
		if entry.Username == in.GetUsername() { //用户名和表里的相等
			if entry.Password == in.GetPassword() { //密码和表里的相等，登录成功
				file1.Seek(int64(cntChar-1), 0) //文件指针偏移到这里
				file1.WriteString("1")          //改写登录状态，1表示登录，0表示未登录
				err1 := Err{
					Code: 0,
					Msg:  "no error: login in",
				}
				loginResult := LoginResult{
					Token: entry.Token,
					Err:   &err1,
				}
				file1.Close()
				return &loginResult, nil
			}
			err1 := Err{ //用户名匹配而密码不相等，说明密码错误，登录失败
				Code: 1,
				Msg:  "error: wrong password",
			}
			loginResult := LoginResult{
				Token: "",
				Err:   &err1,
			}
			file1.Close()
			return &loginResult, nil

		}
		cntChar += 1
	}

	//是新的用户，需要在文件追加一行，表明注册
	file1.WriteString("\n" + strconv.Itoa(cnt) + " " + strconv.Itoa(cnt) + " " + in.GetUsername() + " " + in.GetPassword() + " avatar/head1.png " + "1\n")

	err1 := Err{
		Code: 0,
		Msg:  "no error: sign up",
	}
	loginResult := LoginResult{
		Token: strconv.Itoa(cnt),
		Err:   &err1,
	}
	file1.Close()
	return &loginResult, nil

}

func (s *server2) GetUserInfo(ctx context.Context, in *Token) (*GetUserInfoResult, error) {
	file1, err := os.OpenFile("table1.txt", os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file1)
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), " ")
		if in.Token == strs[0] { //拿到个人信息
			userInfo := UserInfo{
				Id:       strs[1],
				Username: strs[2],
				Password: strs[3],
				Head:     strs[4],
			}

			err1 := Err{
				Code: 0,
				Msg:  "no error",
			}

			getUserInfoResult := GetUserInfoResult{
				UserInfo: &userInfo,
				Err:      &err1,
			}
			file1.Close()
			return &getUserInfoResult, nil
		}
	}

	err1 := Err{ //在表中没有找到token
		Code: 1,
		Msg:  "error: no such user",
	}

	getUserInfoResult := GetUserInfoResult{
		UserInfo: nil,
		Err:      &err1,
	}
	file1.Close()
	return &getUserInfoResult, nil
}

func (s *server2) SetUserInfo(ctx context.Context, in *SetUserInfoArgs) (*SetUserInfoResult, error) {

	//改用户信息（用户名和密码），文件里需要改动
	file1, err := os.OpenFile("table1.txt", os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file1)
	var cnt = 0
	cntChar := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		strs := strings.Split(scanner.Text(), " ")

		entry := Entry{
			Token:    strs[0],
			Id:       strs[1],
			Username: strs[2],
			Password: strs[3],
			Url:      strs[4],
			Mark:     strs[5],
		}

		cnt++
		cntChar += len(scanner.Text())

		if in.GetUsername() == entry.Username { //找到存在的用户名
			if in.GetToken() == entry.Token { //就是自己，不是别人
				file1.Seek(int64(cntChar-len(scanner.Text())), 0)
				entry.Password = in.GetPassword()
				entry.Url = in.GetHead()

				file1.Seek(int64(cntChar-len(scanner.Text())), 0)                                                                                         //文件指针偏移到要修改的地方
				file1.WriteString(entry.Token + " " + entry.Id + " " + entry.Username + " " + entry.Password + " " + entry.Url + " " + entry.Mark + "\n") //这里有bug，假设现在要写入的行比原来短，就没法完全覆盖，换行的时候，这一行没有被覆盖掉的内容会覆盖掉下一行

				err1 := Err{
					Code: 0,
					Msg:  "no error",
				}

				serUserInfoResult := SetUserInfoResult{
					Err: &err1,
				}
				return &serUserInfoResult, nil

			}

			//已经存在该用户名，而且不是自己的，不能改
			err1 := Err{
				Code: 1,
				Msg:  "same username exist",
			}

			serUserInfoResult := SetUserInfoResult{
				Err: &err1,
			}
			return &serUserInfoResult, nil
		}

		cntChar++

	}

	//没有存在该用户名，可以改，再遍历一次
	cntChar = 0 //字符计数清0
	file1.Seek(0, 0)
	scanner2 := bufio.NewScanner(file1)
	for scanner2.Scan() {
		fmt.Println(scanner2.Text())
		strs := strings.Split(scanner2.Text(), " ")

		entry := Entry{
			Token:    strs[0],
			Id:       strs[1],
			Username: strs[2],
			Password: strs[3],
			Url:      strs[4],
			Mark:     strs[5],
		}

		cntChar += len(scanner2.Text())
		if entry.Token == in.Token { //存在这个token
			entry.Username = in.GetUsername()
			entry.Password = in.GetPassword()
			entry.Url = in.GetHead()

			file1.Seek(int64(cntChar-len(scanner2.Text())), 0)                                                                                        //文件指针偏移到要修改的地方
			file1.WriteString(entry.Token + " " + entry.Id + " " + entry.Username + " " + entry.Password + " " + entry.Url + " " + entry.Mark + "\n") //这里有bug，假设现在要写入的行比原来短，就没法完全覆盖，换行的时候，这一行没有被覆盖掉的内容会覆盖掉下一行

			err1 := Err{
				Code: 0,
				Msg:  "no error",
			}

			setUserInfoResult := SetUserInfoResult{
				Err: &err1,
			}

			file1.Close()
			return &setUserInfoResult, nil
		}

		cntChar++

	}

	//找不到token对应的用户
	err1 := Err{
		Code: 1,
		Msg:  "error: no such user",
	}

	setUserInfoResult := SetUserInfoResult{
		Err: &err1,
	}

	file1.Close()
	return &setUserInfoResult, nil
}

func (s *server2) GetOthersInfo(ctx context.Context, in *GetOthersInfoArgs) (*GetOthersInfoResult, error) {

	file1, err := os.OpenFile("table1.txt", os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file1)
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), " ")
		if in.Id == strs[1] { //拿到别人信息
			userInfo := UserInfo{
				Id:       in.Id,
				Username: strs[2],
				Password: "", //别人的密码拿不到
				Head:     strs[4],
			}
			err1 := Err{
				Code: 0,
				Msg:  "no error",
			}
			getOthersInfoResult := GetOthersInfoResult{
				UserInfo: &userInfo,
				Err:      &err1,
			}
			return &getOthersInfoResult, nil
		}
	}

	err1 := Err{
		Code: 1,
		Msg:  "error: no such user",
	}

	getOthersInfoResult := GetOthersInfoResult{
		UserInfo: nil,
		Err:      &err1,
	}

	return &getOthersInfoResult, nil
}

func (s *server2) Logout(ctx context.Context, in *LogoutArgs) (*LogoutResult, error) {
	file1, err := os.OpenFile("table1.txt", os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file1)
	cntChar := 0
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), " ")
		cntChar += len(scanner.Text())
		if in.Token == strs[1] {
			//需要在文件里把这一行的mark改了
			file1.Seek(int64(cntChar-1), 0)
			file1.WriteString("0") //改成未登录状态

			err1 := Err{
				Code: 0,
				Msg:  "no error",
			}
			LogoutResult := LogoutResult{
				Err: &err1,
			}
			return &LogoutResult, nil
		}
		cntChar++
	}

	err1 := Err{
		Code: 1,
		Msg:  "error: no such user",
	}
	LogoutResult := LogoutResult{
		Err: &err1,
	}
	return &LogoutResult, nil

}

func main() {
	var server server2

	fmt.Println("--------Test Login start--------")
	loginArgs := LoginArgs{ //测试Login
		Username: "username1",
		Password: "password1",
	}
	var ctx context.Context
	loginResult, err := server.Login(ctx, &loginArgs)
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
	getUserInfoResult, err := server.GetUserInfo(ctx, &token)
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
	getOthersInfoResult, err := server.GetOthersInfo(ctx, &getOthersInfoArgs)
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
		Username: "ser2",
		Password: "assword2",
		Head:     "avatar/head1.png",
	}
	setUserInfoResult, err := server.SetUserInfo(ctx, &setUserInfoArgs)
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
	logoutResult, err := server.Logout(ctx, &logoutArgs)
	if err != nil {
		panic(err)
	}
	fmt.Println(logoutResult.GetErr().GetCode())
	fmt.Println(logoutResult.GetErr().GetMsg())
	fmt.Println("--------Test Logout end--------")
}
