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
	case "meta_event":
		// heartbeat, ignore it
		return
	case "message":
		/*
		 * Except the Event.Message, I also provide you with another kind of messages,
		 * simply parsed into several strings by my own parser.
		 * If you'd like so, just use the messages params simply.
		 */
		if config.Settings.ParseCommandOnly {
			msg := perEvent.RawMessage
			if (len(msg) > 0) && (msg[0] == '/') {
				messages = tools.AutoParser(msg)
			}
		} else {
			messages = tools.AutoParser(perEvent.RawMessage)
		}
		switch perEvent.MessageType {
		case "private":
			log.Println("received private message:\""+perEvent.RawMessage+"\" from sender:", perEvent.UserID)
			PrivateEventHandler(A3bot, event.PrivateEvent{Event: perEvent}, messages, "PrivateMsgHandler")
		case "group":
			log.Println("received group message:\""+perEvent.RawMessage+"\" from sender:", perEvent.Sender.UserID, "at group:", perEvent.GroupID)
			GroupEventHandler(A3bot, event.GroupEvent{Event: perEvent}, messages, "GroupMsgHandler")
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
		case "group_upload":
			log.Println("user", perEvent.UserID, "upload a file", perEvent.File.Name, "in group:", perEvent.GroupID)
			GroupEventHandler(A3bot, event.GroupEvent{Event: perEvent}, messages, "GroupUploadHandler")
		case "group_admin":
			log.Println("user", perEvent.UserID, "'s admin", perEvent.MessageID, "in group:", perEvent.GroupID, "is", perEvent.SubType)
			GroupEventHandler(A3bot, event.GroupEvent{Event: perEvent}, messages, "GroupAdminHandler")
		case "group_increase":
			log.Println("user", perEvent.UserID, "joined the group", perEvent.GroupID, ", type:", perEvent.SubType)
			GroupEventHandler(A3bot, event.GroupEvent{Event: perEvent}, messages, "GroupIncreaseHandler")
		case "group_decrease":
			log.Println("user", perEvent.UserID, "left the group", perEvent.GroupID, ", type:", perEvent.SubType)
			GroupEventHandler(A3bot, event.GroupEvent{Event: perEvent}, messages, "GroupDecreaseHandler")
		default:
			log.Println("Unsupported notice type:", perEvent.NoticeType)
		}
	case "request":
		// no implementation yet
		return
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

		if handlerFunc.Call(args)[0].Int() == int64(plugin.MESSAGE_BLOCK) {
			return
		}
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

		if handlerFunc.Call(args)[0].Int() == int64(plugin.MESSAGE_BLOCK) {
			return
		}
	}
}
