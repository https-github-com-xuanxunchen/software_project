package main

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

type Entry struct {
	Token    string
	Id       string
	Username string
	Password string
	Url      string
	Mark     string // 0 for logout, 1 for login
}

type Service struct {
	sync.RWMutex

	filename string   // For read
	tail     *os.File // For write
}

func (s *Service) Login(in *LoginArgs) (*LoginResult, error) {

	s.Lock()
	defer s.Unlock()

	if in.GetState() == 0 { //
		if err := checkInvalid(in); err != nil { //先检查登录参数是否合法（有无用户名和密码）
			return &LoginResult{ //登录参数不合法
				Err: err,
			}, nil
		}

		var entry *Entry
		if entry = s.noLockGetUserInfoByUsername(in.Username); entry != nil { //以username为查找关键字，在表中查询，entry != nil说明找到username相同的记录，注意这里是把表扫一遍，所以取出来的是最后一条username相同的记录
			if entry.Password == in.Password { //找到的记录的password匹配得上
				entry.Mark = "1" //该用户登陆成功，Mark置为1
			} else { //找到的记录的password匹配不上
				return &LoginResult{
					State: 0,
					Err: &Err{
						Code: 2,
						Msg:  "Incorrect password",
					},
				}, nil
			}
		} else { //以username为查找关键字，在表中查询不到记录，说明是新用户，注册
			return &LoginResult{
				State:      1,
				VerifyCode: "verify code",
			}, nil
		}

	}
	if in.GetState() == 1 {
		if in.GetVerifyCode() == "verify code" { //验证码正确
			id := randStringRunes(32) //id随机生成

			entry := &Entry{
				Token:    id, //Token通过id生成，这里设成和id一样
				Id:       id,
				Username: in.Username,
				Password: in.Password,
				Url:      "", //头像暂时为空
				Mark:     "1",
			}

			if err := s.noLockWriteEntry(entry); err != nil { //在表中追加一项
				return nil, err
			}

			return &LoginResult{
				Token: entry.Token,
				State: 2,
			}, nil
		}
		return &LoginResult{
			State: 0,
			Err: &Err{
				Code: 7,
				Msg:  "Verify code wrong",
			},
		}, nil

	}
	return &LoginResult{
		State: in.GetState(),
		Err: &Err{
			Code: 8,
			Msg:  "Login state wrong",
		},
	}, nil
}

func (s *Service) GetUserInfo(in *Token) (*GetUserInfoResult, error) { //获取自己的userInfo
	s.RLock()
	defer s.RUnlock()

	if entry := s.noLockGetUserInfoByToken(in.Token); entry != nil { //以Token为查找关键字，在表中查询记录，entry != nil 说明查询到了匹配上这个Token的记录
		if entry.Mark == "1" { //该用户已经登录，才能获取到信息
			return &GetUserInfoResult{
				UserInfo: &UserInfo{
					Id:       entry.Id,
					Username: entry.Username,
					Password: "",
					Head:     entry.Url,
				},
			}, nil
		}
		return &GetUserInfoResult{ //该用户未登录，即使Token匹配得上也查询不到记录
			Err: &Err{
				Code: 5,
				Msg:  "Not log in",
			},
		}, nil
	}

	return &GetUserInfoResult{ //Token匹配不上
		Err: &Err{
			Code: 3,
			Msg:  "User not found",
		},
	}, nil
}

func (s *Service) SetUserInfo(in *SetUserInfoArgs) (*SetUserInfoResult, error) { //改动用户信息
	s.Lock()
	defer s.Unlock()

	if entry := s.noLockGetUserInfoByToken(in.Token); entry != nil { //先要根据Token查找到相应的记录，说明有这个用户，才能改动
		if len(in.Username) != 0 { //Username改动了
			entry.Username = in.Username
		}
		if len(in.Password) != 0 { //Password改动了
			entry.Password = in.Password
		}
		if len(in.Head) != 0 { //头像改动了
			entry.Url = in.Head
		}

		if err := s.noLockWriteEntry(entry); err != nil { //追加
			return nil, err
		}

		return &SetUserInfoResult{}, nil
	}

	return &SetUserInfoResult{ //Token匹配不上
		Err: &Err{
			Code: 4,
			Msg:  "Incorrect token",
		},
	}, nil
}

func (s *Service) GetOthersInfo(in *GetOthersInfoArgs) (*GetOthersInfoResult, error) {
	s.RLock()
	defer s.RUnlock()
	if entry := s.noLockGetUserInfoById(in.Id); entry != nil { //根据Id查询记录，并且找到了

		return &GetOthersInfoResult{
			UserInfo: &UserInfo{
				Id:       entry.Id,
				Username: entry.Username,
				Head:     entry.Url,
			},
		}, nil
	}

	return &GetOthersInfoResult{ //没有找到匹配这个Id的记录
		Err: &Err{
			Code: 6,
			Msg:  "Incorrect Id",
		},
	}, nil
}

func (s *Service) Logout(in *LogoutArgs) (*LogoutResult, error) {
	s.Lock()
	defer s.Unlock()

	if entry := s.noLockGetUserInfoByToken(in.Token); entry != nil { //根据Token查询记录，并且找到了
		entry.Mark = "0" //退出，Mark置为0

		if err := s.noLockWriteEntry(entry); err != nil { //追加
			return nil, err
		}

		return &LogoutResult{}, nil
	}

	return &LogoutResult{ //没有找到匹配Token的记录
		Err: &Err{
			Code: 4,
			Msg:  "Incorrect token",
		},
	}, nil
}

func (s *Service) noLockGetUserInfoByUsername(username string) *Entry {
	return s.noLockGetUserInfo(func(entry *Entry) bool {
		return entry.Username == username
	})
}

func (s *Service) noLockGetUserInfoByToken(token string) *Entry {
	return s.noLockGetUserInfo(func(entry *Entry) bool {
		return entry.Token == token
	})
}

func (s *Service) noLockGetUserInfoById(Id string) *Entry {
	return s.noLockGetUserInfo(func(entry *Entry) bool {
		return entry.Id == Id
	})
}

func (s *Service) noLockGetUserInfo(pred func(*Entry) bool) *Entry {
	file, err := os.Open(s.filename)
	if err != nil {
		return nil
	}
	reader := bufio.NewReader(file)
	var entry *Entry

	line, err := reader.ReadString('\n')
	for err == nil {
		if len(line) == 0 {
			break
		}

		nEntry := decode(line[:len(line)-1])
		if pred(nEntry) {
			entry = nEntry
		}

		line, err = reader.ReadString('\n')
	}

	return entry
}

func (s *Service) noLockWriteEntry(entry *Entry) error { //在表中追加一项
	if _, err := s.tail.WriteString(encode(entry)); err != nil {
		return err
	}
	return nil
}

func decode(line string) *Entry { //解码，把字符串解析成entry
	strs := strings.Split(line, " ")
	return &Entry{
		Token:    strs[0],
		Id:       strs[1],
		Username: strs[2],
		Password: strs[3],
		Url:      strs[4],
		Mark:     strs[5],
	}
}

func checkInvalid(args *LoginArgs) *Err { //检查登录参数是否合法
	if len(args.Username) == 0 { //没有用户名
		return &Err{
			Code: 0,
			Msg:  "Empty username",
		}
	} else if len(args.Password) == 0 { //没有密码
		return &Err{
			Code: 1,
			Msg:  "Empty password",
		}
	}
	return nil
}

func encode(entry *Entry) string {
	return entry.Token + " " + entry.Id + " " + entry.Username + " " + entry.Password + " " + entry.Url + " " + entry.Mark + "\n"
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//func main() {
//	file, err := ioutil.TempFile("service", "")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer os.Remove(file.Name())
//}
