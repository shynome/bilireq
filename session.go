package bilireq

import "github.com/google/go-querystring/query"

type SessionGetParams struct {
	Talker
	ClientInfo
}

func (api *Client) Session(params SessionGetParams) (resp Response[MsgSession], err error) {
	p, err := query.Values(params)
	if err != nil {
		return
	}
	_, err = api.client.R().
		SetQueryParamsFromValues(p).
		SetResult(&resp).
		Get("https://api.vc.bilibili.com/session_svr/v1/session_svr/session_detail")
	return
}
