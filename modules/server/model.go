package server

import (
	"database/sql"
	"errors"
	"github.com/LouisTrinczek-KevinDanne/Chat-API/database"
	"github.com/LouisTrinczek-KevinDanne/Chat-API/modules/channel"
	"github.com/LouisTrinczek-KevinDanne/Chat-API/modules/user"
	"time"
)

type Server struct {
	ID           int               `json:"id"`
	Name         string            `json:"name"`
	Channels     []channel.Channel `json:"channels"`
	Members      []user.User       `json:"members"`
	CreationTime time.Time         `json:"creationTime"`
}

func init() {
	_, err := database.Instance.Exec(`CREATE TABLE IF NOT EXISTS server (
		id INT NOT NULL AUTO_INCREMENT,
		name VARCHAR(20),
    	creation_time DATETIME,
    	PRIMARY KEY(id)
	)`)
	if err != nil {
		panic(err)
	}

	_, err = database.Instance.Exec(`CREATE TABLE IF NOT EXISTS server_has_member (
		server_id INT NOT NULL,
		user_id INT NOT NULL,
		FOREIGN KEY(server_id) REFERENCES server(id),
		FOREIGN KEY(user_id) REFERENCES user(id),
    	PRIMARY KEY(server_id, user_id)
	)`)
	if err != nil {
		panic(err)
	}
}

func CreateServer(server *Server) error {
	creationTime := time.Now()

	res, err := database.Instance.Exec("INSERT INTO server (name, creation_time) VALUES (?, ?)", server.Name, creationTime)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	server.ID = int(id)
	server.CreationTime = creationTime

	return nil
}

func FetchServers() ([]Server, error) {
	servers := make([]Server, 0)

	rows, err := database.Instance.Query("SELECT id, name, creation_time FROM server")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var server Server
		err = rows.Scan(&server.ID, &server.Name, &server.CreationTime)
		if err != nil {
			return nil, err
		}

		servers = append(servers, server)
	}
	return servers, nil
}

func FetchServerByID(id int) (*Server, error) {
	row := database.Instance.QueryRow("SELECT id, name, creation_time FROM server")

	server := &Server{}
	err := row.Scan(&server.ID, &server.Name, &server.CreationTime)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	members, err := FetchServerMembers(id)
	if err != nil {
		return nil, err
	}
	server.Members = members

	return server, nil
}

func FetchServerMembers(id int) ([]user.User, error) {
	members := make([]user.User, 0)

	rows, err := database.Instance.Query("SELECT user_id WHERE server_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var userId int
		err = rows.Scan(&userId)
		if err != nil {
			return nil, err
		}

		member, err := user.FetchUserById(userId)
		if err != nil {
			return nil, err
		}
		if member == nil {
			return nil, errors.New("Member not found")
		}
		members = append(members, *member)
	}

	return members, nil
}

func UpdateServer(server *Server) error {
	_, err := database.Instance.Exec("UPDATE server SET name = ? WHERE id = ?", server.Name, server.ID)
	return err
}

func DeleteServer(id int) error {
	channels, err := channel.FetchChannelsByServerId(id)
	if err != nil {
		return err
	}
	for _, c := range channels {
		err = channel.DeleteChannel(c.ID)
		if err != nil {
			return err
		}
	}

	_, err = database.Instance.Exec("DELETE FROM server_has_member WHERE server_id = ?", id)
	if err != nil {
		return err
	}

	_, err = database.Instance.Exec("DELETE FROM server WHERE id = ?", id)
	return err
}
