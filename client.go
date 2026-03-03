package bilireq

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/shynome/err0"
)

type Client struct {
	client *resty.Client
	host   string
}

var (
	ErrBilibili   = errors.New("bilibili resp data error")
	ErrUnkownCode = fmt.Errorf("%w. unkown code error", ErrBilibili)
	ErrNotLogin   = fmt.Errorf("%w. user is not login", ErrBilibili)
	ErrRespFormat = fmt.Errorf("%w. resp unmarshal json failed", ErrBilibili)
)

func New(host string) (_ *Client) {
	client := resty.New().
		SetTimeout(5 * time.Second).
		SetProxy(fmt.Sprintf("http://%s:1080", host)).
		SetTLSClientConfig(&tls.Config{
			InsecureSkipVerify: true,
		})

	api := &Client{
		client: client,
		host:   host,
	}

	client.OnBeforeRequest(func(c *resty.Client, r *resty.Request) (err error) {

		r.Header.Set("js.fetch.credentials", "include")
		return
	})

	client.OnAfterResponse(func(c *resty.Client, r *resty.Response) (err error) {
		if ct := r.Header().Get("Content-Type"); !strings.HasPrefix(ct, "application/json") {
			return
		}

		defer err0.Then(&err, nil, nil)

		var resp Response[any]
		if err = json.Unmarshal(r.Body(), &resp); err != nil {
			return fmt.Errorf("%w. unmarshal resp json failed", ErrRespFormat)
		}

		if resp.Code != 0 {
			return fmt.Errorf("%w. status code %v, msg %s", ErrUnkownCode, resp.Code, resp.Message)
		}
		return
	})

	return api
}

func (api *Client) SetDebug(v bool) {
	api.client.SetDebug(v)
}

type Response[T any] struct {
	Code    int    `json:"code"`
	Data    T      `json:"data"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
}

type Page[T any] struct {
	Results  int32  `json:"results,omitempty"`
	Page     string `json:"page,omitempty"`
	Pagesize string `json:"pagesize,omitempty"`
	List     []T    `json:"list,omitempty"`
}
