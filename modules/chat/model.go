package chat

import (
	"github.com/LouisTrinczek-KevinDanne/Chat-API/modules/user"
	"time"
)

type Server struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	Channels     []Channel   `json:"channels"`
	Members      []user.User `json:"members"`
	CreationTime time.Time   `json:"creationTime"`
}

type Channel struct {
	ID       int              `json:"id"`
	Name     string           `json:"name"`
	Messages []ChannelMessage `json:"messages"`
}

type ChannelMessage struct {
	ID        int       `json:"id"`
	Sender    user.User `json:"sender"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func CreateServer(server *Server) error {
	// TODO create logic
	server.ID = 999
	return nil
}

func FetchServers() ([]Server, error) {
	servers := make([]Server, 0)
	// TODO fetch logic
	return servers, nil
}

func FetchServerByID(id int) (*Server, error) {
	// todo fetch logic
	return nil, nil
}

func UpdateServer(server *Server) error {
	// todo update logic
	return nil
}

func DeleteServer(id int) error {
	// todo delete logic
	return nil
}
