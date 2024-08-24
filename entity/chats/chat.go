package chats

type Chat struct {
	id        uint32 `json:"id"`
	name      string `json:"name"`
	owner     string `json:"owner"`
	createdAt string `json:"created_at`
}

func (chat *Chat) Name() string {
	return chat.name
}
func (chat *Chat) IsApiEnabled() bool {
	return true
}
