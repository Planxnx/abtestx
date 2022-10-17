package abtestx

// Client is a a/b testing instance interface.
type Client interface {
	Run() error
	Pick() (id string, callback func() error)
}
