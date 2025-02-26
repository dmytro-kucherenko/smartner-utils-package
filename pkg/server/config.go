package server

type RequestConfig[M any] struct {
	Path           string
	Status         int
	Middlewares    []M
	Meta           RequestMeta
	ProvideSession bool
}

func NewConfig[M any](meta RequestMeta) *RequestConfig[M] {
	middlewares := make([]M, 2)

	return &RequestConfig[M]{Meta: meta, Middlewares: middlewares}
}

func (config *RequestConfig[M]) WithSession() *RequestConfig[M] {
	config.ProvideSession = true

	return config
}

func (config *RequestConfig[M]) WithMiddleware(middleware M) *RequestConfig[M] {
	config.Middlewares = append(config.Middlewares, middleware)

	return config
}

func (config *RequestConfig[M]) MapRoute(path string, status int) *RequestConfig[M] {
	config.Path = path
	config.Status = status

	return config
}
