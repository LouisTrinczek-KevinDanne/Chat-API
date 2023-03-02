package channel

import (
	"database/sql"
	"errors"
	"github.com/LouisTrinczek-KevinDanne/Chat-API/database"
	"github.com/LouisTrinczek-KevinDanne/Chat-API/modules/user"
	"time"
)

type Channel struct {
	ID       int              `json:"id"`
	Name     string           `json:"name"`
	ServerId int              `json:"serverId"`
	Messages []ChannelMessage `json:"messages"`
}

type ChannelMessage struct {
	ID        int       `json:"id"`
	ChannelId int       `json:"channelId"`
	Sender    user.User `json:"sender"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func init() {
	_, err := database.Instance.Exec(`CREATE TABLE IF NOT EXISTS channel (
		id INT NOT NULL AUTO_INCREMENT,
		server_id INT NOT NULL,
		name VARCHAR(20),
    	FOREIGN KEY(server_id) REFERENCES server(id),
    	PRIMARY KEY(id)
	)`)
	if err != nil {
		panic(err)
	}

	_, err = database.Instance.Exec(`CREATE TABLE IF NOT EXISTS channel_message (
		id INT NOT NULL AUTO_INCREMENT,
		channel_id INT NOT NULL,
		sender_id INT NOT NULL,
		messsage TEXT,
		sent_date DATETIME,
    	FOREIGN KEY(channel_id) REFERENCES channel(id),
    	FOREIGN KEY(sender_id) REFERENCES user(id),
    	PRIMARY KEY(id)
	)`)
	if err != nil {
		panic(err)
	}
}

func CreateChannel(channel *Channel) error {
	res, err := database.Instance.Exec("INSERT INTO channel (server_id, name) VALUES (?, ?)", channel.ServerId, channel.Name)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	channel.ID = int(id)

	return nil
}

func FetchChannels() ([]Channel, error) {
	channels := make([]Channel, 0)

	rows, err := database.Instance.Query("SELECT id, server_id, name FROM channel")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var channel Channel
		err = rows.Scan(&channel.ID, &channel.ServerId, &channel.Name)
		if err != nil {
			return nil, err
		}

		channels = append(channels, channel)
	}
	return channels, nil
}

func FetchChannelByID(id int) (*Channel, error) {
	row := database.Instance.QueryRow("SELECT id, server_id, name FROM channel WHERE id = ?", id)

	channel := &Channel{}
	err := row.Scan(&channel.ID, &channel.ServerId, &channel.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	messages, err := FetchChannelMessages(id)
	if err != nil {
		return nil, err
	}
	channel.Messages = messages

	return channel, nil
}

func FetchChannelMessages(id int) ([]ChannelMessage, error) {
	messages := make([]ChannelMessage, 0)

	rows, err := database.Instance.Query("SELECT id, channel_id, sender_id, message, sent_date FROM channel_message WHERE channel_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var message ChannelMessage
		var senderId int
		err = rows.Scan(&message.ID, &senderId, &message.Message, &message.Timestamp)
		if err != nil {
			return nil, err
		}

		sender, err := user.FetchUserById(senderId)
		if err != nil {
			return nil, err
		}
		if sender == nil {
			return nil, errors.New("Sender not found")
		}
		message.Sender = *sender

		messages = append(messages, message)
	}

	return messages, nil
}

func FetchChannelsByServerId(id int) ([]Channel, error) {
	channels := make([]Channel, 0)

	rows, err := database.Instance.Query("SELECT id, server_id, name FROM channel WHERE server_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var channel Channel
		err = rows.Scan(&channel.ID, &channel.ServerId, &channel.Name)
		if err != nil {
			return nil, err
		}

		channels = append(channels, channel)
	}

	return channels, nil
}

func UpdateChannel(channel *Channel) error {
	_, err := database.Instance.Exec("UPDATE channel SET name = ? WHERE id = ?", channel.Name, channel.ID)
	return err
}

func DeleteChannel(id int) error {
	_, err := database.Instance.Exec("DELETE FROM channel_message WHERE channel_id = ?", id)
	if err != nil {
		return err
	}

	_, err = database.Instance.Exec("DELETE FROM channel WHERE id = ?", id)
	return err
}
