package handlers

type ping struct {
	msg string
}

func Ping() *handler {
	return New(ping{
		msg: "pong",
	})
}
