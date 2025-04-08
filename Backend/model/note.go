package model


type Note struct {
	ID int 'json:"id"'
	Title string 'json: "title"'
	Content string 'json: "content"'
	Created_At string 'json: "created_at"'
}
	