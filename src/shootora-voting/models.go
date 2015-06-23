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

func GetSummary() ([]float64, error) {
	res := []float64{0, 0, 0, 0}
	rows, err := db.Query("SELECT num FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var num int
	count := 0.0
	for rows.Next() {
		rows.Scan(&num)
		res[num] += 1.0
		count += 1.0
	}

	for i, v := range res {
		res[i] = v / count * 100.0
	}

	return res, nil
}
