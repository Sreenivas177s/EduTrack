package chats

type Chat struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Owner       string `json:"owner"`
	CreatedTime string `json:"created_time"`
}
