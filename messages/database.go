package Message

import (
	"fmt"
	"time"

	Cassandra "../../../src/messageapi/cassandra"
	"github.com/gocql/gocql"
)

type Message struct {
	ID          gocql.UUID `json:"id"`
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	Email       string     `json:"email"`
	MagicNumber int        `json:"magic_number"`
	CreateTime  time.Time  `json:"create_time"`
}

func InsertMessage(message Message) (uuid gocql.UUID, err error) {
	uuid = gocql.TimeUUID()

	query := Cassandra.Session.Query(`
		INSERT INTO message (id, title, content, email, magic_number, create_time)
		VALUES (?, ?, ?, ?, ?, ?)`,
		uuid, message.Title, message.Content, message.Email, message.MagicNumber, time.Now())

	err = query.Exec()

	return uuid, err
}

func FindByEmail(email string) (messageList []Message) {
	dict := map[string]interface{}{}

	fmt.Println(email)

	query := Cassandra.Session.Query(`
			SELECT *
			FROM message
			WHERE email LIKE '?'`, email)

	iterable := query.Iter()
	fmt.Println(iterable.NumRows())

	for iterable.MapScan(dict) {
		fmt.Println(dict["id"])
		messageList = append(messageList, Message{
			ID:          dict["id"].(gocql.UUID),
			Title:       dict["title"].(string),
			Content:     dict["content"].(string),
			Email:       dict["email"].(string),
			MagicNumber: dict["magic_number"].(int),
			CreateTime:  dict["create_time"].(time.Time),
		})
		dict = map[string]interface{}{}
	}

	return messageList
}
