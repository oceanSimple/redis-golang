package model

import "strings"

type setInterface interface {
	Init()
	Add(str ...string)
	Delete(str ...string)
	Has(str string) bool
	ToString() string
	Intersection(set *Set) *Set
	Union(set *Set) *Set
	Difference(set *Set) *Set
}

type Set struct {
	data map[string]bool
}

func (s *Set) Init() {
	s.data = make(map[string]bool)
}

func (s *Set) Add(str ...string) {
	for _, v := range str {
		s.data[v] = true
	}
}

func (s *Set) Delete(str ...string) {
	for _, v := range str {
		s.data[v] = false
	}
}

func (s *Set) Has(str string) bool {
	return s.data[str]
}

func (s *Set) ToString() string {
	builder := strings.Builder{}
	builder.WriteString("{")
	for k, v := range s.data {
		if v {
			builder.WriteString(k)
			builder.WriteString(", ")
		}
	}
	result := builder.String()
	if builder.Len() > 1 {
		result = result[:builder.Len()-2]
	}
	result += "}"
	return result
}

func (s *Set) Intersection(set *Set) *Set {
	result := &Set{}
	result.Init()

	for k, v := range s.data {
		if v && set.data[k] {
			result.data[k] = true
		}
	}

	return result
}

func (s *Set) Union(set *Set) *Set {
	result := &Set{}
	result.Init()

	for k, v := range s.data {
		if v {
			result.data[k] = true
		}
	}

	for k, v := range set.data {
		if v {
			result.data[k] = true
		}
	}

	return result
}

func (s *Set) Difference(set *Set) *Set {
	result := &Set{}
	result.Init()

	for k, v := range s.data {
		if v && !set.data[k] {
			result.data[k] = true
		}
	}

	return result
}
