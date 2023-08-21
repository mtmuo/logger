package logger

import "testing"

func TestName(t *testing.T) {
	err := New(Config{
		WxPusher: &WxPusher{
			ApiToken: "",
			TopicIds: []int{},
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	// git config --global http.https://github.com.proxy socks5://127.0.0.1:10808
	module := WithModule("TestName")

	for i := 0; i < 100; i++ {
		module.Error("err")
	}
}
