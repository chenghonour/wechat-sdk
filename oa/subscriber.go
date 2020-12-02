package oa

import (
	"encoding/json"
	"net/url"

	"github.com/shenghui0779/gochat/wx"
	"github.com/tidwall/gjson"
)

// MaxSubscriberListCount 公众号订阅者列表的最大数目
const MaxSubscriberListCount = 10000

// SubscriberInfo 微信公众号订阅者信息
type SubscriberInfo struct {
	Subscribe      int     `json:"subscribe"`
	OpenID         string  `json:"openid"`
	NickName       string  `json:"nickname"`
	Sex            int     `json:"sex"`
	Language       string  `json:"language"`
	City           string  `json:"city"`
	Province       string  `json:"province"`
	Country        string  `json:"country"`
	AvatarURL      string  `json:"headimgurl"`
	SubscribeTime  int64   `json:"subscribe_time"`
	UnionID        string  `json:"unionid"`
	Remark         string  `json:"remark"`
	GroupID        int64   `json:"groupid"`
	TagidList      []int64 `json:"tagid_list"`
	SubscribeScene string  `json:"subscribe_scene"`
	QRScene        int64   `json:"qr_scene"`
	QRSceneStr     string  `json:"qr_scene_str"`
}

// SubscriberList 微信公众号订阅者列表
type SubscriberList struct {
	Total      int                `json:"total"`
	Count      int                `json:"count"`
	Data       SubscriberListData `json:"data"`
	NextOpenID string             `json:"next_openid"`
}

// SubscriberListData 微信公众号订阅者列表数据
type SubscriberListData struct {
	OpenID []string `json:"openid"`
}

// GetSubscriberInfo 获取微信公众号订阅者信息
func GetSubscriberInfo(openid string, dest *SubscriberInfo) wx.Action {
	query := url.Values{}

	query.Set("openid", openid)
	query.Set("lang", "zh_CN")

	return wx.NewOpenGetAPI(SubscriberGetURL, query, func(resp []byte) error {
		return json.Unmarshal(resp, dest)
	})
}

// BatchGetSubscriberInfo 批量获取微信公众号订阅者信息
func BatchGetSubscriberInfo(openids []string, dest *[]SubscriberInfo) wx.Action {
	return wx.NewOpenPostAPI(SubscriberBatchGetURL, url.Values{}, wx.NewPostBody(func() ([]byte, error) {
		userList := make([]map[string]string, 0, len(openids))

		for _, v := range openids {
			userList = append(userList, map[string]string{
				"openid": v,
				"lang":   "zh_CN",
			})
		}

		return json.Marshal(map[string][]map[string]string{"user_list": userList})
	}), func(resp []byte) error {
		r := gjson.GetBytes(resp, "user_info_list")

		return json.Unmarshal([]byte(r.Raw), dest)
	})
}

// GetSubscriberList 获取微信公众号订阅者列表
func GetSubscriberList(nextOpenID string, dest *SubscriberList) wx.Action {
	query := url.Values{}

	query.Set("next_openid", nextOpenID)

	return wx.NewOpenGetAPI(SubscriberListURL, query, func(resp []byte) error {
		return json.Unmarshal(resp, dest)
	})
}

// GetBlackList 获取用户黑名单列表
func GetBlackList(beginOpenID string, dest *SubscriberList) wx.Action {
	return wx.NewOpenPostAPI(BlackListGetURL, url.Values{}, wx.NewPostBody(func() ([]byte, error) {
		return json.Marshal(wx.X{
			"begin_openid": beginOpenID,
		})
	}), func(resp []byte) error {
		return json.Unmarshal(resp, dest)
	})
}

// BatchBlackList 拉黑用户
func BatchBlackList(openids ...string) wx.Action {
	return wx.NewOpenPostAPI(BatchBlackListURL, url.Values{}, wx.NewPostBody(func() ([]byte, error) {
		return json.Marshal(wx.X{
			"openid_list": openids,
		})
	}), nil)
}

// BatchUnBlackList 取消拉黑用户
func BatchUnBlackList(openids ...string) wx.Action {
	return wx.NewOpenPostAPI(BatchUnBlackListURL, url.Values{}, wx.NewPostBody(func() ([]byte, error) {
		return json.Marshal(wx.X{
			"openid_list": openids,
		})
	}), nil)
}

// SetUserRemark 设置用户备注名（该接口暂时开放给微信认证的服务号）
func SetUserRemark(openid, remark string) wx.Action {
	return wx.NewOpenPostAPI(UserRemarkSetURL, url.Values{}, wx.NewPostBody(func() ([]byte, error) {
		return json.Marshal(wx.X{
			"openid": openid,
			"remark": remark,
		})
	}), nil)
}