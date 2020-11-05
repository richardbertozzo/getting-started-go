package person

import "testing"

func TestPerson(t *testing.T) {
	p := Person{Name: "Jane", Age: 20}
	p = p.SetAge(22)
	if p.Age != 22 {
		t.Error("Esta pessoa deveria ter 22 anos")
	}
}
