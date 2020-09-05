package todos

type Todo struct {
	TodoID string `json:"todoId"`
	Text   string `json:"text"`
	Status string `json:"status"`
	Date   string `json:"date"`
	UserID string `json:"userId"`
}

type Todos struct {
	Data []Todo `json:"data"`
}
