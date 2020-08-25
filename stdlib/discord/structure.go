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

type Embed struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Fields      []string `json:"fields"`
}
