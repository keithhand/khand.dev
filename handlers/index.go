package handlers

func (h *handler) Index() *handler {
	h.viewer = Home()
	return h
}
