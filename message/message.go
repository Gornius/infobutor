package message

type Message struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Origin  string `json:"origin"`
}

type MessageEnvelope struct {
	Message      Message
	SinkToken string
}
