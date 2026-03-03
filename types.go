package bilireq

type IntBool int64

const (
	IntTrue  IntBool = 1
	IntFalse IntBool = 0
)

type ClientInfo struct {
	BuildVersion int64  `url:"build,omitempty"`    // 客户端内部版本号
	MobiApp      string `url:"mobi_app,omitempty"` // 平台标识. 可为 web 等
}

func (c *ClientInfo) fill() {
	if c.MobiApp == "" {
		c.MobiApp = "web"
	}
}
