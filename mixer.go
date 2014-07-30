package mixer

import (
	"net/http"
)

type Mixer struct {
	handlers map[string]http.Handler
}

func NewMixer() *Mixer {
	return &Mixer{
		handlers: make(map[string]http.Handler),
	}
}

func (m *Mixer) AddHost(host string, handler http.Handler) {
	m.handlers[host] = handler
}

func (m *Mixer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if handler := m.handlers[req.Host]; handler != nil {
		handler.ServeHTTP(res, req)
	} else {
		http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}
