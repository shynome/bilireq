package bilireq

import "time"

func (api *Client) UnreadMsgs(m MsgSession) ([]PrivateMsgItem, error) {
	q := MsgsGetParams{
		Talker:     m.Talker,
		Size:       m.UnreadCount,
		BeginSeqNo: m.AckSeqNo,
	}
	resp, err := api.Msgs(q)
	if err != nil {
		return nil, err
	}

	return resp.Data.Messages, nil
}

func (api *Client) YieldUnreadMsgs(d time.Duration) func(func(MsgSession, []PrivateMsgItem) bool) {
	if d == 0 {
		d = time.Second
	}
	return func(yield func(MsgSession, []PrivateMsgItem) bool) {
		q := SessionsGetParams{
			SessionType: SessionTypeUser,
			GroupFold:   IntTrue,
			SortRule:    2,
		}
		for {
			sessions := []MsgSession{}
			for ss := range api.YieldSessions(q) {
				meetReaded := false
				for _, s := range ss.List {
					if s.UnreadCount == 0 {
						meetReaded = true
						break
					}
					sessions = append(sessions, s)
				}
				if meetReaded {
					break
				}
			}
			for _, s := range sessions {
				msgs, err := api.UnreadMsgs(s)
				if err != nil {
					continue
				}
				if !yield(s, msgs) {
					return
				}
			}
			time.Sleep(d)
		}
	}
}
