package greeter_server

import (
	"context"
	"helloworld/helloworld"
)

type UserInfo struct {
	Username helloworld.Username;
	Password helloworld.Password;
}

var UserInfos []UserInfo

type server struct {
	helloworld.UnimplementedGreeterServer
}

func (s *server) Login(ctx context.Context, in *helloworld.LoginArgs) (*helloworld.LoginResult, error) {
	for i := range UserInfos {    //遍历已经注册了的账号，比较LoginArgs的Username和Password是否和UserInfos[i]的Username和Password相等
		if in.Username == &UserInfos[i].Username && helloworld.LoginArgs.GetPassword() == &UserInfos[i].Password {
				//找到这个用户名，且密码正确
				//取地址操作可能有问题，但是*报错，还有括号为什么会有有警告
			token := helloworld.Token{    //生成token
				Token:                "all the tokens are same",
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			}

			err := helloworld.Err{    //生成err（没有错误，错误码为0）
				Code:                 0,
				Msg:                  "no error",
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			}

			loginResult := helloworld.LoginResult{    //把token和err写到loginResult
				Token:                &token,
				Err:                  &err,
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			}

			return &loginResult, nil    //注意返回的第二个参数不是err，返回的第二个参数是表示grpc调用有没有问题的，而err是表示该用户是否存在的，就算该用户不存在，grpc调用也是正常的
		}
	}

	//不存在这个用户或者密码错误
	err := helloworld.Err{
		Code:                 -1,
		Msg:                  "there is an error: wrong username or password.",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	loginResult := helloworld.LoginResult{
		Token:                nil,    //无法生成token
		Err:                  &err,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	return &loginResult, nil
}

func (s *server) Signup(ctx context.Context, in *helloworld.SignupArgs) (*helloworld.SignupResult, error) {
	for i := range UserInfos {
		if in.Username == &UserInfos[i].Username {
				//有可能已经存在相同的用户名，也就是说这个用户名已经被注册过
				//和Login一样的问题
			token := helloworld.Token{
				Token:                in.Username.Username,
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			}

			err := helloworld.Err{
				Code:                 -1,
				Msg:                  "there is an error: Username already exists.",
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			}

			signupResult := helloworld.SignupResult{
				Token:                &token,
				Err:                  &err,
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			}

			return &signupResult, nil
		}
	}

	//该用户名还没有注册过
	token := helloworld.Token{
		Token:                in.Username.Username,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	err := helloworld.Err{
		Code:                 0,
		Msg:                  "no error",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	signupResult := helloworld.SignupResult{
		Token:                &token,
		Err:                  &err,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	return &signupResult, nil
}

func (s *server) GetUserInfo(ctx context.Context, in *helloworld.Token) (*helloworld.GetUserInfoResult, error){
	userInfo := helloworld.UserInfo{
		Id:                   "random ID",
		Nickname:             "random Nickname",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	err := helloworld.Err{
		Code:                 0,
		Msg:                  "",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	getUserInfoResult := helloworld.GetUserInfoResult{
		UserInfo:             &userInfo,
		Err:                  &err,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	return &getUserInfoResult, nil
}

func (s *server) SetUserInfo(ctx context.Context, in *helloworld.SetUserInfoArgs) (*helloworld.SetUserInfoResult, error){
	empty := helloworld.Empty{
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	err := helloworld.Err{
		Code:                 0,
		Msg:                  "no error",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	setUserInfoResult := helloworld.SetUserInfoResult{
		Empty:                &empty,
		Err:                  &err,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	return &setUserInfoResult, nil
}

func (s *server) GetOthersInfo(ctx context.Context, in *helloworld.GetOthersInfoArgs) (*helloworld.GetOthersInfoResult, error) {
	userInfo := helloworld.UserInfo{
		Id:                   "random ID",
		Nickname:             "random Nickname",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	err := helloworld.Err{
		Code:                 0,
		Msg:                  "no error",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	getOthersInfoResult := helloworld.GetOthersInfoResult{
		UserInfo:             &userInfo,
		Err:                  &err,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	return &getOthersInfoResult, nil
}