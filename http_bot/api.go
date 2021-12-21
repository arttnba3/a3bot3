package http_bot

import (
	"a3bot3/api"
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"time"
)

func (bot *HTTPBot) SendPrivateMsg(userId int64, message interface{}, autoEscape bool) int {
	// send request to cq-http
	var uuid = strconv.Itoa(bot.NextUUID()) + strconv.Itoa(time.Now().Nanosecond())
	var data = api.Action{
		Action: "send_private_msg",
		Params: api.PrivateParam{
			UserID:     userId,
			Message:    message,
			AutoEscape: autoEscape,
		},
		UUID: uuid,
	}
	log.Println("send message to user", userId, ":", message)
	bot.SendRequest(data)

	// get response
	resp, _ := bot.GetResponse(uuid)
	var respData api.ResponseBody
	err := json.Unmarshal(resp, &respData)
	if err != nil {
		log.Println("Invalid response from cq-http:", resp)
		return -1
	}
	return respData.Data.MessageID
}

func (bot *HTTPBot) SendGroupMsg(groupId int64, message interface{}, autoEscape bool) int {
	// send request to cq-http
	var uuid = strconv.Itoa(bot.NextUUID()) + strconv.Itoa(time.Now().Nanosecond())
	var data = api.Action{
		Action: "send_group_msg",
		Params: api.GroupParam{
			GroupID:    groupId,
			Message:    message,
			AutoEscape: autoEscape,
		},
		UUID: uuid,
	}
	log.Println("send message to group", groupId, ":", message)
	bot.SendRequest(data)

	// get response
	resp, _ := bot.GetResponse(uuid)
	var respData api.ResponseBody
	err := json.Unmarshal(resp, &respData)
	if err != nil {
		log.Println("Invalid response from cq-http:", resp)
		return -1
	}
	return respData.Data.MessageID
}

func (bot *HTTPBot) SendMsg(messageType string, userId int64, groupId int64, message interface{}, autoEscape bool) (int, error) {
	switch messageType {
	case string("private"):
		return bot.SendPrivateMsg(userId, message, autoEscape), nil
	case string("group"):
		return bot.SendGroupMsg(groupId, message, autoEscape), nil
	default:
		log.Println("Invalid type of message! Check your code!")
		return -1, errors.New("invalid message type")
	}
}

func (bot *HTTPBot) DeleteMsg(messageId int) {
	// send request to cq-http
	var uuid = strconv.Itoa(bot.NextUUID()) + strconv.Itoa(time.Now().Nanosecond())
	var data = api.Action{
		Action: "delete_msg",
		Params: api.MessageParam{
			MessageID: messageId,
		},
		UUID: uuid,
	}
	log.Println("try to delete message, id:", messageId)
	bot.SendRequest(data)
	_, _ = bot.GetResponse(uuid)
}

func (bot *HTTPBot) GetMsg(messageId int) (api.Message, error) {
	// send request to cq-http
	var uuid = strconv.Itoa(bot.NextUUID()) + strconv.Itoa(time.Now().Nanosecond())
	var data = api.Action{
		Action: "get_msg",
		Params: api.MessageParam{
			MessageID: messageId,
		},
		UUID: uuid,
	}
	bot.SendRequest(data)

	// get response
	resp, _ := bot.GetResponse(uuid)
	var respData api.ResponseBody
	err := json.Unmarshal(resp, &respData)
	if err != nil {
		log.Println("Invalid response from cq-http:", resp)
		return api.Message{}, errors.New("invalid response")
	}
	return api.Message{
		Message:     respData.Data.Message,
		MessageID:   respData.Data.MessageID,
		MessageType: respData.Data.MessageType,
		Sender:      respData.Data.Sender,
		ReadID:      respData.Data.ReadID,
		Time:        respData.Data.Time,
	}, nil
}
