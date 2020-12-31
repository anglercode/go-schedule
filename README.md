# go-schedule
Launch a go function(s) on a schedule.

# Usage

```go
import (
    "context"
    "os"
    "os/signal"
    "github.com/mcneilcode/go-schedule"
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