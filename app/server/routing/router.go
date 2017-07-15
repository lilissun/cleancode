package routing

import "errors"

// Errors
var (
	ErrorRouteAlreadyRegistered = errors.New("route already registered")
	ErrorRouteDoesNotExist      = errors.New("route does not exist")
)

// Router is router
type Router struct {
	Services map[string]func([]byte) ([]byte, error)
}

// NewRouter new router
func NewRouter() *Router {
	return &Router{Services: make(map[string]func([]byte) ([]byte, error))}
}

// Register handle under name
func (router *Router) Register(name string, handle func([]byte) ([]byte, error)) error {
	if _, exist := router.Services[name]; exist {
		return ErrorRouteAlreadyRegistered
	}
	router.Services[name] = handle
	return nil
}

// Route request to name
func (router Router) Route(name string, request []byte) ([]byte, error) {
	if handle, exist := router.Services[name]; exist {
		return handle(request)
	}
	return nil, ErrorRouteDoesNotExist
}
