package bilireq

import "github.com/google/go-querystring/query"

type MsgUnreadParams struct {
	Type             MsgUnreadType `url:"unread_type,omitempty"`        // 未读类型
	ShowUnfollowList IntBool       `url:"show_unfollow_list,omitempty"` // 是否返回未关注人推送消息数
	ShowDustbin      IntBool       `url:"show_dustbin,omitempty"`       // 是否返回未关注人推送消息数
}

type MsgUnreadType int64

const (
	MsgUnreadTypeAll          MsgUnreadType = 0 // 所有
	MsgUnreadTypeFollowOnly   MsgUnreadType = 1 // 仅已关注
	MsgUnreadTypeUnfollowOnly MsgUnreadType = 2 // 仅未关注
	MsgUnreadTypeDustbinOnly  MsgUnreadType = 3 // 仅被拦截 (须同时设置参数 show_dustbin=1)
)

type MsgUnread struct {
	Unfollow     int64 `json:"unfollow_unread"`   // 未读未关注用户私信数
	Follow       int64 `json:"follow_unread"`     // 未读已关注用户私信数
	UnfollowPush int64 `json:"unfollow_push_msg"` // 未读未关注用户推送消息数
	MsgUnreadDustbin
	MsgUnreadBiz
	Custom int64 `json:"custom_unread"` // 未读客服消息数
}

type MsgUnreadBiz struct {
	Unfollow int64 `json:"biz_msg_unfollow_unread"` // 未读未关注用户通知数
	Follow   int64 `json:"biz_msg_follow_unread"`   // 未读已关注用户通知数
}

type MsgUnreadDustbin struct {
	PushMsg int64 `json:"dustbin_push_msg"` // 未读被拦截的推送消息数
	Unread  int64 `json:"dustbin_unread"`   // 未读被拦截的私信数
}

func (api *Client) MsgUnread(params MsgUnreadParams) (resp Response[MsgUnread], err error) {
	p, err := query.Values(params)
	if err != nil {
		return
	}
	_, err = api.client.R().
		SetQueryParamsFromValues(p).
		SetResult(&resp).
		Get("https://api.vc.bilibili.com/session_svr/v1/session_svr/single_unread")
	return
}
