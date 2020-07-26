package builtin

import "testing"

type man interface {
	SetName(string)
}

type Person struct {
	name string
}

func (p *Person) SetName(name string) {
	p.name = name
}

type Student struct {
	Person
}

func TestMan(t *testing.T) {
	var m man
	var s Student
	m = &s
	m.SetName("123")
	t.Log(m)
}
