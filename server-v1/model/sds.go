package model

type Sds struct {
	Data []byte
}

type SdsInterface interface {
	Len() int
	String() string
	Append(str string)
	Update(str string) bool
}

func (s *Sds) Len() int {
	return len(s.Data)
}

func (s *Sds) String() string {
	return string(s.Data)
}

func (s *Sds) Append(str string) {
	s.Data = append(s.Data, str...)
}

func (s *Sds) Update(str string) bool {
	s.Data = []byte(str)
	return true
}
