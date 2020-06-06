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

	filename string  // For read
	tail     os.File // For write
}

func (s *Service) Login(in *LoginArgs) (*LoginResult, error) {
	if err := checkInvalid(in); err != nil {
		return &LoginResult{
			Token: "",
			Err:   err,
		}, nil
	}

	s.Lock()
	defer s.Unlock()

	var entry *Entry
	if entry = s.noLockGetUserInfo(in.Username); entry != nil {
		if entry.Password == in.Password {
			entry.Mark = "1"
		} else {
			return &LoginResult{
				Token: "",
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

	if _, err := s.tail.WriteString(encode(entry)); err != nil {
		return nil, err
	}

	return &LoginResult{
		Token: entry.Token,
		Err:   nil,
	}, nil
}

func (s *Service) GetUserInfo(in *Token) (*GetUserInfoResult, error) {
	s.RLock()
	defer s.RUnlock()

	return nil, nil
}

func (s *Service) SetUserInfo(in *SetUserInfoArgs) (*SetUserInfoResult, error) {
	s.Lock()
	defer s.Unlock()

	return nil, nil
}

func (s *Service) GetOthersInfo(in *GetOthersInfoArgs) (*GetOthersInfoResult, error) {
	s.RLock()
	defer s.RUnlock()

	return nil, nil
}

func (s *Service) Logout(in *LogoutArgs) (*LogoutResult, error) {
	s.Lock()
	defer s.Unlock()

	return nil, nil
}

func (s *Service) noLockGetUserInfo(username string) *Entry {
	file, err := os.Open(s.filename)
	if err != nil {
		return nil
	}
	reader := bufio.NewReader(file)
	for line, err := reader.ReadString('\n'); err != nil; line, err = reader.ReadString('\n') {
		entry := decode(line)
		if entry.Username == username {
			return entry
		}
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
