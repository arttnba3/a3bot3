package plugin

import (
	"a3bot3/api"
	"a3bot3/event"
)

type TodoPlugin struct {
	PluginInfo
}

func init() {

}

func (p *TodoPlugin) PrivateMsgHandler(bot api.BotAPI, privateEvent event.PrivateEvent, messages []string) int {
	return MESSAGE_IGNORE
}

func (p *TodoPlugin) GroupMsgHandler(bot api.BotAPI, groupEvent event.GroupEvent, messages []string) int {
	return MESSAGE_IGNORE
}
