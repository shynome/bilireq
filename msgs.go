package bilireq

import "github.com/google/go-querystring/query"

type MsgsGetParams struct {
	Talker
}

func (api *Client) Msgs(params MsgsGetParams) (resp Response[MsgSession], err error) {
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

type Talker struct {
	ID          int64          `url:"talker_id"`    // 聊天对象的id. session_type 为 1 时表示用户 mid，为 2 时表示粉丝团 id
	SessionType MsgSessionType `url:"session_type"` // 聊天对象的类型
}
