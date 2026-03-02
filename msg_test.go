package bilireq_test

import (
	"testing"

	"github.com/shynome/bilireq"
)

func TestMsg(t *testing.T) {
	resp, err := bclient.MsgUnread(bilireq.MsgUnreadParams{})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
