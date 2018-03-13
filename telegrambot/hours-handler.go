package telegrambot

func NewHoursHandler() Handler {
	return Handler{
		Description: "hours handler",
		Route:       "/hours",
	}
}
