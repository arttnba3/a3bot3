package bot

import (
	"a3bot3/api"
	"a3bot3/config"
	"a3bot3/event"
	"a3bot3/plugin"
	"a3bot3/tools"
	"log"
)

var A3bot api.BotAPI

// MessageHandler
/*
 * All the messages will be passed by from Plugins as the order in array Plugins,
 * unless it met the MESSAGE_BLOCK as ret val, it'll be blocked, the traversal ends.
 * You can control status of each plugin with the /plugin command.
 */
func MessageHandler(perEvent event.Event) {
	var messages []string

	switch perEvent.PostType {
	// heartbeat, ignore
	case "meta_event":
		return
	case "message":
		/*
		 * Except the Event.Message, I also provide you with another kind of messages,
		 * simply parsed into several strings by my own parser.
		 * If you'd like so, just use the messages params simply.
		 */
		if config.Settings.ParseCommandOnly {
			msg := perEvent.Message
			if (len(msg) > 0) && (msg[0] == '/') {
				messages = tools.AutoParser(msg)
			}
		} else {
			messages = tools.AutoParser(perEvent.Message)
		}
		switch perEvent.MessageType {
		case "private":
			log.Println("received private message:\""+perEvent.Message+"\" from sender:", perEvent.UserID)
			for _, perPlugin := range plugin.Plugins {
				if !perPlugin.IsEnable() {
					continue
				}
				if perPlugin.SendPrivateMsg(A3bot, event.PrivateEvent{Event: perEvent}, messages) == plugin.MESSAGE_BLOCK {
					return
				}
			}
			return
		case "group":
			log.Println("received group message:\""+perEvent.Message+"\" from sender:", perEvent.Sender.UserID, "at group:", perEvent.GroupID)
			for _, perPlugin := range plugin.Plugins {
				if !perPlugin.IsEnable() {
					continue
				}
				if perPlugin.SendGroupMsg(A3bot, event.GroupEvent{Event: perEvent}, messages) == plugin.MESSAGE_BLOCK {
					return
				}
			}
			return
		default:
			log.Println("Unsupported message type:", perEvent.MessageType)
		}
	case "notice":
		return // not done yet
	default:
		log.Println("Invalid post type! Check your cq-http or contact the developer!")
	}
}
