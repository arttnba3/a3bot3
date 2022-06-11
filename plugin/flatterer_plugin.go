package plugin

import (
	"a3bot3/api"
	"a3bot3/event"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type FlattererPlugin struct {
	PluginInfo
}

type FlattererRespData struct {
	Comment string `json:"comment"`
}

type FlattererResp struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data FlattererRespData `json:"data"`
}

func getFlattererText() string {
	var fResp FlattererResp

	resp, err := http.Get("https://api.muxiaoguo.cn/api/tiangourj?api_key=")
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	respData, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(respData, &fResp)
	if err != nil {
		return err.Error()
	}

	return fResp.Data.Comment
}

func (p *FlattererPlugin) PrivateMsgHandler(bot api.BotAPI, privateEvent event.PrivateEvent, messages []string) int {
	if p.MatchCommand(messages[0]) {
		bot.SendPrivateMsg(privateEvent.UserID, getFlattererText(), false)
		return MESSAGE_BLOCK
	} else {
		return MESSAGE_IGNORE
	}
}

func (p *FlattererPlugin) GroupMsgHandler(bot api.BotAPI, groupEvent event.GroupEvent, messages []string) int {
	if p.MatchCommand(messages[0]) {
		bot.SendGroupMsg(groupEvent.UserID, getFlattererText(), false)
		return MESSAGE_BLOCK
	} else {
		return MESSAGE_IGNORE
	}
}
