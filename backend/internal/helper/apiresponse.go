package helper

import "github.com/haobuhaoo/gossip-with-go/internal/api"

// ParseResponseDataAndMessage packages the data into an API Response with the success message.
func ParseResponseDataAndMessage(data []byte, msg string) api.Response {
	return api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{msg},
	}
}

// ParseResponseMessage packages the message into an API Response.
func ParseResponseMessage(msg string) api.Response {
	return api.Response{
		Messages: []string{msg},
	}
}
