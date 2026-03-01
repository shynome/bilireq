package bilireq

import (
	"fmt"

	"github.com/shynome/err0"
	"github.com/shynome/err0/try"
)

type LiveFeedList = Page[LiveInfo]

const LiveFeedPageSize = 20

func (api *Client) LiveFeedList(page int) (resp Response[LiveFeedList], err error) {
	_, err = api.client.R().
		SetQueryParams(map[string]string{
			"page":     fmt.Sprint(page),
			"pageSize": fmt.Sprint(LiveFeedPageSize),
		}).
		SetResult(&resp).
		Get("https://api.live.bilibili.com/xlive/web-ucenter/v1/xfetter/FeedList")
	return
}

type LiveInfo struct {
	Cover        string `json:"cover,omitempty"`
	Face         string `json:"face,omitempty"`
	Uname        string `json:"uname,omitempty"`
	Title        string `json:"title,omitempty"`
	Roomid       int64  `json:"roomid,omitempty"`
	Pic          string `json:"pic,omitempty"`
	Online       int64  `json:"online,omitempty"`
	Link         string `json:"link,omitempty"`
	Uid          int64  `json:"uid,omitempty"`
	ParentAreaId int64  `json:"parent_area_id,omitempty"`
	AreaId       int64  `json:"area_id,omitempty"`
}

func (api *Client) LiveFeedAll() (feeds []LiveInfo, err error) {
	defer err0.Then(&err, nil, nil)
	for i := 1; true; i++ {
		resp := try.To1(api.LiveFeedList(i))
		nfeeds := resp.Data.List
		feeds = append(feeds, nfeeds...)
		if resp.Data.Results == 0 {
			break
		}
		if len(nfeeds) < 20 {
			break
		}
	}
	return
}

type LiveStatus int

const (
	LiveOff       LiveStatus = 0 //未开播
	LiveOnline    LiveStatus = 1 //直播中
	LiveVideoLoop LiveStatus = 2 //视频轮播
)

type LiveRoomInfo struct {
	LiveTime   BilibiliTime `json:"live_time"`
	LiveStatus LiveStatus   `json:"live_status"`
}

func (api *Client) LiveRoomInfo(id string) (resp Response[LiveRoomInfo], err error) {
	_, err = api.client.R().
		SetQueryParams(map[string]string{
			"id": id,
		}).
		SetResult(&resp).
		Get("https://api.live.bilibili.com/room/v1/Room/get_info")
	return
}
