package logger

import (
	"github.com/wxpusher/wxpusher-sdk-go"
	"github.com/wxpusher/wxpusher-sdk-go/model"
)

type WxPusher struct {
	ApiToken  string   `json:"apiToken" yaml:"apiToken"`
	TopicIds  []int    `json:"topicIds" yaml:"topicIds"`
	Recipient []string `json:"recipient" yaml:"recipient"`
}

func (n *WxPusher) Send(subject, context string) error {
	msg := model.NewMessage(n.ApiToken).SetSummary(subject).SetContent(context)
	if len(n.TopicIds) != 0 {
		msg.AddTopicId(n.TopicIds[0], n.TopicIds[1:]...)
	}
	if len(n.Recipient) != 0 {
		msg.AddUId(n.Recipient[0], n.Recipient[1:]...)
	}
	_, err := wxpusher.SendMessage(msg)
	return err
}
