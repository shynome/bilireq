package bilireq_test

import (
	"testing"

	"github.com/shynome/err0/try"
)

func TestDynamicHistoryGenerator(t *testing.T) {

	hg := bclient.DynamicHistoryGenerator("")

	for hg.Next() {
		val := hg.Value()
		t.Log(val)
	}

}

func TestLiveAll(t *testing.T) {

	feeds := try.To1(bclient.LiveFeedAll())
	t.Log(feeds)

}

func TestLiveInfo(t *testing.T) {

	info := try.To1(bclient.LiveRoomInfo("898286"))
	t.Log(info)

}
