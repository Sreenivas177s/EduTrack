package entity

import "github.com/gofiber/fiber/v2/log"

type Chat struct {
	ApiBase
	Name        string `json:"name"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
	Scope       string `json:"scope"`
}

func (c Chat) New() Entity {
	return Chat{}
}

func (chat *Chat) Validate() {
	chat.Id = 1
	log.Debug("validator reached")
}
