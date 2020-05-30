package helloworld

import (
	"context"
)

var UserInfos map[string]string

type server struct {
	//helloworld.UnimplementedGreeterServer
}

func (s *server) Login(ctx context.Context, in *LoginArgs) (*LoginResult, error) {

	if UserInfos[in.Username.GetUsername()] == in.GetPassword().GetPassword() { //在哈希表UserInfos中存在这个用户名和密码
		token := Token{ //Token的字段Token暂时先设置成和用户名相同，都是string
			Token: in.Username.GetUsername(),
		}

		err := Err{ //生成err（没有错误，错误码为0）
			Code: 0,
			Msg:  "no error",
		}

		loginResult := LoginResult{ //把token和err写到loginResult
			Token: &token,
			Err:   &err,
		}

		return &loginResult, nil //注意返回的第二个参数不是err，返回的第二个参数是表示grpc调用有没有问题的，而err是表示该用户是否存在的，就算该用户不存在，grpc调用也是正常的

	}

	//不存在这个用户或者密码错误
	err := Err{
		Code: 1,
		Msg:  "there is an error: wrong username or password.",
	}

	loginResult := LoginResult{
		Token: nil, //无法生成token
		Err:   &err,
	}

	return &loginResult, nil
}

func (s *server) Signup(ctx context.Context, in *SignupArgs) (*SignupResult, error) {

	if _, ok := UserInfos[in.GetUsername().GetUsername()]; ok { //有可能已经存在相同的用户名，也就是说这个用户名已经被注册过

		err := Err{
			Code: 1,
			Msg:  "there is an error: Username already exists.",
		}

		signupResult := SignupResult{
			Token: nil, //因为用户名被注册过，所以拿不到token
			Err:   &err,
		}

		return &signupResult, nil
	}

	//该用户名还没有注册过
	token := Token{
		Token: in.GetUsername().GetUsername(),
	}

	err := Err{
		Code: 0,
		Msg:  "no error",
	}

	signupResult := SignupResult{
		Token: &token,
		Err:   &err,
	}

	return &signupResult, nil
}

func (s *server) GetUserInfo(ctx context.Context, in *Token) (*GetUserInfoResult, error) {
	userInfo := UserInfo{
		Username: in.GetToken(), //Token和Username的字符串相同
		Password: UserInfos[in.GetToken()],
	}

	err := Err{
		Code: 0,
		Msg:  "no error",
	}

	getUserInfoResult := GetUserInfoResult{
		UserInfo: &userInfo,
		Err:      &err,
	}

	return &getUserInfoResult, nil
}

func (s *server) SetUserInfo(ctx context.Context, in *SetUserInfoArgs) (*SetUserInfoResult, error) {

	//改用户信息（用户名和密码），需要改动哈希表UserInfos，还未实现

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

func (s *server) GetOthersInfo(ctx context.Context, in *GetOthersInfoArgs) (*GetOthersInfoResult, error) {
	userInfo := UserInfo{
		Username: "random Username",
		//别人的Password肯定是拿不到
	}

	err := Err{
		Code: 0,
		Msg:  "no error",
	}

	getOthersInfoResult := GetOthersInfoResult{
		UserInfo: &userInfo,
		Err:      &err,
	}

	return &getOthersInfoResult, nil
}

func main() {

}
