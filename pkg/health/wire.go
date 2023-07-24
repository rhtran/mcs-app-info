package health

func GetHealthRouter() *Router {
	healthCheckService := NewHealthCheckService()
	return NewCheckRouter(healthCheckService)
}
