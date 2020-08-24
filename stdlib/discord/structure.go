package discord

// Args : command argument object
type Args struct {
	Author  User     `json:"author"`
	Message Message  `json:"message"`
	Guild   Guild    `json:"guild"`
	Channel Channel  `json:"channel"`
	Args    []string `json:"args"`
}

type Message struct {
	ID      string  `json:"id"`
	Channel Channel `json:"channel"`
	Author  User    `json:"author"`
	Content string  `json:"content"`
}

type Guild struct {
	ID string `json:"id"`
}

type Channel struct {
	ID string `json:"id"`
}

type User struct {
	ID string `json:"id"`
}
