package bot

import (
	"a3bot3/api"
	"a3bot3/config"
	"a3bot3/event"
	"a3bot3/plugin"
	"a3bot3/tools"
	"log"
	"reflect"
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
			PrivateEventHandler(A3bot, event.PrivateEvent{Event: perEvent}, messages, "SendPrivateMsg")
		case "group":
			log.Println("received group message:\""+perEvent.Message+"\" from sender:", perEvent.Sender.UserID, "at group:", perEvent.GroupID)
			GroupEventHandler(A3bot, event.GroupEvent{Event: perEvent}, messages, "SendGroupMsg")
		default:
			log.Println("Unsupported message type:", perEvent.MessageType)
		}
	case "notice":
		switch perEvent.NoticeType {
		case "friend_recall":
			log.Println("user", perEvent.UserID, "'s message", perEvent.MessageID, "is recalled.")
			PrivateEventHandler(A3bot, event.PrivateEvent{Event: perEvent}, messages, "FriendRecallHandler")
		case "group_recall":
			log.Println("user", perEvent.UserID, "'s message", perEvent.MessageID, "in group:", perEvent.GroupID, "is recalled.")
			GroupEventHandler(A3bot, event.GroupEvent{Event: perEvent}, messages, "GroupRecallHandler")
		default:
			log.Println("Unsupported notice type:", perEvent.NoticeType)
		}
	default:
		log.Println("Invalid post type! Check your cq-http or contact the developer!")
	}
}

// PrivateEventHandler :
// Use the reflection to call specific event handler for every plugin.
// This func deal with the private message.
func PrivateEventHandler(bot api.BotAPI, privateEvent event.PrivateEvent, messages []string, handlerName string) {
	for _, perPlugin := range plugin.Plugins {
		if !perPlugin.IsEnable() {
			continue
		}

		val := reflect.ValueOf(perPlugin)
		handlerFunc := val.MethodByName(handlerName)
		args := []reflect.Value{reflect.ValueOf(bot), reflect.ValueOf(privateEvent), reflect.ValueOf(messages)}

		handlerFunc.Call(args)
	}
}

// GroupEventHandler :
// Use the reflection to call specific event handler for every plugin.
// This func deal with the group message.
func GroupEventHandler(bot api.BotAPI, groupEvent event.GroupEvent, messages []string, handlerName string) {
	for _, perPlugin := range plugin.Plugins {
		if !perPlugin.IsEnable() {
			continue
		}

		val := reflect.ValueOf(perPlugin)
		handlerFunc := val.MethodByName(handlerName)
		args := []reflect.Value{reflect.ValueOf(bot), reflect.ValueOf(groupEvent), reflect.ValueOf(messages)}

		handlerFunc.Call(args)
	}
}
