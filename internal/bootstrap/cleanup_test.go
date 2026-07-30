package bootstrap

import (
	"context"
	"reflect"
	"testing"
)

func TestCleanupStackRunsOnceInReverseOrder(t *testing.T) {
	var order []int
	stack := &CleanupStack{}
	stack.Add(func() { order = append(order, 1) })
	stack.AddContext(func(context.Context) error {
		order = append(order, 2)
		return nil
	})
	if err := stack.Run(context.Background()); err != nil {
		t.Fatal(err)
	}
	if err := stack.Run(context.Background()); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(order, []int{2, 1}) {
		t.Fatalf("cleanup order = %v, want [2 1]", order)
	}
}
