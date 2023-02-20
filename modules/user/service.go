package user

import (
	"errors"
	"github.com/AstroFireWasTaken/ChatAPI/database"
	"strconv"
)

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

func UpdateUser(user *User) error {
	_, err := database.Instance.Exec("UPDATE users SET email = ?, username = ?, password = ? WHERE id = ?", user.Email, user.Username, user.Password, user.ID)
	return err
}

func DeleteUser(id int) error {
	res, err := database.Instance.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count < 1 {
		return errors.New("No user found with id " + strconv.Itoa(int(id)))
	}

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
	rows, err := database.Instance.Query("SELECT id, email, username, creation_time FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, nil
	}

	user := &User{}
	err = rows.Scan(&user.ID, &user.Email, &user.Username, &user.CreationTime)
	if err != nil {
		return nil, err
	}

	return user, nil
}
