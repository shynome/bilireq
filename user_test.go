package bilireq_test

import (
	"testing"

	"github.com/shynome/bilireq"
	"github.com/shynome/err0/try"
)

func TestUser(t *testing.T) {
	user := try.To1(bclient.UserInfo())
	t.Log(user.Data.Mid)
}

func TestRelationModify(t *testing.T) {
	resp, err := bclient.RelationModify("6660627", bilireq.ActSub)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}
