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
	if err := checkInvalid(in); err != nil {
		return &LoginResult{
			Err: err,
		}, nil
	}

	s.Lock()
	defer s.Unlock()

	var entry *Entry
	if entry = s.noLockGetUserInfoByUsername(in.Username); entry != nil {
		if entry.Password == in.Password {
			entry.Mark = "1"
		} else {
			return &LoginResult{
				Err: &Err{
					Code: 2,
					Msg:  "Incorrect password",
				},
			}, nil
		}
	} else {
		id := randStringRunes(32)
		entry = &Entry{
			Token:    id,
			Id:       id,
			Username: in.Username,
			Password: in.Password,
			Url:      "",
			Mark:     "1",
		}
	}

	if err := s.noLockWriteEntry(entry); err != nil {
		return nil, err
	}

	return &LoginResult{
		Token: entry.Token,
	}, nil
}

func (s *Service) GetUserInfo(in *Token) (*GetUserInfoResult, error) {
	s.RLock()
	defer s.RUnlock()

	if entry := s.noLockGetUserInfoByToken(in.Token); entry != nil {
		return &GetUserInfoResult{
			UserInfo: &UserInfo{
				Id:       entry.Id,
				Username: entry.Username,
				Password: "",
				Head:     entry.Url,
			},
		}, nil
	}

	return &GetUserInfoResult{
		Err: &Err{
			Code: 3,
			Msg:  "User not found",
		},
	}, nil
}

func (s *Service) SetUserInfo(in *SetUserInfoArgs) (*SetUserInfoResult, error) {
	s.Lock()
	defer s.Unlock()

	if entry := s.noLockGetUserInfoByToken(in.Token); entry != nil {
		if len(in.Username) != 0 {
			entry.Username = in.Username
		}
		if len(in.Password) != 0 {
			entry.Password = in.Password
		}
		if len(in.Head) != 0 {
			entry.Url = in.Head
		}

		if err := s.noLockWriteEntry(entry); err != nil {
			return nil, err
		}

		return &SetUserInfoResult{}, nil
	}

	return &SetUserInfoResult{
		Err: &Err{
			Code: 4,
			Msg:  "Incorrect token",
		},
	}, nil
}

func (s *Service) GetOthersInfo(in *GetOthersInfoArgs) (*GetOthersInfoResult, error) {
	s.RLock()
	defer s.RUnlock()

	return nil, nil
}

func (s *Service) Logout(in *LogoutArgs) (*LogoutResult, error) {
	s.Lock()
	defer s.Unlock()

	if entry := s.noLockGetUserInfoByToken(in.Token); entry != nil {
		entry.Mark = "0"

		if err := s.noLockWriteEntry(entry); err != nil {
			return nil, err
		}

		return &LogoutResult{}, nil
	}

	return &LogoutResult{
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

func (s *Service) noLockWriteEntry(entry *Entry) error {
	if _, err := s.tail.WriteString(encode(entry)); err != nil {
		return err
	}
	return nil
}

func decode(line string) *Entry {
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

func checkInvalid(args *LoginArgs) *Err {
	if len(args.Username) == 0 {
		return &Err{
			Code: 0,
			Msg:  "Empty username",
		}
	} else if len(args.Password) == 0 {
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
