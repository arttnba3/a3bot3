package main

import (
	"a3bot3/bot"
	"a3bot3/tools"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func Shell() {
	time.Sleep(time.Second * 1)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("a3bot3 shell> ")
		command, err := reader.ReadBytes('\n')
		if err != nil {
			log.Println(err)
			break
		}

		exit := execCommand(tools.AutoParser(string(command[:len(command)-1])))
		if exit {
			fmt.Println("see you next time~")
			break
		}
	}
}

func execCommand(commands []string) bool {
	if commands == nil || len(commands) == 0 {
		return false
	}

	switch commands[0] {
	case "exit":
		return true
	case "help":
		helpInfo()
	case "echo":
		echoMsg(commands)
	default:
		fmt.Println("[x] Command \"" + commands[0] + "\" not found.")
	}
	return false
}

func helpInfo() {
	fmt.Println("echo [type] [number] [message] --- send a message to a user/group")
	fmt.Println("help                           ---show help info")
	fmt.Println("Redundant params will be ignored.")
}

func echoMsg(commands []string) {
	if len(commands) < 4 {
		fmt.Println("[x] No enough params.")
	} else {
		targetId, err := strconv.ParseInt(commands[2], 10, 64)
		if err != nil {
			fmt.Println("[x] invalid number.")
		} else {
			switch strings.ToLower(commands[1]) {
			case "user":
				bot.A3bot.SendPrivateMsg(targetId, commands[3], false)
			case "group":
				bot.A3bot.SendGroupMsg(targetId, commands[3], false)
			default:
				fmt.Println("[x] invalid type.")
			}
		}
	}
}
