package model

import (
	"fmt"
)

type User struct {
	ID  string `json:"id"`
	PWD string `json:"pwd"`
}

func (user *User) Create() error {
	err := Db.Set(Ctx, user.ID, user.PWD, 0).Err()
	if err != nil {
		fmt.Println("Couldn't set the pair.: ", err)
		panic(err)
	}
	return err
}

// TODO: 更新メソッドの追加
// func (user *User) Update() error {

// }

// TODO: 削除メソッドの追加
// func (user *User) Delete() error {

// }

func GetUserByID(id string) (User, error) {
	var user User
	pwd, err := Db.Get(Ctx, id).Result()
	if err != nil {
		fmt.Println("Couldn't get the value.: ", err)
		panic(err)
	}
	user.ID = id
	user.PWD = pwd
	return user, err
}

// ログインリクエストのPWDとDBに保存されたPWDをを照合
func Authenticate(loginReq User) bool {
	user, err := GetUserByID(loginReq.ID)
	if err != nil {
		fmt.Println("Couldn't get the user.: ", err)
		panic(err)
	}

	if user.PWD == "" {
		fmt.Println("No such user.")
		return false
	}
	if loginReq.PWD != user.PWD {
		return false
	}
	return true
}
