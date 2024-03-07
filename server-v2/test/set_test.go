package test

import (
	"github.com/stretchr/testify/assert"
	"server-v2/builder"
	"testing"
)

func TestBuilder(t *testing.T) {
	set1 := builder.SetBuilder.New()
	set2 := builder.SetBuilder.New()

	assert.NotNil(t, set1)
	assert.NotNil(t, set2)
}

func TestAdd(t *testing.T) {
	set := builder.SetBuilder.New()
	set.Add("a", "b", "c")
	println(set.ToString())
}

func TestDelete(t *testing.T) {
	set := builder.SetBuilder.New()
	set.Add("a", "b", "c")
	set.Delete("a", "b")
	println(set.ToString())
}

func TestIntersect(t *testing.T) {
	set1 := builder.SetBuilder.New()
	set1.Add("a", "b", "c")

	set2 := builder.SetBuilder.New()
	set2.Add("b", "c", "d")

	newSet := set1.Intersection(set2)
	println(newSet.ToString())
}

func TestUnion(t *testing.T) {
	set1 := builder.SetBuilder.New()
	set1.Add("a", "b", "c")

	set2 := builder.SetBuilder.New()
	set2.Add("b", "c", "d")

	newSet := set1.Union(set2)
	println(newSet.ToString())
}

func TestDifference(t *testing.T) {
	set1 := builder.SetBuilder.New()
	set1.Add("a", "b", "c")

	set2 := builder.SetBuilder.New()
	set2.Add("b", "c", "d")

	newSet := set1.Difference(set2)
	println(newSet.ToString())
}
