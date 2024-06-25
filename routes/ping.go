package routes

import (
	"io"
	"net/http"
)

type pingService struct{}

func NewPingService() pingService {
	return pingService{}
}

func (h pingService) Get(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "pong")
}
