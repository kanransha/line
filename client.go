package line

const endpoint = "https://api.line.me"

//Channel Line Channel
type Channel struct {
	token string
}

//NewChannel New Channel
func NewChannel(token string) *Channel {
	return &Channel{token}
}

//Header Header
func (c Channel) Header(isToIncludeContentType bool) map[string]string {
	if isToIncludeContentType {
		return map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer " + c.token,
		}
	}
	return map[string]string{
		"Authorization": "Bearer " + c.token,
	}
}
