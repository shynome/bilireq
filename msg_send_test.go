package bilireq_test

import (
	"os"
	"testing"
)

func TestMsgSend(t *testing.T) {
	resp, err := bclient.MsgSend2User(os.Getenv("MSG_RECEIVER"), "hello world")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}
