package oa

import (
	"encoding/json"
	"net/url"

	"github.com/shenghui0779/gochat/public"
	"github.com/tidwall/gjson"
)

// TemplateInfo 模板信息
type TemplateInfo struct {
	TemplateID      string `json:"template_id"`
	Title           string `json:"title"`
	PrimaryIndustry string `json:"primary_industry"`
	DeputyIndustry  string `json:"deputy_industry"`
	Content         string `json:"content"`
	Example         string `json:"example"`
}

// GetTemplateList 获取模板列表
func GetTemplateList(dest *[]TemplateInfo) public.Action {
	return public.NewOpenGetAPI(TemplateListURL, url.Values{}, func(resp []byte) error {
		r := gjson.GetBytes(resp, "template_list")

		return json.Unmarshal([]byte(r.Raw), dest)
	})
}

// DeleteTemplate 删除模板
func DeleteTemplate(templateID string) public.Action {
	return public.NewOpenPostAPI(TemplateDeleteURL, url.Values{}, public.NewPostBody(func() ([]byte, error) {
		return json.Marshal(public.X{"template_id": templateID})
	}), nil)
}

// MessageBody 消息内容体
type MessageBody map[string]map[string]string

// TemplateMessage 公众号模板消息
type TemplateMessage struct {
	OpenID      string      // 接收者（用户）的 openid
	TemplateID  string      // 模板ID
	RedirectURL string      // 模板跳转链接（海外帐号没有跳转能力）
	MPAppID     string      // 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏）
	MPPagePath  string      // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar），要求该小程序已发布，暂不支持小游戏
	Data        MessageBody // 模板内容，格式形如：{"key1":{"value":"V","color":"#"},"key2":{"value": "V","color":"#"}}
}

// SendTemplateMessage 发送模板消息
func SendTemplateMessage(msg *TemplateMessage) public.Action {
	return public.NewOpenPostAPI(TemplateMessageSendURL, url.Values{}, public.NewPostBody(func() ([]byte, error) {
		params := public.X{
			"touser":      msg.OpenID,
			"template_id": msg.TemplateID,
			"data":        msg.Data,
		}

		if msg.RedirectURL != "" {
			params["url"] = msg.RedirectURL
		}

		if msg.MPAppID != "" {
			params["miniprogram"] = map[string]string{
				"appid":    msg.MPAppID,
				"pagepath": msg.MPPagePath,
			}
		}

		return json.Marshal(params)
	}), nil)
}

// SendSubscribeMessage 发送订阅消息
func SendSubscribeMessage(scene, title string, msg *TemplateMessage) public.Action {
	return public.NewOpenPostAPI(SubscribeMessageSendURL, url.Values{}, public.NewPostBody(func() ([]byte, error) {
		params := public.X{
			"scene":       scene,
			"title":       title,
			"touser":      msg.OpenID,
			"template_id": msg.TemplateID,
			"data":        msg.Data,
		}

		if msg.RedirectURL != "" {
			params["url"] = msg.RedirectURL
		}

		if msg.MPAppID != "" {
			params["miniprogram"] = map[string]string{
				"appid":    msg.MPAppID,
				"pagepath": msg.MPPagePath,
			}
		}

		return json.Marshal(params)
	}), nil)
}
