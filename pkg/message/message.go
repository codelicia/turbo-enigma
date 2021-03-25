package message

type Provider interface {
	SendPullRequestEvent(URL, title, author string) error
}