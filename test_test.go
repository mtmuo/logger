package logger

import "testing"

func TestName(t *testing.T) {
	err := New(Config{
		WxPusher: &WxPusher{
			ApiToken: "AT_ap9VpRWycDS73FlX4Dza8NI9e3PyxSmF",
			TopicIds: []int{11326},
		},
	})
	if err != nil {
		t.Error(err)
		return
	}

	module := WithModule("TestName")

	for i := 0; i < 100; i++ {
		module.Error("err")
	}
}
