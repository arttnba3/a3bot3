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
	MessageID int32 `json:"message_id"`
}

type ResponseBody struct {
	Data    ResponseData `json:"data"`
	Retcode int          `json:"retcode"`
	Status  string       `json:"status"`
}

type ResponseData struct {
	Message     string `json:"message"`
	MessageID   int32  `json:"message_id"`
	Group       bool   `json:"group"`
	MessageSeq  int64  `json:"message_seq"`
	MessageType string `json:"message_type"`
	ReadID      int32  `json:"real_id"`
	Sender      Sender `json:"sender"`
	Time        int32  `json:"time"`
}

type Sender struct {
	Nickname string `json:"nickname"`
	UserID   int64  `json:"user_id"`
}

type Message struct {
	Message     string `json:"message"`
	MessageID   int32  `json:"message_id"`
	MessageType string `json:"message_type"`
	ReadID      int32  `json:"real_id"`
	Sender      Sender `json:"sender"`
	Time        int32  `json:"time"`
}

type NodeData struct {
	ID      int32       `json:"id"`
	Name    string      `json:"name"`
	UIN     int64       `json:"uin"`
	Content interface{} `json:"content"`
	Seq     interface{} `json:"seq"`
}

type Node struct {
	Type string   `json:"type"`
	Data NodeData `json:"data"`
}

type BotAPI interface {
	SendPrivateMsg(userId int64, message interface{}, autoEscape bool) int32
	SendGroupMsg(groupId int64, message interface{}, autoEscape bool) int32
	SendGroupForwardMsg(groupId int64, message interface{})
	SendMsg(messageType string, userId int64, groupId int64, message interface{}, autoEscape bool) (int32, error)
	DeleteMsg(messageId int32)
	GetMsg(messageId int32) (Message, error)
}

func (bot *Bot) NextUUID() int {
	var uuid int

	bot.Lock.Lock()
	uuid = bot.UUID
	bot.UUID += 1
	defer bot.Lock.Unlock()

	return uuid
}
