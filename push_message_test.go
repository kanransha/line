package line_test

import (
	"os"
	"testing"

	"github.com/kanransha/line"
)

func TestPushTextMessage(t *testing.T) {
	userID := os.Getenv("LINE_USER_ID")
	if userID == "" {
		panic("LINE_USER_ID")
	}
	token := os.Getenv("LINE_TOKEN")
	if token == "" {
		panic("LINE_TOKEN")
	}
	c := line.NewChannel(token)
	line.PushTextMessage(userID, "Text Message From TestPushTextMessage in push_message_test.go", c)
}
