package shared

type Message struct {
	Text 	string
	Color string
}

type Page struct {
	Title string
	Data string // json data to be injected
	Messages []Message
}