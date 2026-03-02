package bilireq

import (
	"encoding/json"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/shynome/err0"
	"github.com/shynome/err0/try"
)

type MsgSendParams struct {
	Sender       string         `url:"msg[sender_uid]"`                 //
	Receiver     string         `url:"msg[receiver_id]"`                //
	ReceiverType MsgSessionType `url:"msg[receiver_type]"`              // 接收者类型
	MsgType      MsgSendType    `url:"msg[msg_type]"`                   // 消息类型*. 此接口仅支持传入 1、2 或 5
	MsgStatus    int64          `url:"msg[msg_status],omitempty"`       //
	FaceVersion  int64          `url:"msg[new_face_version],omitempty"` //
	Timestamp    int64          `url:"msg[timestamp]"`                  //
	ContentStr   string         `url:"msg[content]"`                    //
	Content      any            `url:"-"`                               //
	CSRF         string         `url:"csrf"`                            //
	DeviceID     string         `url:"msg[dev_id]"`                     //
	ClientInfo
}

// 说是需要 wbi 签名, 但其实不需要, 只需要 csrf 就够了
func (api *Client) MsgSend(params MsgSendParams) (resp Response[json.RawMessage], err error) {
	defer err0.Then(&err, nil, nil)
	ss := try.To1(api.getCSRF3())
	params.CSRF = ss[0]
	params.Sender = ss[1]
	params.DeviceID = ss[2]
	if params.ContentStr == "" {
		content := try.To1(json.Marshal(params.Content))
		params.ContentStr = string(content)
	}
	params.Timestamp = time.Now().Unix()
	if params.MobiApp == "" {
		params.MobiApp = "web"
	}
	p := try.To1(query.Values(params))
	_, err = api.client.R().
		SetFormDataFromValues(p).
		SetResult(&resp).
		Post("https://api.vc.bilibili.com/web_im/v1/web_im/send_msg")
	return
}

func (api *Client) MsgSend2User(uid string, msg string) (resp Response[json.RawMessage], err error) {
	p := MsgSendParams{
		Receiver:     uid,
		ReceiverType: MsgSessionTypeUser,
		MsgType:      MsgSendTypeText,
		Content:      MsgContent{Content: msg},
	}
	return api.MsgSend(p)
}

type MsgContent struct {
	Content string `json:"content"`
}

// 私信消息类型
type MsgSendType int64

const (
	MsgSendTypeText   MsgSendType = 1 // 文字消息
	MsgSendTypeImg    MsgSendType = 2 // 图片消息
	MsgSendTypeRevoke MsgSendType = 5 // 撤回消息
)
