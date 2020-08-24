package discord

// Message : message object
type Message struct {
	UserID    string   `json:"userid"`
	MessageID string   `json:"messageid"`
	ChannelID string   `json:"channelid"`
	Content   string   `json:"content"`
	Args      []string `json:"args"`
}
