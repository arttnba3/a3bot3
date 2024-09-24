package plugin

import (
	"a3bot3/api"
	"a3bot3/event"
)

var lastMsgGroup = make(map[int64]string)
var perGroupRepeated = make(map[int64]bool)

type RepeaterPlugin struct {
	PluginInfo
}

func (p *RepeaterPlugin) PrivateMsgHandler(bot api.BotAPI, privateEvent event.PrivateEvent, messages []string) int {
	// we only repeat in group
	return MESSAGE_IGNORE
}

func (p *RepeaterPlugin) GroupMsgHandler(bot api.BotAPI, groupEvent event.GroupEvent, messages []string) int {
	groupId := groupEvent.GroupID

	if lastMsgGroup[groupId] == groupEvent.RawMessage {
		if !perGroupRepeated[groupId] {
			bot.SendGroupMsg(groupId, groupEvent.RawMessage, false)
			perGroupRepeated[groupId] = true
		}
	} else {
		perGroupRepeated[groupId] = false
	}

	lastMsgGroup[groupId] = groupEvent.RawMessage

	return MESSAGE_BLOCK
}
