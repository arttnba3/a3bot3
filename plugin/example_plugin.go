package plugin

import (
	"a3bot3/api"
	"a3bot3/event"
)

// Following are an example of plugin

type ExamplePlugin struct {
	PluginInfo
}

func (p *ExamplePlugin) SendPrivateMsg(bot api.BotAPI, privateEvent event.PrivateEvent, messages []string) int {
	// continuous passing message to next plugin
	if messages == nil || len(messages) < 1 || !p.MatchCommand(messages[0]) {
		return MESSAGE_IGNORE
	}

	// do not pass message to next plugin
	bot.SendPrivateMsg(privateEvent.UserID, "Hello world! I'm a3bot3!", false)
	return MESSAGE_BLOCK

}

func (p *ExamplePlugin) SendGroupMsg(bot api.BotAPI, groupEvent event.GroupEvent, messages []string) int {
	// continuous passing message to next plugin
	if messages == nil || len(messages) < 1 || !p.MatchCommand(messages[0]) {
		return MESSAGE_IGNORE
	}

	// do not pass message to next plugin
	bot.SendGroupMsg(groupEvent.GroupID, "Hello world! I'm a3bot3!", false)
	return MESSAGE_BLOCK
}
