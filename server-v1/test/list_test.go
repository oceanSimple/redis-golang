package test

import (
	"github.com/stretchr/testify/assert"
	"server-v1/model"
	"testing"
)

func TestList(t *testing.T) {
	list := &model.List{}
	list.Init()
	list.PushBack(1)
	list.PushBack("test")
	list.PushFront("front")
	assert.Equal(t, 3, list.Len())
	assert.Equal(t, "front", list.Front().Value)
	assert.Equal(t, "test", list.Back().Value)

	list.Remove(list.Front())
	assert.Equal(t, 2, list.Len())
	assert.Equal(t, 1, list.Front().Value)
}
