# ABTestX

[![Go Reference](https://pkg.go.dev/badge/github.com/Planxnx/abtestx.svg)](https://pkg.go.dev/github.com/Planxnx/abtestx)
[![Go Report Card](https://goreportcard.com/badge/github.com/Planxnx/abtestx)](https://goreportcard.com/report/github.com/Planxnx/abtestx)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/f86ffbf33b4c49ffa06c4c79bef302b9)](https://www.codacy.com/gh/Planxnx/abtestx/dashboard?utm_source=github.com&utm_medium=referral&utm_content=Planxnx/abtestx&utm_campaign=Badge_Grade)
[![Code Analysis & Tests](https://github.com/Planxnx/abtestx/actions/workflows/code-analysis.yml/badge.svg)](https://github.com/Planxnx/abtestx/actions/workflows/code-analysis.yml)
[![DeepSource](https://deepsource.io/gh/Planxnx/abtestx.svg/?label=active+issues&token=zmx0Q9rLHR-7XdizRnSknm7d)](https://deepsource.io/gh/Planxnx/abtestx/?ref=repository-badge)
[![license](https://img.shields.io/badge/license-MIT-green.svg)](https://github.com/Planxnx/http-wrapper/blob/main/LICENSE)

> **A Simple A/B Testing library in Go**

abtestx is a tool to help you run A/B tests for your golang applications with minimal effort and multiple strategies.

## Requirements

- go 1.18 or higher
- [golangci-lint](https://github.com/golangci/golangci-lint)

## Installation

```shell
go get github.com/Planxnx/abtestx
```

## Strategies

- Round Robin
- Weighted Random
- Random (soon)

## Example

```go
import (
	"fmt"
	"error"
	"github.com/Planxnx/abtestx"
)

func RoundRobin() {
	abtest := abtestx.NewRoundRobin([]abtestx.RoundRobinTest{
		{
			ID: "A",
			Callback: func() error {
				fmt.Println("execute A!")
				return nil
			},
		},
		{
			ID: "B",
			Callback: func() error {
				fmt.Println("execute B!")
				return nil
			},
		},
		{
			ID: "C",
			Callback: func() error {
				return errors.New("Error C")
			},
		},
	})

	err := abtest.Run() // execute A!
	err := abtest.Run() // execute B!
	err := abtest.Run() // err is not nil, got error "Error C"
	err := abtest.Run() // execute A!
	err := abtest.Run() // execute B!

}

func WeightedRandom() {
	abtest := abtestx.NewWeightedRandom([]abtestx.WeightedRandomTest{
		{
			ID:     "A",
			Weight: 0.8,
			Callback: func() error {
				fmt.Println("execute A!")
				return nil
			},
		},
		{
			ID:     "B",
			Weight: 0.2,
			Callback: func() error {
				fmt.Println("execute B!")
				return nil
			},
		},
	})

	// will execute A 80% of the time and B 20% of the time
	_ = abtest.Run() // execute A!
	_ = abtest.Run() // execute B!
	_ = abtest.Run() // execute A!
	_ = abtest.Run() // execute A!
	_ = abtest.Run() // execute A!
}
```

## LICENSE

abtestx released under MIT license, refer [LICENSE](https://github.com/Planxnx/abtestx/blob/main/LICENSE) file.
