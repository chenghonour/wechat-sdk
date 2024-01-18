package linkedcorp

import (
	"encoding/json"

	"github.com/chenghonour/wechat-sdk/urls"
	"github.com/chenghonour/wechat-sdk/wx"
)

type ResultAgentPermList struct {
	UserIDs       []string `json:"userids"`
	DepartmentIDs []string `json:"department_ids"`
}

// ListAgentPerm 获取应用的可见范围
func ListAgentPerm(result *ResultAgentPermList) wx.Action {
	return wx.NewPostAction(urls.CorpLinkedcorpPermList,
		wx.WithDecode(func(b []byte) error {
			return json.Unmarshal(b, result)
		}),
	)
}
