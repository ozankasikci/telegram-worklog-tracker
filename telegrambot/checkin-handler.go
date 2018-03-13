package telegrambot

func NewCheckinHandler() Handler {
	return Handler{
		Description: "Check in handler",
		Route:       "/checkin",
	}
}
