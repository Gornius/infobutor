package message

// struct that represents message being passed
type Message struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Origin  string `json:"origin"`
}

// wrapper around message with additional information
type MessageEnvelope struct {
	Message      Message
	SinkToken string
}
