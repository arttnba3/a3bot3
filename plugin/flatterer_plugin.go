package plugin

import (
	"a3bot3/api"
	"a3bot3/event"
	"a3bot3/tools"
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
	// re-parsing if necessary
	if messages == nil || len(messages) < 1 {
		messages = tools.AutoParser(privateEvent.Message)
	}

	if len(messages) < 1 || !p.MatchCommand(messages[0]) {
		return MESSAGE_IGNORE
	} else {
		bot.SendPrivateMsg(privateEvent.UserID, getFlattererText(), false)
		return MESSAGE_BLOCK
	}
}

func (p *FlattererPlugin) GroupMsgHandler(bot api.BotAPI, groupEvent event.GroupEvent, messages []string) int {
	// re-parsing if necessary
	if messages == nil || len(messages) < 1 {
		messages = tools.AutoParser(groupEvent.Message)
	}

	if len(messages) < 1 || !p.MatchCommand(messages[0]) {
		return MESSAGE_IGNORE
	} else {
		bot.SendGroupMsg(groupEvent.GroupID, getFlattererText(), false)
		return MESSAGE_BLOCK
	}
}
