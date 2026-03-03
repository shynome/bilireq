package bilireq

import (
	"encoding/json"

	"github.com/google/go-querystring/query"
)

type MsgsGetParams struct {
	Talker
	Size           int64 `url:"size,omitempty"`             // 返回消息数量. 默认为 0，最大为 2000. 当本参数为 0 或不存在时，只返回系统提示
	BeginSeqNo     int64 `url:"begin_seqno,omitempty"`      // 开始的序列号. 提供本参数时返回以本序列号开始（不包括本序列号）的消息
	EndSeqNo       int64 `url:"end_seqno,omitempty"`        // 结束的序列号. 提供本参数时返回以本序列号结束（不包括本序列号）的消息
	SenderDeviceID int64 `url:"sender_device_id,omitempty"` // 发送者设备. 默认为 1
	ClientInfo
}

func (api *Client) Msgs(params MsgsGetParams) (resp Response[SessionMsgs], err error) {
	if params.Size == 0 {
		params.Size = 20
	}
	params.ClientInfo.fill()
	p, err := query.Values(params)
	if err != nil {
		return
	}
	_, err = api.client.R().
		SetQueryParamsFromValues(p).
		SetResult(&resp).
		Get("https://api.vc.bilibili.com/svr_sync/v1/svr_sync/fetch_session_msgs")
	return
}

type SessionMsgs struct {
	Messages   []PrivateMsgItem `json:"messages"`
	HasMore    int64            `json:"has_more"`
	MinSeqNo   int64            `json:"min_seqno"`
	MaxSeqNo   int64            `json:"max_seqno"`
	EmojiInfos json.RawMessage  `json:"e_infos"`
}
