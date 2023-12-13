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

func TestGoBasicMap(t *testing.T) {
	m := make(map[string]int) //map[key type]value type
	m["tucker"] = 100
	m["golang"] = 200
	m["awesome"] = 300

	assert.Equal(t, 100, m["tucker"])
	assert.Equal(t, 200, m["golang"])
	assert.Equal(t, 300, m["awesome"])
	assert.Equal(t, 0, m["aaa"]) // 없는 값을 조회했을 때는 int의 default값 0이 나온다

	_, ok := m["aaa"]   // 이런식으로 값이 존재하는지 알 수 있다 ok
	assert.False(t, ok) // 값이 없기때문에 false 나옴

	delete(m, "tucker") // delete를 사용해서 맵에서 삭제 가능
	val, ok := m["tucker"]
	assert.False(t, ok)
	assert.Equal(t, val, 0)
}
