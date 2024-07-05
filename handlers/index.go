package handlers

type index struct{}

func Index() *handler {
	return New(index{})
}
