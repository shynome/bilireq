package bilireq_test

import (
	"testing"

	"github.com/shynome/bilireq"
)

func TestSessions(t *testing.T) {
	resp, err := bclient.Sessions(bilireq.SessionsGetParams{
		SessionType: bilireq.SessionTypeAll,
		GroupFold:   bilireq.IntTrue,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
