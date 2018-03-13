package telegrambot

func NewCheckoutHandler() Handler {
	return Handler{
		Description: "Check out handler",
		Route:       "/checkout",
	}
}
