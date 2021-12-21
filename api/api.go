package api

import (
	"sync"
)

type Bot struct {
	Host      string
	Port      int
	Responses sync.Map
	UUID      int
	Lock      sync.Mutex
}

type Action struct {
	Action string      `json:"action"`
	Params interface{} `json:"params"`
	UUID   string      `json:"uuid"`
}

type PrivateParam struct {
	UserID     int64       `json:"user_id"`
	Message    interface{} `json:"message"`
	AutoEscape bool        `json:"auto_escape"`
}

type GroupParam struct {
	GroupID    int64       `json:"group_id"`
	Message    interface{} `json:"message"`
	AutoEscape bool        `json:"auto_escape"`
}

type MessageParam struct {
	MessageID int `json:"message_id"`
}

type ResponseBody struct {
	Data    ResponseData `json:"data"`
	Retcode int          `json:"retcode"`
	Status  string       `json:"status"`
}

type ResponseData struct {
	Message     string `json:"message"`
	MessageID   int    `json:"message_id"`
	Group       bool   `json:"group"`
	MessageSeq  int    `json:"message_seq"`
	MessageType string `json:"message_type"`
	ReadID      int    `json:"real_id"`
	Sender      Sender `json:"sender"`
	Time        int    `json:"time"`
}

type Sender struct {
	Nickname string `json:"nickname"`
	UserID   int64  `json:"user_id"`
}

type Message struct {
	Message     string `json:"message"`
	MessageID   int    `json:"message_id"`
	MessageType string `json:"message_type"`
	ReadID      int    `json:"real_id"`
	Sender      Sender `json:"sender"`
	Time        int    `json:"time"`
}

type BotAPI interface {
	SendPrivateMsg(userId int64, message interface{}, autoEscape bool) int
	SendGroupMsg(groupId int64, message interface{}, autoEscape bool) int
	SendMsg(messageType string, userId int64, groupId int64, message interface{}, autoEscape bool) (int, error)
	DeleteMsg(messageId int)
	GetMsg(messageId int) (Message, error)
}

func (bot *Bot) NextUUID() int {
	var uuid int

	bot.Lock.Lock()
	uuid = bot.UUID
	bot.UUID += 1
	defer bot.Lock.Unlock()

	return uuid
}
