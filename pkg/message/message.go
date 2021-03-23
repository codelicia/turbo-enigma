package message

type Message interface {
	SendPullRequestEvent(URL, title, author string) error
}