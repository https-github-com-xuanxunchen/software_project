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
	Mark     string
	//Url      string
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
			//Url:      strs[4],
			Mark: strs[4],
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
	file1.WriteString("\n" + strconv.Itoa(cnt) + " " + strconv.Itoa(cnt) + " " + in.GetUsername() + " " + in.GetPassword() + " " + "1\n")

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

	err1 := Err{
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

	//改用户信息（用户名和密码），文件里需要改动，未完成
	/*file1, err := os.OpenFile("table1.txt", os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file1)
	var cnt = 0
	cntChar := 0
	for scanner.Scan() {
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
		cntChar += len(strs)
		if entry.Token == in.Token() {

		}

	}*/

	empty := Empty{}

	err := Err{
		Code: 0,
		Msg:  "no error",
	}

	setUserInfoResult := SetUserInfoResult{
		Empty: &empty,
		Err:   &err,
	}

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
			file1.WriteString("0")

			err1 := Err{
				Code: 0,
				Msg:  "no error",
			}
			LogoutResult := LogoutResult{
				Err: &err1,
			}
			return &LogoutResult, nil
		}
		cntChar += 1
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

	logoutArgs := LogoutArgs{ //测试Logout
		Token: "2",
	}
	logoutResult, err := server.Logout(ctx, &logoutArgs)
	if err != nil {
		panic(err)
	}
	fmt.Println(logoutResult.GetErr().GetCode())
	fmt.Println(logoutResult.GetErr().GetMsg())

}
