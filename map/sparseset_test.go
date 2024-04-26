package _map

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSparseSet(t *testing.T) {
	s := NewSparseSet[string, int]()
	s.Add("aaa", 30)
	s.Add("bbb", 20)
	s.Add("ccc", 10)

	v, ok := s.Get("bbb")
	assert.True(t, ok)
	assert.Equal(t, 20, v)

	v, ok = s.Get("ddd")
	assert.False(t, ok)

	removed := s.Remove("bbb")
	assert.True(t, removed)

	_, ok = s.Get("bbb")
	assert.False(t, ok)

	for i := s.Iterator(); !i.IsEnd(); i.Next() {
		elem := i.Get()
		if elem.Key == "aaa" {
			assert.Equal(t, 30, elem.Value)
		} else if elem.Key == "ccc" {
			assert.Equal(t, 10, elem.Value)
		} else {
			t.Fail()
		}
	}
}
