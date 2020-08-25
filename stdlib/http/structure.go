package http

type Response struct {
	Status int  `json:"status"`
	Data   Data `json:"data"`
}

type Data struct {
	Text    string            `json:"text"`
	Headers map[string]string `json:"headers"`
}
