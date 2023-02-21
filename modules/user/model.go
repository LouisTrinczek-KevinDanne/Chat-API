package user

import (
	"database/sql"
	"errors"
	"github.com/LouisTrinczek-KevinDanne/Chat-API/database"
	"time"
)

type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	Password     string    `json:"password,omitempty"`
	CreationTime time.Time `json:"creationTime"`
}

func CreateUser(user *User) error {
	res, err := database.Instance.Exec("INSERT INTO users (email, username, password) VALUES (?, ?, ?)", user.Email, user.Username, user.Password)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(id)

	return nil
}

func FetchUsers() ([]User, error) {
	users := make([]User, 0)

	rows, err := database.Instance.Query("SELECT id, email, username, creation_time FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Email, &user.Username, &user.CreationTime)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func FetchUserById(id int) (*User, error) {
	row := database.Instance.QueryRow("SELECT id, email, username, creation_time FROM users WHERE id = ?", id)

	user := &User{}
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.CreationTime)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func UpdateUser(user *User) error {
	_, err := database.Instance.Exec("UPDATE users SET email = ?, username = ?, password = ? WHERE id = ?", user.Email, user.Username, user.Password, user.ID)
	return err
}

func DeleteUser(id int) error {
	_, err := database.Instance.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
