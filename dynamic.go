package bilireq

import (
	"fmt"
)

type DynamicNewList struct {
	NewNum        int32   `json:"new_num,omitempty"`
	ExistGap      int32   `json:"exist_gap,omitempty"`
	UpdateNum     int32   `json:"update_num,omitempty"`
	OpenRcmd      int32   `json:"open_rcmd,omitempty"`
	Cards         []*Card `json:"cards,omitempty"`
	MaxDynamicId  int64   `json:"max_dynamic_id,omitempty"`
	HistoryOffset int64   `json:"history_offset,omitempty"`
}

func (api *Client) DynamicNew() (resp Response[DynamicNewList], err error) {
	_, err = api.client.R().
		SetQueryParams(map[string]string{
			"platform":  "web",
			"from":      "weball",
			"type_list": "268435455",
		}).
		SetResult(&resp).
		Get("https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/dynamic_new")
	return
}

type DynamicHistoryList struct {
	Cards      []*Card `json:"cards,omitempty"`
	HasMore    int32   `json:"has_more,omitempty"`
	NextOffset int64   `json:"next_offset,omitempty"`
}

func (api *Client) DynamicHistory(offset string) (resp Response[DynamicHistoryList], err error) {
	_, err = api.client.R().
		SetQueryParams(map[string]string{
			"offset_dynamic_id": offset,
			"type_list":         "268435455",
		}).
		SetResult(&resp).
		Get("https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/dynamic_history")
	return
}

type DynamicHistory struct {
	err   error
	value []*Card

	api *Client

	currentOffset string
	nextOffset    *string
}

func (api *Client) DynamicHistoryGenerator(offset string) *DynamicHistory {
	return &DynamicHistory{
		api: api,

		currentOffset: offset,
		nextOffset:    &offset,
	}
}

func (d *DynamicHistory) Next() (hasNext bool) {
	if d.err != nil || d.nextOffset == nil {
		return false
	}

	if d.currentOffset == "" {
		if resp, err := d.api.DynamicNew(); err != nil {
			d.err = err
			return
		} else {
			d.next(fmt.Sprint(resp.Data.HistoryOffset), resp.Data.Cards)
		}
	} else {
		if resp, err := d.api.DynamicHistory(*d.nextOffset); err != nil {
			d.err = err
			return
		} else {
			d.next(fmt.Sprint(resp.Data.NextOffset), resp.Data.Cards)
		}
	}
	d.currentOffset = *d.nextOffset
	return true
}

func (d *DynamicHistory) next(nextOffset string, value []*Card) {
	d.nextOffset = &nextOffset
	d.value = value
}

func (d *DynamicHistory) Value() []*Card { return d.value }
func (d *DynamicHistory) Error() error   { return d.err }
