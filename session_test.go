package bilireq_test

import (
	"testing"

	"github.com/shynome/bilireq"
)

func TestSession(t *testing.T) {

	resp, err := bclient.Session(bilireq.SessionGetParams{
		Talker: bilireq.TalkerUser(msgReceiver),
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
