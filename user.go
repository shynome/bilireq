package bilireq

import (
	"fmt"

	"github.com/shynome/err0"
	"github.com/shynome/err0/try"
)

type UserInfo struct {
	IsLogin bool   `json:"isLogin,omitempty"`
	Mid     int64  `json:"mid,omitempty"`
	Uname   string `json:"uname,omitempty"`
}

func (api *Client) UserInfo() (resp Response[UserInfo], err error) {
	_, err = api.client.R().
		SetResult(&resp).
		Get("https://api.bilibili.com/x/web-interface/nav")
	return
}

func (api *Client) RelationModify(mid string, act ActType) (resp Response[any], err error) {
	defer err0.Then(&err, nil, nil)
	csrf := try.To1(api.getCSRF())
	_, err = api.client.R().
		SetFormData(map[string]string{
			"fid":  mid,
			"act":  fmt.Sprint(act),
			"csrf": csrf,
		}).
		SetResult(&resp).
		Post("https://api.bilibili.com/x/relation/modify")
	return
}

type ActType int

const (
	ActSub ActType = 1 + iota
	ActUnsub
	ActHiddenSub
	ActUnhiddenSub
	ActBlock
	ActUnblock
	ActRemoveFollower
)

type SpaceInfo struct {
	Mid   int64  `json:"mid,omitempty"`
	Name  string `json:"name,omitempty"`
	Level int32  `json:"level,omitempty"`
}

func (api *Client) SpaceInfo(uid string) (resp Response[SpaceInfo], err error) {
	_, err = api.client.R().
		SetQueryParam("mid", uid).
		SetQueryParam("platform", "web").
		SetResult(&resp).
		Get("https://api.bilibili.com/x/space/acc/info")
	return
}

type AttentionList struct {
	List []int64 `json:"list,omitempty"`
}

func (api *Client) AttentionList() (resp Response[AttentionList], err error) {
	_, err = api.client.R().
		SetResult(&resp).
		Get("https://api.vc.bilibili.com/vc/feed/v1/feed/get_attention_list")
	return
}
