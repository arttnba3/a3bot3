package plugin

import (
	"a3bot3/api"
	"a3bot3/event"
)

type Plugin interface {
	SendPrivateMsg(bot api.BotAPI, privateEvent event.PrivateEvent, messages []string) int
	SendGroupMsg(bot api.BotAPI, groupEvent event.GroupEvent, messages []string) int
	IsEnable() bool
	GetName() string
	MatchCommand(cmd string) bool
	SetEnable(enable bool)
	GroupRecallHandler(bot api.BotAPI, groupEvent event.GroupEvent, messages []string) int
	FriendRecallHandler(bot api.BotAPI, privateEvent event.PrivateEvent, messages []string) int
	GroupUploadHandler(bot api.BotAPI, groupEvent event.GroupEvent, messages []string) int
	GroupAdminHandler(bot api.BotAPI, groupEvent event.GroupEvent, messages []string) int
	GroupIncreaseHandler(bot api.BotAPI, groupEvent event.GroupEvent, messages []string) int
	GroupDecreaseHandler(bot api.BotAPI, groupEvent event.GroupEvent, messages []string) int
}

var MESSAGE_BLOCK = 1
var MESSAGE_IGNORE = 0

// SendPrivateMsg :
// handler with private message
func (p *PluginInfo) SendPrivateMsg(bot api.BotAPI, privateEvent event.PrivateEvent, messages []string) int {
	return MESSAGE_IGNORE
}

// SendGroupMsg :
// handler with group message
func (p *PluginInfo) SendGroupMsg(bot api.BotAPI, groupEvent event.GroupEvent, messages []string) int {
	return MESSAGE_IGNORE
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

// GroupRecallHandler :
// handle with group-recalling event
func (p *PluginInfo) GroupRecallHandler(bot api.BotAPI, groupEvent event.GroupEvent, messages []string) int {
	return MESSAGE_IGNORE
}

// FriendRecallHandler :
// handle with friend-recalling event
func (p *PluginInfo) FriendRecallHandler(bot api.BotAPI, privateEvent event.PrivateEvent, messages []string) int {
	return MESSAGE_IGNORE
}

// GroupUploadHandler :
// handle with file-uploading event in group
func (p *PluginInfo) GroupUploadHandler(bot api.BotAPI, groupEvent event.GroupEvent, messages []string) int {
	return MESSAGE_IGNORE
}

// GroupAdminHandler :
// handle with someone-becoming-admin event in group
func (p *PluginInfo) GroupAdminHandler(bot api.BotAPI, groupEvent event.GroupEvent, messages []string) int {
	return MESSAGE_IGNORE
}

// GroupIncreaseHandler :
// handle with: someone joined the group.
// Check event.SubType for type of joining.
func (p *PluginInfo) GroupIncreaseHandler(bot api.BotAPI, groupEvent event.GroupEvent, messages []string) int {
	return MESSAGE_IGNORE
}

// GroupDecreaseHandler :
// handle with: someone left the group.
// Check event.SubType for type of leaving.
func (p *PluginInfo) GroupDecreaseHandler(bot api.BotAPI, groupEvent event.GroupEvent, messages []string) int {
	return MESSAGE_IGNORE
}
