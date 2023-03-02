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

func init() {
	_, err := database.Instance.Exec(`CREATE TABLE IF NOT EXISTS user (
    	id INT NOT NULL AUTO_INCREMENT,
    	email VARCHAR(20),
   		username VARCHAR(20),
    	password VARCHAR(20),
    	creation_time DATETIME,
    	PRIMARY KEY(id)
	)`)
	if err != nil {
		panic(err)
	}
}

func CreateUser(user *User) error {
	creationTime := time.Now()

	res, err := database.Instance.Exec("INSERT INTO user (email, username, password, creation_time) VALUES (?, ?, ?, ?)", user.Email, user.Username, user.Password, creationTime)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(id)
	user.CreationTime = creationTime

	return nil
}

func FetchUsers() ([]User, error) {
	users := make([]User, 0)

	rows, err := database.Instance.Query("SELECT id, email, username, creation_time FROM user")
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
	row := database.Instance.QueryRow("SELECT id, email, username, creation_time FROM user WHERE id = ?", id)

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
	_, err := database.Instance.Exec("UPDATE user SET email = ?, username = ?, password = ? WHERE id = ?", user.Email, user.Username, user.Password, user.ID)
	return err
}

func DeleteUser(id int) error {
	_, err := database.Instance.Exec("DELETE FROM user WHERE id = ?", id)
	return err
}
