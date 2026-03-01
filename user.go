package bilireq

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

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

func (api *Client) getCSRF() (_ string, err error) {
	defer err0.Then(&err, nil, nil)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	req := try.To1(http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s/csrf", api.host), nil))
	resp := try.To1(http.DefaultClient.Do(req))
	defer resp.Body.Close()
	csrf := try.To1(io.ReadAll(resp.Body))
	return string(csrf), nil
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
