package main

import (
	"fmt"
)

type User struct {
	Id   string
	Name string
	Num  int
}

func (user *User) Save() error {
	// simple validation
	if user.Num < 0 || user.Num > 3 {
		return fmt.Errorf("invalid num")
	}

	_, err := db.Exec("UPDATE users SET num = ? WHERE id = ?", user.Num, user.Id)
	if err != nil {
		return err
	}

	return nil
}

func GetUser(id string) (*User, error) {
	var user User
	rows, err := db.Query("SELECT id, name, num FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Num)
		break
	}
	return &user, nil
}
