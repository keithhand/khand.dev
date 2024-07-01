package handlers

type ping struct {
	msg string
}

func (h *handler) Ping() *handler {
	h.viewer = PingIndex("pong")
	return h
}
