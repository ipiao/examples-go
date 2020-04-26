package json

import (
	"encoding/json"
	"testing"
)

type User struct {
	Name string
}

func TestUnmarshal(t *testing.T) {
	var users []*User
	str := `[{"name": "tom"}]`
	json.Unmarshal([]byte(str), &users)
	t.Logf("%v", users)
}
