package entites

func GetEntityMapping(entityName string) MiddlewareChainFactory {
	mapping := map[string]MiddlewareChainFactory{
		"chats": &ChatMWCFactory{},
	}

	mappedEntity := mapping[entityName]
	if mappedEntity == nil {
		return &Common{}
	}
	return mappedEntity
}
