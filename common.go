package bilireq

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/shynome/err0"
	"github.com/shynome/err0/try"
)

func (api *Client) getCSRF() (_ string, err error) {
	defer err0.Then(&err, nil, nil)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	req := try.To1(http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s/csrf", api.host), nil))
	resp := try.To1(http.DefaultClient.Do(req))
	defer resp.Body.Close()
	csrf := try.To1(io.ReadAll(resp.Body))
	return string(csrf), nil
}
