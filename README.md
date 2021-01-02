# go-schedule ![CI](https://github.com/mcneilcode/go-schedule/workflows/Builds/badge.svg) [![GoDoc](https://godoc.org/github.com/mcneilcode/go-schedule?status.svg)](https://godoc.org/github.com/mcneilcode/go-schedule) [![Go Report Card](https://goreportcard.com/badge/github.com/mcneilcode/go-schedule)](https://goreportcard.com/report/github.com/mcneilcode/go-schedule)

Launch a go function(s) on a schedule.

# Usage

Sample usage where two functions get scheduled on different timers,
the first runs 5 seconds, the second every 15 seconds:

```go
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	schedule "github.com/mcneilcode/go-schedule"
)

func someFunc(ctx context.Context) {
	test := "Testing 1, 2, 3"
	fmt.Println(test)
}

func someOtherFunc(ctx context.Context) {
	test := "Testing 4, 5, 6"
	fmt.Println(test)
}

func main() {
	ctx := context.Background()
	worker := schedule.New()
	worker.Add(ctx, someFunc, time.Second*time.Duration(5))
	worker.Add(ctx, someOtherFunc, time.Second*time.Duration(15))
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit
	worker.Stop()
}
```
