package bilireq

import (
	"encoding/json"

	"github.com/google/go-querystring/query"
)

type MsgsGetParams struct {
	Talker
	Size           int64  `url:"size,omitempty"`             // 返回消息数量. 默认为 0，最大为 2000. 当本参数为 0 或不存在时，只返回系统提示
	EndSeqNo       uint64 `url:"end_seqno,omitempty"`        // 最新的序列号开始. 提供本参数时返回以本序列号结束（不包括本序列号）的消息
	BeginSeqNo     uint64 `url:"begin_seqno,omitempty"`      // 到最老的序列号结束. 提供本参数时返回以本序列号开始（不包括本序列号）的消息
	SenderDeviceID int64  `url:"sender_device_id,omitempty"` // 发送者设备. 默认为 1
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

func (api *Client) YieldMsgs(params MsgsGetParams) func(yield func(uint64, *SessionMsgs) bool) {
	next := params.EndSeqNo
	return func(yield func(uint64, *SessionMsgs) bool) {
		for {
			params.EndSeqNo = next
			resp, err := api.Msgs(params)
			if err != nil {
				return
			}
			d := resp.Data
			next = d.MinSeqNo
			if !yield(next, &d) {
				return
			}
			if d.HasMore == 0 {
				return
			}
		}
	}
}

type SessionMsgs struct {
	Messages   []PrivateMsgItem `json:"messages"`
	HasMore    int64            `json:"has_more"`
	MinSeqNo   uint64           `json:"min_seqno"`
	MaxSeqNo   uint64           `json:"max_seqno"`
	EmojiInfos json.RawMessage  `json:"e_infos"`
}
