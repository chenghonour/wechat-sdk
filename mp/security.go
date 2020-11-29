package mp

import (
	"encoding/json"
	"io/ioutil"
	"net/url"

	"github.com/shenghui0779/gochat/public"
	"github.com/tidwall/gjson"
)

// SecCheckMediaType 检测的素材类型
type SecCheckMediaType int

// 微信支持的素材类型
var (
	SecCheckMediaAudio = 1 // 音频
	SecCheckMediaImage = 2 // 图片
)

// ImageSecCheck 校验一张图片是否含有违法违规内容
func ImageSecCheck(filename string) public.Action {
	return public.NewOpenUploadAPI(ImageSecCheckURL, url.Values{}, public.NewUploadBody("media", filename, func() ([]byte, error) {
		return ioutil.ReadFile(filename)
	}), nil)
}

// MediaCheckAsyncInfo 任务id，用于匹配异步推送结果
type MediaCheckAsyncInfo struct {
	TraceID string
}

// MediaCheckAsync 异步校验图片/音频是否含有违法违规内容
func MediaCheckAsync(mediaType SecCheckMediaType, mediaURL string, dest *MediaCheckAsyncInfo) public.Action {
	return public.NewOpenPostAPI(MediaCheckAsyncURL, url.Values{}, public.NewPostBody(func() ([]byte, error) {
		return json.Marshal(public.X{
			"media_type": mediaType,
			"media_url":  mediaURL,
		})
	}), func(resp []byte) error {
		dest.TraceID = gjson.GetBytes(resp, "trace_id").String()

		return nil
	})
}

// MsgSecCheck 检查一段文本是否含有违法违规内容
func MsgSecCheck(content string) public.Action {
	return public.NewOpenPostAPI(MsgSecCheckURL, url.Values{}, public.NewPostBody(func() ([]byte, error) {
		return json.Marshal(public.X{
			"content": content,
		})
	}), nil)
}
