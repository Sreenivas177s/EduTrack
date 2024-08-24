package chats

type Chat struct {
	id        uint32
	name      string
	owner     string
	createdAt string //unix-epoch
}

func (chat *Chat) Name() string {
	return chat.name
}
func (chat *Chat) IsApiEnabled() bool {
	return true
}
