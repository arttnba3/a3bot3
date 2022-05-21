package plugin

import (
	"a3bot3/api"
	"a3bot3/event"
)

type PluginInfo struct {
	Enable  bool
	Name    string
	Command string
	Plugin
}

type Plugin interface {
	SendPrivateMsg(bot api.BotAPI, privateEvent event.PrivateEvent, messages []string) int
	SendGroupMsg(bot api.BotAPI, groupEvent event.GroupEvent, messages []string) int
	IsEnable() bool
	GetName() string
	MatchCommand(cmd string) bool
	SetEnable(enable bool)
}

var MESSAGE_BLOCK = 1
var MESSAGE_IGNORE = 0

// Plugins :
// This is the array of instances of all available plugins.
var Plugins = [...]Plugin{
	&PluginManager{
		PluginInfo: PluginInfo{
			Enable:  true,
			Name:    "PluginManager",
			Command: "/plugin",
		},
	},
	&ExamplePlugin{
		PluginInfo: PluginInfo{
			Enable:  false,
			Name:    "ExamplePlugin",
			Command: "/hello",
		},
	},
	&RepeaterPlugin{
		PluginInfo: PluginInfo{
			Enable:  true,
			Name:    "RepeaterPlugin",
			Command: "",
			Plugin:  nil,
		},
	},
}

func (p *PluginInfo) IsEnable() bool {
	return p.Enable
}

func (p *PluginInfo) GetName() string {
	return p.Name
}

func (p *PluginInfo) MatchCommand(cmd string) bool {
	return cmd == p.Command
}

func (p *PluginInfo) SetEnable(enable bool) {
	p.Enable = enable
}
