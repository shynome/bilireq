package bilireq_test

import (
	"testing"

	"github.com/shynome/bilireq"
	"github.com/shynome/err0/try"
)

func TestUnreadMsgs(t *testing.T) {
	resp := try.To1(bclient.Session(bilireq.SessionGetParams{
		Talker: bilireq.TalkerUser(msgReceiver),
	}))

	msgs, err := bclient.UnreadMsgs(resp.Data)
	if err != nil {
		t.Error(err)
	}
	t.Log(msgs)
}

func TestYieldUnreadMsgs(t *testing.T) {
	f := bclient.YieldUnreadMsgs(0)
	for msgs := range f {
		t.Log(msgs)
	}
}
