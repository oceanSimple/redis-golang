package builder

import "server-v3/model"

type setBuilder struct {
	set *model.Set
}

func (b *setBuilder) New() *model.Set {
	newSet := &model.Set{}
	newSet.Init()
	return newSet
}
