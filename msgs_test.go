package bilireq_test

import (
	"log"
	"testing"

	"github.com/shynome/bilireq"
)

func TestMsgs(t *testing.T) {
	resp, err := bclient.Msgs(bilireq.MsgsGetParams{
		Talker: bilireq.TalkerUser(msgReceiver),
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
	for _, m := range resp.Data.Messages {
		var s = m.Content
		log.Println("[", m.MsgSeqNo, "]", m.Sender, ":", s)
	}
}

func TestYield(t *testing.T) {
	f := bclient.YieldMsgs(bilireq.MsgsGetParams{
		Talker: bilireq.TalkerUser(msgReceiver),
	})
	for minSeqno, v := range f {
		log.Println("mmmmmmmmmmmmmmmin", minSeqno)
		for _, m := range v.Messages {
			var s = m.Content
			log.Println("[", m.MsgSeqNo, "]", m.Sender, ":", s)
		}
	}
	log.Println("8")
}
