package plugin

import (
	"a3bot3/api"
	"a3bot3/event"
)

type Plugin struct {
	Enable bool
	Name   string
}

type Operations interface {
	SendPrivateMsg(bot api.BotAPI, privateEvent event.PrivateEvent) int
	SendGroupMsg(bot api.BotAPI, groupEvent event.GroupEvent) int
}

var MESSAGE_BLOCK = 1
var MESSAGE_IGNORE = 0

// Following are an example of plugin

var ExamplePLugin = Plugin{
	Enable: true,
	Name:   "ExamplePlugin",
}

func (p *Plugin) SendPrivateMsg(bot api.BotAPI, privateEvent event.PrivateEvent) int {
	// do not pass message to next plugin
	if privateEvent.Message == "Hello a3bot3!" {
		bot.SendPrivateMsg(privateEvent.UserID, "Hello world!", false)
	}
	// continuous passing message to next plugin
	return MESSAGE_IGNORE
}

func (p *Plugin) SendGroupMsg(bot api.BotAPI, groupEvent event.GroupEvent) int {
	// do not pass message to next plugin
	if groupEvent.Message == "Hello a3bot3!" {
		bot.SendGroupMsg(groupEvent.GroupID, "Hello group!", false)
		return MESSAGE_BLOCK
	} else {
		// continuous passing message to next plugin
		return MESSAGE_IGNORE
	}
}
