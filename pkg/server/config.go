package server

type RequestConfig[M any] struct {
	Path         string
	Status       int
	Interceptors []M
}

func NewConfig[I any]() *RequestConfig[I] {
	interceptors := make([]I, 0)

	return &RequestConfig[I]{Interceptors: interceptors}
}

func (config *RequestConfig[I]) WithInterceptor(interceptors ...I) *RequestConfig[I] {
	config.Interceptors = append(config.Interceptors, interceptors...)

	return config
}

func (config *RequestConfig[M]) MapRoute(path string, status int) *RequestConfig[M] {
	config.Path = path
	config.Status = status

	return config
}
