package models

import "errors"

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

var users = make(map[int]User)

func SetDefaultUser() {
	user := User{Id: 1, UserName: "Abraham", PassWord: "Popodepollito"}
	users[user.Id] = user
}

func GetUser(id int) (User, error) {
	if user, ok := users[id]; ok {
		return user, nil
	}
	return User{}, errors.New("El usuario no se encontro")
}
