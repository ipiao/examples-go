package context

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	ctx := context.WithValue(context.Background(), "key", "value")
	defer ctx.Done()
	GetCtxValue(ctx, "key")
}

func GetCtxValue(ctx context.Context, key string) {
	defer ctx.Done()
	v := ctx.Value(key)
	fmt.Printf("get value of key: %s, %v", key, v)
}
