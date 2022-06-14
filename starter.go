package main

import (
	"a3bot3/api"
	"a3bot3/bot"
	"a3bot3/bot/http_bot"
	"a3bot3/config"
	"a3bot3/info"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

/*
 * The default json parser of golang is a piece of shit
 * which cannot work correctly while the first char
 * of the filed is not an upper char
 */

func BotStart() {
	info.GreenMsg("================================================")
	info.GreenMsg("================     a3bot3     ================")
	info.GreenMsg("================ version  0.0.1 ================")
	info.GreenMsg("================================================")

	log.Println("Initialize log...")
	InitLog()
	log.Println("Log initialization done.")

	log.Println("Loading settings...")
	LoadingSettings()
	log.Println("Connecting...")

	switch config.Settings.Type {
	case string("http"):
		log.Println("http listening on:", config.Settings.Listen.Host, "port:", config.Settings.Listen.Port)
		log.Println("http posting on:", config.Settings.Post.Host, "port:", config.Settings.Post.Port)
		bot.A3bot = &http_bot.HTTPBot{
			Bot: api.Bot{
				Host: config.Settings.Post.Host,
				Port: config.Settings.Post.Port,
				UUID: 0,
			},
		}
		go http_bot.CQHTTPListener(config.Settings.Listen.Host, config.Settings.Listen.Port)
	case string("ws"):
		log.Panicln("websocket is coming soon.....")
	case string("reverse_ws"):
		log.Panicln("reverse websocket is coming soon...")
	default:
		log.Panicln("Invalid type of bot! Check your config.json file!")
	}
}

func LoadingSettings() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		info.RedMsg("FAILED TO READ CONFIG FILE config.json !")
		panic(err)
	}

	err = json.Unmarshal([]byte(data), &config.Settings)
	if err != nil {
		info.RedMsg("ERROR OCCUR WHILE PARSING JSON FILE!")
		panic(err)
	}
}

func InitLog() {
	os.Mkdir("log", 0764)
	logFile, err := os.OpenFile("./log/"+time.Now().Format("2006-01-02_15-04-05")+".txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0664)
	if err != nil {
		info.RedMsg("FAILED TO CREATE LOG FILE!")
		panic(err)
	}
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))
}
