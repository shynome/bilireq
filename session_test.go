package bilireq_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/shynome/bilireq"
)

func TestSession(t *testing.T) {
	idStr := os.Getenv("MSG_RECEIVER")
	id, _ := strconv.Atoi(idStr)
	resp, err := bclient.Session(bilireq.SessionGetParams{
		Talker: bilireq.Talker{ID: int64(id), SessionType: bilireq.MsgSessionTypeUser},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
