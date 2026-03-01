package bilireq

import (
	"strings"
	"time"

	"github.com/shynome/err0/try"
)

type BilibiliTime time.Time

const BilibiliTimeLayout = "2006-01-02 15:04:05"
const BilibiliTimeEmpty = "0000-00-00 00:00:00"

var BilibiliTimeLoc = try.To1(time.LoadLocation("Asia/Shanghai"))

func (c *BilibiliTime) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" || value == BilibiliTimeEmpty {
		return nil
	}

	t, err := time.ParseInLocation(BilibiliTimeLayout, value, BilibiliTimeLoc) //parse time
	if err != nil {
		return err
	}
	*c = BilibiliTime(t) //set result using the pointer
	return nil
}
