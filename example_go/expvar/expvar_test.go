package expvar

import (
	"expvar"
	"testing"
)

func TestExpvar(t *testing.T) {
	i := expvar.NewInt("testInt")
	i.Set(8)
	t.Log(i)

	expvar.Publish("testInt2", i)

	v := expvar.Get("testInt2")

	t.Log(v == i)
}
