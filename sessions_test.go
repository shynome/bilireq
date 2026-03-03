package bilireq_test

import (
	"log"
	"testing"

	"github.com/shynome/bilireq"
)

func TestSessions(t *testing.T) {
	resp, err := bclient.Sessions(bilireq.SessionsGetParams{
		SessionType: bilireq.SessionTypeUser,
		GroupFold:   bilireq.IntTrue,
		SortRule:    2,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

func TestYieldSessions(t *testing.T) {
	f := bclient.YieldSessions(bilireq.SessionsGetParams{
		SessionType: bilireq.SessionTypeUser,
		GroupFold:   bilireq.IntTrue,
		SortRule:    2,
	})
	bclient.SetDebug(false)
	i := 0
	for ss := range f {
		for _, s := range ss.List {
			i++
			log.Println("talker", s.Talker.ID)
		}
	}
	log.Println("count", i)
}
