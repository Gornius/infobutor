package message

type Message struct {
	Title   string
	Content string
	Origin  string
}

type MessageEnvelope struct {
	Message      Message
	ChannelToken string
}
