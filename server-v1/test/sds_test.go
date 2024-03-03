package test

import (
	"github.com/stretchr/testify/assert"
	"server-v1/model"
	"server-v1/service/serviceImpl"
	"testing"
)

func TestSds(t *testing.T) {
	s := model.Sds{Data: []byte("test")}
	assert.Equal(t, 4, s.Len(), "s.Len() should return 4")
	assert.Equal(t, "test", s.String(), "s.String() should return 'test'")
	s.Append("ing")
	assert.Equal(t, "testing", s.String(), "s.String() should return 'testing'")
	assert.True(t, s.Update("new"), "s.Update() should return true")
	assert.Equal(t, "new", s.String(), "s.String() should return 'new'")
}

func TestSdsService(t *testing.T) {
	builder := serviceImpl.SdsServiceImpl{}
	s1 := builder.NewSds("test")
	s2 := builder.NewEmptySds()
	assert.True(t, builder.Equal(s1, model.Sds{Data: []byte("test")}),
		"builder.Equal() should return true")
	assert.True(t, builder.Equal(s2, model.Sds{Data: []byte{}}),
		"builder.Equal() should return true")
}
