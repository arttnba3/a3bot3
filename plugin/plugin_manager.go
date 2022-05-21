package plugin

import (
	"a3bot3/api"
	"a3bot3/config"
	"a3bot3/event"
	"a3bot3/tools"
)

type PluginManager struct {
	PluginInfo
}

func (p *PluginManager) SendPrivateMsg(bot api.BotAPI, privateEvent event.PrivateEvent, messages []string) int {
	var replyMsg string

	// re-parsing if necessary
	if messages == nil || len(messages) < 1 {
		messages = tools.AutoParser(privateEvent.Message)
	}

	if len(messages) < 1 || !p.MatchCommand(messages[0]) {
		return MESSAGE_IGNORE
	}

	// check authentication
	if privateEvent.UserID != config.Settings.Admin {
		replyMsg = "Permission denied. Authentication limited."
	} else {
		// operate
		replyMsg = pluginManagerSolver(messages)
	}

	bot.SendPrivateMsg(privateEvent.UserID, replyMsg, false)

	return MESSAGE_BLOCK
}

func (p *PluginManager) SendGroupMsg(bot api.BotAPI, groupEvent event.GroupEvent, messages []string) int {
	var replyMsg string

	// re-parsing if necessary
	if messages == nil || len(messages) < 1 {
		messages = tools.AutoParser(groupEvent.Message)
	}

	if len(messages) < 1 || !p.MatchCommand(messages[0]) {
		return MESSAGE_IGNORE
	}

	// check authentication
	if groupEvent.UserID != config.Settings.Admin {
		replyMsg = "Permission denied. Authentication limited."
	} else {
		// operate
		replyMsg = pluginManagerSolver(messages)
	}

	bot.SendGroupMsg(groupEvent.GroupID, replyMsg, false)

	return MESSAGE_BLOCK
}

func pluginManagerSolver(messages []string) string {
	var replyMsg = ""

	if len(messages) < 2 {
		replyMsg = "Available plugins:\n"
		for _, perPlugin := range Plugins {
			if perPlugin.IsEnable() {
				replyMsg += perPlugin.GetName() + "\n"
			}
		}
	} else {
		switch messages[1] {
		case "help":
			replyMsg = "Usage: /plugin [operations] [params]\n" +
				"redundant args will be ignore automatically\n" +
				"Available operations:\n" +
				"(none)     ------ show enabled plugins\n" +
				"all        ------ show all the plugins\n" +
				"help       ------ show help info\n" +
				"Available for admin only:\n" +
				"load [name]   ---- load a plugin\n" +
				"unload [name] ---- unload a plugin"
		case "all":
			replyMsg = "All plugins ('*' means it's enabled):"
			for _, perPlugin := range Plugins {
				replyMsg += "\n" + perPlugin.GetName() + " ["
				if perPlugin.IsEnable() {
					replyMsg += "*"
				} else {
					replyMsg += " "
				}
				replyMsg += "]"
			}
		case "load", "unload":
			if len(messages) < 3 {
				replyMsg = "uncompleted command."
			} else {
				// search by plugin name
				var obj *Plugin = nil
				for _, perPlugin := range Plugins {
					if perPlugin.GetName() == messages[2] {
						obj = &perPlugin
						break
					}
				}
				// operate
				if obj == nil {
					replyMsg = "plugin not found."
				} else {
					if messages[1] == "load" {
						(*obj).SetEnable(true)
					} else {
						(*obj).SetEnable(false)
					}
					replyMsg = "Done."
				}
			}
		default:
			replyMsg = "invalid operations."
		}
	}

	return replyMsg
}
