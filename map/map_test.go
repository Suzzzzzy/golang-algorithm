package _map

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashMap(t *testing.T) {
	var h HashMap[int]
	h.Add("tucker", 100)

	val, ok := h.Get("tucker")
	assert.True(t, ok)
	assert.Equal(t, 100, val)

	h.Add("golang", 200)

	val, ok = h.Get("golang")
	assert.True(t, ok)
	assert.Equal(t, 200, val)

	h.Add("awesome", 300)

	val, ok = h.Get("awesome")
	assert.True(t, ok)
	assert.Equal(t, 300, val)

}
