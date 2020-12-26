package helper

const (
	ERROR   = "ERROR"
	MESSAGE = "MESSAGE"
)

type response struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Error       bool        `json:"error"`
	Data        interface{} `json:"data"`
}

func ResponseJSON(
	messagetype, message string,
	err bool, data interface{},
) *response {
	return &response{messagetype, message, err, data}
}
