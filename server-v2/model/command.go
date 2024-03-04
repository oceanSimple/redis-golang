package model

type Command struct {
	Command string
	Key     string
	Value   []string
}

func (c Command) ToString() string {
	return c.Command + " " + c.Key + " " + c.Value[0]
}
