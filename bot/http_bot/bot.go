package http_bot

import (
	"a3bot3/api"
	"a3bot3/bot"
	"a3bot3/event"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type MessageHandler func(perEvent event.Event)

var messageHandler MessageHandler

type HTTPBot struct {
	api.Bot
}

// SendRequest
/*
 * @Description
 */
func (bot *HTTPBot) SendRequest(i interface{}) {
	data := i.(api.Action)
	postUrl := fmt.Sprintf("http://%v:%v/%v", bot.Host, bot.Port, data.Action)
	byteData, err := json.Marshal(data.Params)
	if err != nil {
		log.Fatalln("Error occur: " + err.Error() + "\nCheck your code!")
		return
	}

	resp, err := http.Post(postUrl,
		"application/json;charset=UTF-8",
		bytes.NewReader(byteData))
	if err != nil {
		log.Fatalln("Error occur: " + err.Error() + "\nCheck cq-http!")
		return
	}
	defer resp.Body.Close()

	respData, _ := ioutil.ReadAll(resp.Body)

	bot.Responses.Store(data.UUID, respData)
}

// GetResponse
/*
 * @Description
 */
func (bot *HTTPBot) GetResponse(uuid string) ([]byte, error) {
	data, ok := bot.Responses.LoadAndDelete(uuid)
	if ok {
		return data.([]byte), nil
	} else {
		return nil, errors.New("invalid uuid")
	}
}

// CQHTTPListener
/*
 *
 */
func CQHTTPListener(host string, port int) {
	fmt.Println("serving:" + host + ":" + strconv.Itoa(port))
	http.HandleFunc("/", CQHTTPServer)
	err := http.ListenAndServe(host+":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func CQHTTPServer(resp http.ResponseWriter, request *http.Request) {
	var perEvent event.Event

	data, _ := ioutil.ReadAll(request.Body)
	err := json.Unmarshal(data, &perEvent)
	if err != nil {
		log.Println("Invalid data from cq-http:", data)
		return
	}

	bot.MessageHandler(perEvent)
}
