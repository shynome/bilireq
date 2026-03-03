package bilireq_test

import (
	"testing"
)

func TestMsgSend(t *testing.T) {
	resp, err := bclient.MsgSend2User(msgReceiver, "hello world")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}
