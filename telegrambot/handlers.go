package telegrambot

type Handler struct {
	Description string
	Route       string
}

func GetHandlers() []Handler {
	var handlers []Handler

	handlers = append(
		handlers,
		NewCheckinHandler(),
		NewCheckoutHandler(),
		NewHoursHandler(),
	)

	return handlers
}
