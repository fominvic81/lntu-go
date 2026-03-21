package notes

//easyjson:json
type Note struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

//easyjson:json
type Notes []Note
