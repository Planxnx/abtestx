package abtestx

type Client interface {
	Run() error
	Pick() (id string, callback func() error)
}
