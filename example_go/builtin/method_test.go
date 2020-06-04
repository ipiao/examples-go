package builtin

import "testing"

type User struct {
	Name string
}

func (u *User) GetName() string {
	if u == nil {
		return ""
	}
	return u.Name
}

func TestNilCall(t *testing.T) {
	u := &User{Name: "nas"}

	t.Log(u.GetName())

	var u1 *User
	t.Log(u1.GetName())
}
