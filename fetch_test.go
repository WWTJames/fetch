package fetch

import (
	"testing"
	// "fmt"

	"github.com/xs23933/core"

)

func Test_Fetch(t *testing.T) {
	w := New(map[string]interface{}{
		// "proxy": "http://127.0.0.1:8080",
		"headers": map[string]string{
			"author": "Beasr ",
		},
	})

	buf, err := w.Payload("https://hk.nt5.net/api/authorize", map[string]interface{}{
		"user":     "admin",
		"password": "admin",
	})
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	core.Log("%s", buf)
	t.Error("")
}

func Test_Proxy(t *testing.T) {
	buf, _ := ProxyPayload("https://hk.nt5.net/api/authorize", "http://127.0.0.1:8888", map[string]interface{}{
		"user":     "admin",
		"password": "admin",
	})
	t.Errorf("%s\n", buf)
}

// Test_get 添加头信息的请求
//   params 是请求参数
//   headers 头信息
func Test_get(t *testing.T) {
	buf, _ := Get("https://hk.nt5.net/api/user/info", map[string]interface{}{
		"params": map[string]string{
			"hello": "world",
		},
		"headers": map[string]string{
			"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiYWRtaW4iLCJuYW1lIjoiQWRtaW4iLCJtb2JpbGUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AbnQ1Lm5ldCIsImdlbmRlciI6MSwidXVpZCI6ImE2ZGNkY2E1LWE3M2EtNDIyMS1hODdlLTEwZDJlNDc1NmU5YiIsInJvbGVzIjpbXSwiY3JlYXRlZF9hdCI6IjIwMTktMDctMTFUMTU6MDk6MjMrMDg6MDAiLCJ1cGRhdGVkX2F0IjoiMjAxOS0wNy0yOVQxMTowNDo1OSswODowMCIsImV4cCI6MTU2Njk3NTgxNSwiaWF0IjoxNTY0MzgzODE1fQ.sJtC5g6VRX60dE8fcqRC9QNL-n5TYD0mTkBtr3QqCHk",
		},
	})
	t.Errorf("%s\n", buf)

}
