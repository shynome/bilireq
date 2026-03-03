package bilireq

import (
	"encoding/json"

	"github.com/google/go-querystring/query"
)

type SessionsGetParams struct {
	SessionType  SessionType `url:"session_type"`            // 会话类型
	GroupFold    IntBool     `url:"group_fold,omitempty"`    // 是否折叠粉丝团消息
	UnfollowFold IntBool     `url:"unfollow_fold,omitempty"` // 是否折叠未关注人消息
	SortRule     int64       `url:"sort_rule,omitempty"`     // 仅当 session_type 不为 4、7 时有效. 1、2：按会话时间逆向排序; 3：按已读时间逆向排序; 其他：用户与系统按会话时间逆向排序，粉丝团按加入时间正向排序
	BeginTs      int64       `url:"begin_ts,omitempty"`      // 起始时间. 微秒级时间戳
	EndTs        int64       `url:"end_ts,omitempty"`        // 终止时间. 微秒级时间戳
	Size         int64       `url:"size,omitempty"`          // 返回的会话数. 默认为 20，最大为 100
	ClientInfo
}

func (api *Client) Sessions(params SessionsGetParams) (resp Response[SessionList], err error) {
	p, err := query.Values(params)
	if err != nil {
		return
	}
	_, err = api.client.R().
		SetQueryParamsFromValues(p).
		SetResult(&resp).
		Get("https://api.vc.bilibili.com/session_svr/v1/session_svr/get_sessions")
	return
}

func (api *Client) YieldSessions(params SessionsGetParams) func(func(SessionList) bool) {
	return func(yield func(SessionList) bool) {
		next := params.EndTs
		for {
			params.EndTs = next
			resp, err := api.Sessions(params)
			if err != nil {
				return
			}
			d := resp.Data
			if !yield(d) {
				return
			}
			if d.HasMore == 0 {
				return
			}
			if len(d.List) == 0 {
				return
			}
			next = d.List[len(d.List)-1].SessionTs
		}
	}
}

type SessionList struct {
	List                []MsgSession    `json:"session_list"`          // 会话列表
	HasMore             IntBool         `json:"has_more"`              // 是否有更多会话
	AntiDistrubCleaning bool            `json:"anti_distrub_cleaning"` // 是否开启了“一键防骚扰”功能
	IsAddressListEmpty  int64           `json:"is_address_list_empty"` // 作用尚不明确
	SystemMsg           json.RawMessage `json:"system_msg"`            // 系统会话列表
	ShowLevel           bool            `json:"show_level"`            // 是否在会话列表中显示用户等级. 目前恒为 true
}

// 会话类型
type SessionType int64

const (
	_                    SessionType = iota //
	SessionTypeUser                         // 用户与系统
	SessionTypeUnfollow                     // 未关注人
	SessionTypeFans                         // 粉丝团
	SessionTypeAll                          // 所有
	SessionTypeBiz                          // 被拦截
	SessionType6                            // 花火商单
	SessionTypeAllSystem                    // 所有系统消息
	SessionType8                            // 陌生人（与 “未关注人” 不同，不包含官方消息）
	SessionType9                            // 关注的人与系统
)
