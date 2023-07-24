package info

func GetInfoRouter() *Router {
	infoService := NewInfoService()
	return NewInfoRouter(infoService)
}
