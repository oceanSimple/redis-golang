package test

import "testing"

func TestMap(t *testing.T) {
	m := make(map[string]string)
	add(m)
}

func add(m map[string]string) {
	m["a"] = "a"
	m["b"] = "b"
}
