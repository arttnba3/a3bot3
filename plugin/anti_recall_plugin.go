package plugin

import (
	"a3bot3/api"
	"a3bot3/event"
	"encoding/json"
	"log"
	"strconv"
)

type AntiRecallPlugin struct {
	PluginInfo
}

func (p *AntiRecallPlugin) GroupRecallHandler(bot api.BotAPI, groupEvent event.GroupEvent, messages []string) int {
	recalledMsg, err := bot.GetMsg(groupEvent.MessageID)
	if err != nil {
		log.Println("failed to get recalled message, id:", groupEvent.MessageID)
		return MESSAGE_IGNORE
	}

	replyMsg := "detected recalled message from " + strconv.FormatInt(groupEvent.UserID, 10) +
		", operated by " + strconv.FormatInt(groupEvent.OperatorID, 10) + ":\n"

	recalledMsgItem := recalledMsg.Message[0]

	if recalledMsgItem.Type == "text" {
		var recalledMsgText api.MessageItemData

		recalledMsgItemJson, _ := json.Marshal(recalledMsgItem.Data)
		err = json.Unmarshal(recalledMsgItemJson, &recalledMsgText)
		if err != nil {
			replyMsg += "(Unsupported data content: \"" + string(recalledMsgItemJson) + "\")"
		} else {
			replyMsg += recalledMsgText.Text
		}
	} else {
		replyMsg += "(Unsupported data type: \"" + recalledMsgItem.Type + "\")"
	}

	bot.SendGroupMsg(groupEvent.GroupID, replyMsg, false)

	return MESSAGE_BLOCK
}

func (p *AntiRecallPlugin) FriendRecallHandler(bot api.BotAPI, privateEvent event.PrivateEvent, messages []string) int {
	recalledMsg, err := bot.GetMsg(privateEvent.MessageID)
	if err != nil {
		log.Println("failed to get recalled message, id:", privateEvent.MessageID)
		return MESSAGE_IGNORE
	}

	replyMsg := "detected your recalled message :\n"

	recalledMsgItem := recalledMsg.Message[0]

	if recalledMsgItem.Type == "text" {
		var recalledMsgText api.MessageItemData

		recalledMsgItemJson, _ := json.Marshal(recalledMsgItem.Data)
		err = json.Unmarshal(recalledMsgItemJson, &recalledMsgText)
		if err != nil {
			replyMsg += "(Unsupported data content: \"" + string(recalledMsgItemJson) + "\")"
		} else {
			replyMsg += recalledMsgText.Text
		}
	} else {
		replyMsg += "(Unsupported data type: \"" + recalledMsgItem.Type + "\")"
	}

	bot.SendPrivateMsg(privateEvent.UserID, replyMsg, false)

	return MESSAGE_BLOCK
}
