package jsapi

import (
	"encoding/json"

	"github.com/chenghonour/wechat-sdk/urls"
	"github.com/chenghonour/wechat-sdk/wx"
)

type ResultTicket struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

func GetQYTicket(result *ResultTicket) wx.Action {
	return wx.NewGetAction(urls.CorpQYTicket,
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}

func GetAgentTicket(result *ResultTicket) wx.Action {
	return wx.NewGetAction(urls.CorpAgentTicket,
		wx.WithQuery("type", "agent_config"),
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}
