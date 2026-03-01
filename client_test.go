package bilireq_test

import (
	"testing"

	"github.com/shynome/bilireq"
)

var bclient *bilireq.Client

func TestMain(m *testing.M) {
	bclient = bilireq.New("192.168.211.2")
	m.Run()
}
