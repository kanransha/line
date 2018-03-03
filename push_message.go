package line

import (
	"github.com/kanransha/request"
)

type textMessage struct {
	TypeText string `json:"type"`
	Message  string `json:"text"`
}

type pushMessageContent struct {
	To       string      `json:"to"`
	Messages interface{} `json:"messages"`
}

func pushMessage(to string, messages interface{}, channel *Channel) error {
	b := &pushMessageContent{to, messages}
	return request.Post(endpoint, "./v2/bot/message/push", nil, b, channel.Header(true), nil)
}

//PushTextMessage Push single text message
func PushTextMessage(to string, message string, channel *Channel) error {
	t := &textMessage{"text", message}
	return pushMessage(to, []*textMessage{t}, channel)
}
