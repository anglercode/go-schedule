package schedule

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestAddStopSingle(t *testing.T) {
	ctx := context.Background()
	worker := New()
	worker.Add(ctx, testFunc, time.Second*time.Duration(5))
	worker.Stop()
}

func TestAddStopMulti(t *testing.T) {
	ctx := context.Background()
	worker := New()
	worker.Add(ctx, testFunc, time.Second*time.Duration(5))
	worker.Add(ctx, testFunc, time.Second*time.Duration(5))
	worker.Add(ctx, testFunc, time.Second*time.Duration(5))
	worker.Stop()
}

// Test Helpers
func testFunc(ctx context.Context) {
	test := "Testing 1, 2, 3"
	fmt.Println(test)
}
