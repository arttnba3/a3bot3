package bot

import (
	"a3bot3/api"
	"a3bot3/event"
	"a3bot3/plugin"
	"log"
)

var A3bot api.BotAPI

var plugins = [...]plugin.Plugin{
	plugin.ExamplePLugin,
}

func MessageHandler(perEvent event.Event) {
	switch perEvent.PostType {
	// heartbeat, ignore
	case string("meta_event"):
		return
	case string("message"):
		switch perEvent.MessageType {
		case string("private"):
			log.Println("received private message:\""+perEvent.Message+"\" from sender:", perEvent.UserID)
			for i := 0; i < len(plugins); i++ {
				if plugins[i].SendPrivateMsg(A3bot, event.PrivateEvent{Event: perEvent}) == plugin.MESSAGE_BLOCK {
					return
				}
			}
			return
		case string("group"):
			log.Println("received group message:\""+perEvent.Message+"\" from sender:", perEvent.Sender.UserID, "at group:", perEvent.GroupID)
			for i := 0; i < len(plugins); i++ {
				if plugins[i].SendGroupMsg(A3bot, event.GroupEvent{Event: perEvent}) == plugin.MESSAGE_BLOCK {
					return
				}
			}
			return
		default:
			log.Println("Invalid message type!")
		}
	case string("notice"):
		return // not done yet
	default:
		log.Println("Invalid post type!")
	}
}
