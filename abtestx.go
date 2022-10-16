package abtestx

var DefaultNewStrategy = NewWeightedRandom

type Test struct {
	ID            string       // Required
	Weight        float64      // Optional
	Callback      func() error // Optional
	runningWeight float64
}

type Client interface {
	Run() error
	Pick() Test
}

func New(tests []Test) Client {
	return DefaultNewStrategy(tests)
}
