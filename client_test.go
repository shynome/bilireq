package bilireq_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/shynome/bilireq"
)

var bclient *bilireq.Client
var msgReceiver int64

func TestMain(m *testing.M) {
	bclient = bilireq.New("192.168.211.2")
	bclient.SetDebug(true)
	idStr := os.Getenv("MSG_RECEIVER")
	id, _ := strconv.Atoi(idStr)
	msgReceiver = int64(id)
	m.Run()
}
