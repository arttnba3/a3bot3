package plugin

import (
	"a3bot3/api"
	"a3bot3/event"
	"log"
	"strconv"
)

type AntiRecallPlugin struct {
	PluginInfo
}

func (p *AntiRecallPlugin) GroupRecallHandler(bot api.BotAPI, groupEvent event.GroupEvent, messages []string) int {
	recallMsg, err := bot.GetMsg(int(groupEvent.MessageID))
	if err != nil {
		log.Println("failed to get recalled message, id:", groupEvent.MessageID)
		return MESSAGE_IGNORE
	}

	replyMsg := "detected recalled message from " + strconv.FormatInt(groupEvent.UserID, 10) +
		", operated by " + strconv.FormatInt(groupEvent.OperatorID, 10) + ":\n"
	replyMsg += recallMsg.Message

	bot.SendGroupMsg(groupEvent.GroupID, replyMsg, false)

	return MESSAGE_BLOCK
}

func (p *AntiRecallPlugin) FriendRecallHandler(bot api.BotAPI, privateEvent event.PrivateEvent, messages []string) int {
	recallMsg, err := bot.GetMsg(int(privateEvent.MessageID))
	if err != nil {
		log.Println("failed to get recalled message, id:", privateEvent.MessageID)
		return MESSAGE_IGNORE
	}

	replyMsg := "detected your recalled message :\n" + recallMsg.Message
	bot.SendPrivateMsg(privateEvent.UserID, replyMsg, false)

	return MESSAGE_BLOCK
}
